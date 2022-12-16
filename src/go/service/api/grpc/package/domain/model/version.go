package model

func (s *Syntax) Version() int {
	switch s.ProtobufVersion {
	case "proto3":
		return 3
	case "proto2":
		return 2
	default:
		return 0
	}
}
