package service

import (
	"grpc/package/model"
	"io"
)

type Service struct {
}

func (s *Service) Transpile(w io.Writer, cfg model.File) error {
	return nil
}
func (s *Service) Builder(w io.Writer, cfg model.File) error {
	return nil
}
