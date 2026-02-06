package providers

import "context"

type Provider interface {
	Complete(ctx context.Context, prompt string) (string, error)
	Health(ctx context.Context) error
}
