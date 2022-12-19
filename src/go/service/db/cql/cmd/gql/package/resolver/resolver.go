package resolver

type Resolver struct {
	ddl provider.DDL
	dml provider.DML
}

func NewResolver(
	ddl provider.DDL,
	dml provider.DML,
) *Resolver {
	return &Resolver{
		ddl: ddl,
		dml: dml,
	}
}
