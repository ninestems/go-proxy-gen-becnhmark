package implementation

import (
	"context"
)

//go:generate go-proxy-gen --log-level=debug

type Converter interface {
	// Convert converts id int64 to string.
	//
	// goproxygen:
	//  log ctx::traceID::trace_id
	//  log input::id:string
	//  log output::int64::out_alias
	//  trace ctx::traceID::trace_id
	//  trace input::id:string
	//  trace output::int64::out_alias
	Convert(ctx context.Context, id string) (int64, error)
}
