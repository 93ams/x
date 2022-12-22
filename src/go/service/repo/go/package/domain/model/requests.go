package model

type (
	CreateReq struct {
		Pattern bool
		Pkg     string
	}
	SearchReq struct {
		File string
		Name string
	}
	ModifyReq struct {
	}
)
