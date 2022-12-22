package service

import (
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/pattern"
	goModel "github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	golang "github.com/tilau2328/x/src/go/service/repo/go/package/service"
	"log"
	"os"
)

//go:generate protoc -I=. --go_out=. --go-grpc_out=. ./testdata/test.proto
type (
	Options struct {
		Golang *golang.Service
	}
	Service struct {
		Options
	}
)

func NewService(opts Options) *Service {
	return &Service{Options: opts}
}

func (s *Service) Generate(props pattern.AdaptorProps) error {
	var modelStructs []goModel.Struct
	var protoStructs []goModel.Struct
	var grpcInterfaces []goModel.Interface
	var providerInterface goModel.Interface
	defer func() {
		log.Println(recover())
		log.Println(providerInterface, grpcInterfaces, protoStructs, modelStructs)
	}()
	if err := s.Golang.Search(os.Stdout, props.Provider, &providerInterface); err != nil {
		return err
	} else if err := s.Golang.Search(os.Stdout, props.Grpc, &grpcInterfaces); err != nil {
		return err
	} else if err := s.Golang.Search(os.Stdout, props.Proto, &protoStructs); err != nil {
		return err
	} else if err := s.Golang.Search(os.Stdout, props.Models, &modelStructs); err != nil {
		return err
	} else if err := x.NewFile(props.Mappers+".go", func(file *os.File) error {
		return s.Golang.Create(file, goModel.CreateReq{}, goModel.Mapper{})
	}); err != nil {
		return err
	} else if err := x.NewFile(props.Handler+".go", func(file *os.File) error {
		return s.Golang.Create(file, goModel.CreateReq{}, goModel.Service{})
	}); err != nil {
		return err
	} else if err := x.NewFile(props.Requester+".go", func(file *os.File) error {
		return s.Golang.Create(file, goModel.CreateReq{}, goModel.Service{})
	}); err != nil {
		return err
	}
	return nil
}
func (s *Service) Create() error {
	return nil
}
func (s *Service) Modify() error {
	return nil
}
