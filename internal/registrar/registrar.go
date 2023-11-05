package registrar

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewNacos)
