package bot_jwt

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"
	
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/security"
	"github.com/forta-network/forta-node/clients"
	"github.com/forta-network/forta-node/clients/messaging"
	"github.com/forta-network/forta-node/config"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// JWTProvider provides jwt tokens to bots, signed with node's private key..
type JWTProvider struct {
	botConfigs      []config.AgentConfig
	botConfigsMutex sync.RWMutex
	
	dockerClient clients.DockerClient
	
	// msgClient to subscribe to bot changes
	msgClient clients.MessageClient
	
	cfg *JWTProviderConfig
	
	lastErr health.ErrorTracker
	
	srv *http.Server
}

type JWTProviderConfig struct {
	Key    *keystore.Key
	Config config.Config
}

func NewBotJWTProvider(
	cfg config.Config,
) (*JWTProvider, error) {
	key, err := security.LoadKey(config.DefaultContainerKeyDirPath)
	if err != nil {
		return nil, err
	}
	
	return initProvider(
		&JWTProviderConfig{
			Key:    key,
			Config: cfg,
		},
	)
}

func initProvider(cfg *JWTProviderConfig) (*JWTProvider, error) {
	globalClient, err := clients.NewDockerClient("")
	if err != nil {
		return nil, fmt.Errorf("failed to create the global docker client: %v", err)
	}
	
	msgClient := messaging.NewClient(
		"jwt-provider", fmt.Sprintf(
			"%s:%s", config.DockerNatsContainerName,
			config.DefaultNatsPort,
		),
	)
	
	return &JWTProvider{dockerClient: globalClient, msgClient: msgClient, cfg: cfg}, nil
}

// Start spawns a jwt provider routine and returns.
func (j *JWTProvider) Start() error {
	return j.StartWithContext(context.Background())
}

func (j *JWTProvider) Stop() error {
	return j.srv.Close()
}

// StartWithContext subscribe to bot updates and spawn a Bot JWT Provider http server.
func (j *JWTProvider) StartWithContext(ctx context.Context) error {
	j.registerMessageHandlers()
	
	if j.cfg.Config.JWTProvider.Addr == "" {
		j.cfg.Config.JWTProvider.Addr = fmt.Sprintf(":%s", config.DefaultBotJWTProviderPort)
	}
	
	// setup routes
	r := mux.NewRouter()
	r.HandleFunc("/create", j.createJWTHandler).Methods(http.MethodPost)
	
	j.srv = &http.Server{
		Addr:    j.cfg.Config.JWTProvider.Addr,
		Handler: r,
	}
	
	go func() {
		err := j.listenAndServeWithContext(ctx)
		if err != nil {
			logrus.WithError(err).Panic("server error")
		}
	}()
	
	return nil
}

func (j *JWTProvider) listenAndServeWithContext(ctx context.Context) error {
	errChan := make(chan error)
	
	go func() {
		logrus.Infof("Starting Bot JWT Provider Service on: %s", j.srv.Addr)
		err := j.srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()
	
	// gracefully handle stopping server
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		_ = j.srv.Close()
		return nil
	}
}

// agentIDReverseLookup reverse lookup from ip to agent id.
func (j *JWTProvider) agentIDReverseLookup(ctx context.Context, remoteAddr string) (string, error) {
	ipAddr, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return "", err
	}
	
	containers, err := j.dockerClient.GetContainers(ctx)
	if err != nil {
		return "", err
	}
	
	var agentContainerNames []string
	for _, container := range containers {
		for _, network := range container.NetworkSettings.Networks {
			if network.IPAddress == ipAddr {
				agentContainerNames = container.Names
				break
			}
		}
		
		if agentContainerNames != nil {
			break
		}
	}
	
	if len(agentContainerNames) == 0 {
		return "", fmt.Errorf("can not find bot id of %s", ipAddr)
	}
	
	containerName := agentContainerNames[0][1:]
	for _, agentConfig := range j.botConfigs {
		if agentConfig.ContainerName() == containerName {
			return agentConfig.ID, nil
		}
	}
	
	return "", fmt.Errorf("no bots with ip: %s exist", ipAddr)
}

func (j *JWTProvider) registerMessageHandlers() {
	j.msgClient.Subscribe(messaging.SubjectAgentsVersionsLatest, messaging.AgentsHandler(j.botUpdateHandler))
}

func (j *JWTProvider) botUpdateHandler(payload messaging.AgentPayload) error {
	j.botConfigsMutex.Lock()
	j.botConfigs = payload
	j.botConfigsMutex.Unlock()
	return nil
}

func (j *JWTProvider) testAPI(_ context.Context) {
	j.lastErr.Set(nil)
}

func (j *JWTProvider) apiHealthChecker(ctx context.Context) {
	j.testAPI(ctx)
	ticker := time.NewTicker(time.Minute * 5)
	for range ticker.C {
		j.testAPI(ctx)
	}
}

func (j *JWTProvider) Name() string {
	return "bot-jwt-provider"
}

func (j *JWTProvider) Health() health.Reports {
	return health.Reports{
		j.lastErr.GetReport("api"),
	}
}

// requestHash used for "hash" claim in JWT token
func requestHash(uri string, payload []byte) common.Hash {
	requestStr := fmt.Sprintf("%s%s", uri, payload)
	
	return crypto.Keccak256Hash([]byte(requestStr))
}

// CreateBotJWT returns a bot JWT token. Basically security.ScannerJWT with bot&request info.
func CreateBotJWT(key *keystore.Key, agentID string, hash string, exp uint64) (string, error) {
	claims := map[string]interface{}{
		"bot":  agentID,
		"hash": hash,
	}
	
	// security.CreateScannerJWT has already a default 30sec expiry
	if exp != 0 {
		claims["exp"] = exp
	}
	
	return security.CreateScannerJWT(key, claims)
}
