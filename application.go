package main

import (
	"context"
	"log/slog"

	"capnproto.org/go/capnp/v3"
	"github.com/gligneul/rollmelette"
)

type MyApplication struct{}

func (a *MyApplication) Advance(
	env rollmelette.Env,
	metadata rollmelette.Metadata,
	deposit rollmelette.Deposit,
	payload []byte,
) error {
	msg, err := capnp.Unmarshal(payload)
	if err != nil {
		return err
	}
	_, err = ReadRootAdvanceRequest(msg)
	if err != nil {
		return err
	}
	return nil
}

func (a *MyApplication) Inspect(env rollmelette.EnvInspector, payload []byte) error {
	// Handle inspect input
	return nil
}

func main() {
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	app := new(MyApplication)
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
