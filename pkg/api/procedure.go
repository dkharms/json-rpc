package api

type ProcedureName string

type ProcedureVersion string

type ProcedureHandler func(*JsonRequest, *JsonResponse) error

type Procedure struct {
	Name    ProcedureName
	Version ProcedureVersion
	Handler ProcedureHandler
}

type ProcedureMap struct {
	m map[ProcedureVersion]Procedure
}

func (m *ProcedureMap) Get(version ProcedureVersion) (Procedure, bool) {
	impl, ok := m.m[version]
	return impl, ok
}

func (m *ProcedureMap) Add(procedure Procedure) {
	if m.m == nil {
		m.m = map[ProcedureVersion]Procedure{}
	}
	m.m[procedure.Version] = procedure
}
