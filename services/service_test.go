package services

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"os"
	"syscall"
	"testing"
	"time"
)

type TestService struct {
	cancelled bool
	ctx       context.Context
}

func (t *TestService) Start() error {
	grp, ctx := errgroup.WithContext(t.ctx)
	grp.Go(func() error {
		select {
		case <-ctx.Done():
			t.cancelled = true
			return ctx.Err()
		}
	})
	return grp.Wait()
}

func (t *TestService) Stop() error {
	return nil
}

func (t *TestService) Name() string {
	return "test"
}

func TestSigIntSignalCancelsService(t *testing.T) {
	sigc = make(chan os.Signal, 1)
	ctx, _ := InitMainContext()

	go func() {
		time.Sleep(1 * time.Second)
		sigc <- syscall.SIGINT
	}()

	svc := &TestService{ctx: ctx}
	err := StartServices(ctx, []Service{svc})
	assert.Error(t, err, context.Canceled)
	assert.True(t, svc.cancelled)
}
