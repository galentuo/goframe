package goframe

import "context"

// assert that defaultContext implementations are fulfilling their interfaces & context.Context
var _ context.Context = &defaultServerContext{}
var _ context.Context = &defaultContext{}
var _ Context = &defaultContext{}
var _ ServerContext = &defaultServerContext{}
