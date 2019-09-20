package utils

import "github.com/google/wire"

// Providers contain all providers from current package
var Providers = wire.NewSet(
	ProvideFile,
	ProvideSignature,
	ProvideString,
	ProvideSystem,
)
