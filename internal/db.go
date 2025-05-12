package internal

import (
	"context"
	"errors"

	"github.com/valkey-io/valkey-go"
)

type DB interface {
	Save(ctx context.Context, key string, message string) error
	Get(ctx context.Context, key string) (string, error)
}

type Valkey struct {
	Client valkey.Client
}

func NewValkey(address string) (*Valkey, error) {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{address}})

	if err != nil {
		return nil, err
	}

	return &Valkey{
		Client: client,
	}, nil
}

func (v *Valkey) Save(ctx context.Context, key string, message string) error {
	exists, _ := v.Client.Do(ctx, v.Client.B().Get().Key(key).Build()).AsBool()

	if exists {
		return errors.New("Key already exists")
	}

	if err := v.Client.Do(ctx, v.Client.B().Set().Key(key).Value(message).Build()).Error(); err != nil {
		return err
	}

	return nil
}

func (v *Valkey) Get(ctx context.Context, key string) (string, error) {
	exists, error := v.Client.Do(ctx, v.Client.B().Get().Key(key).Build()).ToString()

	if error != nil {
		return "", error
	}

	return exists, nil
}
