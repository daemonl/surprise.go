package cx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestCancelCause(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(fmt.Errorf("ERR"))

	runFunc(ctx, t)
}

func TestCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	runFunc(ctx, t)
}

func runFunc(ctx context.Context, t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = req.WithContext(ctx)

	_, err = http.DefaultClient.Do(req)
	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, context.Canceled) {
		// This fails on TestCancelCause
		t.Errorf("Do() returned non cancelErr: %s", err)
	}
}
