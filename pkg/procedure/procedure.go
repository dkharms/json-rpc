package procedure

import "github.com/dkharms/json-rpc/pkg/server"

type Implementation func(*server.JsonRequest, *server.JsonResponse) error

type Procedure struct {
	Name    string
	Version string
	Impl    Implementation
}

func New(name string, version string, impl Implementation) Procedure {
	return Procedure{
		Name:    name,
		Version: version,
		Impl:    impl,
	}
}
