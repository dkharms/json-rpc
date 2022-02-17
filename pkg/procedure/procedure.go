package procedure

import "github.com/dkharms/json-rpc/pkg/api"

type Name string

type Version string

type Implementation func(*api.JsonRequest, *api.JsonResponse) error

type Procedure struct {
	Name    Name
	Version Version
	Impl    Implementation
}

func New(name string, version string, impl Implementation) Procedure {
	return Procedure{
		Name:    Name(name),
		Version: Version(version),
		Impl:    impl,
	}
}

type Map struct {
	m map[Version]Procedure
}

func (m *Map) Get(version Version) (Procedure, bool) {
	impl, ok := m.m[version]
	return impl, ok
}

func (m *Map) Add(procedure Procedure) {
	if m.m == nil {
		m.m = map[Version]Procedure{}
	}
	m.m[procedure.Version] = procedure
}
