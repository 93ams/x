package model

func ToKeySpace(in KeySpace) model.KeySpace {
	return model.KeySpace{}
}
func FromKeySpaces(in []model.KeySpace) []*KeySpace {
	ret := make([]*KeySpace, len(in))
	for i, v := range in {
		ret[i] = FromKeySpace(v)
	}
	return ret
}
func FromKeySpace(in model.KeySpace) *KeySpace {
	return &KeySpace{
		Name:    in.KeySpaceKey.String(),
		Durable: &in.Durable,
	}
}
