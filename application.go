package main

import (
	"context"
	"fmt"
	"log/slog"

	"capnproto.org/go/capnp/v3"
	"github.com/gligneul/rollmelette"
)

type CalculatorApp struct {
	Value int64
}

func (a *CalculatorApp) Advance(
	env rollmelette.Env,
	metadata rollmelette.Metadata,
	deposit rollmelette.Deposit,
	payload []byte,
) error {
	msg, err := capnp.Unmarshal(payload)
	if err != nil {
		return err
	}
	req, err := ReadRootAdvanceRequest(msg)
	if err != nil {
		return err
	}
	switch req.Which() {
	case AdvanceRequest_Which_add:
		a.Value += req.Add().Operand()
	case AdvanceRequest_Which_mul:
		a.Value *= req.Mul().Operand()
	case AdvanceRequest_Which_div:
		a.Value /= req.Div().Operand()
	default:
		return fmt.Errorf("unknown advance request: %v", req)
	}
	return nil
}

func (a *CalculatorApp) Inspect(env rollmelette.EnvInspector, payload []byte) error {
	msg := fmt.Sprintf("%v", a.Value)
	env.Report([]byte(msg))
	return nil
}

func main() {
	ctx := context.Background()
	opts := rollmelette.NewRunOpts()
	app := new(CalculatorApp)
	err := rollmelette.Run(ctx, opts, app)
	if err != nil {
		slog.Error("application error", "error", err)
	}
}
