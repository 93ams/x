package service

import (
	"grpc/package/domain/builder"
	"io"
)

type Service struct {
}

func (s *Service) Transpile(w io.Writer, cfg builder.File) error {
	return nil
}
func (s *Service) Builder(w io.Writer, cfg builder.File) error {
	return nil
}
