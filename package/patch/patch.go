package patch

type (
	Op    string
	Patch struct {
		Path string
		Op   string
		Val  any
	}
)

const (
	OpAdd Op = "add"
	OpSet Op = "set"
	OpRem Op = "rem"
)
