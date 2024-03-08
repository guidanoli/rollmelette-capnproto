package main

import (
	"context"
	"fmt"
	"log/slog"

	"capnproto.org/go/capnp/v3"
	"github.com/gligneul/rollmelette"
)

type CalculatorApp struct {
	Value   int64
	OpCount int64
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
	switch which := req.Which(); which {
	case AdvanceRequest_Which_add:
		a.Value += req.Add().Operand()
		a.OpCount += 1
	case AdvanceRequest_Which_mul:
		a.Value *= req.Mul().Operand()
		a.OpCount += 1
	case AdvanceRequest_Which_div:
		a.Value /= req.Div().Operand()
		a.OpCount += 1
	default:
		return fmt.Errorf("unknown advance request: %v", which)
	}
	return nil
}

func (a *CalculatorApp) Inspect(env rollmelette.EnvInspector, payload []byte) error {
	msg, err := capnp.Unmarshal(payload)
	if err != nil {
		return err
	}
	req, err := ReadRootInspectRequest(msg)
	if err != nil {
		return err
	}
	switch which := req.Which(); which {
	case InspectRequest_Which_value:
		msg := fmt.Sprintf("%v", a.Value)
		env.Report([]byte(msg))
	case InspectRequest_Which_opCount:
		msg := fmt.Sprintf("%v", a.OpCount)
		env.Report([]byte(msg))
	default:
		return fmt.Errorf("unknown inspect request: %v", which)
	}
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
