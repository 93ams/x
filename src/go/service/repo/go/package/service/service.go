package service

import "context"

type (
	CreateIn struct {
	}
	SearchIn struct {
	}
	UpdateIn struct {
	}
	DeleteIn struct {
	}
	GoService struct {
	}
)

func NewGoService() *GoService {
	return &GoService{}
}
func (s *GoService) Create(ctx context.Context, in CreateIn) error {

	return nil
}
func (s *GoService) Search(ctx context.Context, in SearchIn) error {
	return nil
}
func (s *GoService) Update(ctx context.Context, in UpdateIn) error {
	return nil
}
func (s *GoService) Delete(ctx context.Context, in DeleteIn) error {
	return nil
}
