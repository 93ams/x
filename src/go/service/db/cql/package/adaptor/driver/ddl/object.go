package ddl

import "strings"

type Object map[string]string

func (o Object) String() string {
	var ret []string
	for k, v := range o {
		ret = append(ret, "'"+k+"':"+"'"+v+"'")
	}
	return "{ " + strings.Join(ret, ", ") + "}"
}
