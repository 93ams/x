package service

import (
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/api/grpc/package/pattern"
	goModel "github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	golang "github.com/tilau2328/x/src/go/service/repo/go/package/services"
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
	modelStructs := new([]goModel.Struct)
	protoStructs := new([]goModel.Struct)
	grpcInterfaces := new([]goModel.Interface)
	providerInterface := new(goModel.Interface)
	defer func() {
		log.Println(providerInterface, grpcInterfaces, protoStructs, modelStructs)
	}()
	if err := s.Golang.Search(os.Stdout, props.Provider, providerInterface); err != nil {
		return err
	} else if err := s.Golang.Search(os.Stdout, props.Grpc, grpcInterfaces); err != nil {
		return err
	} else if err := s.Golang.Search(os.Stdout, props.Proto, protoStructs); err != nil {
		return err
	} else if err := s.Golang.Search(os.Stdout, props.Models, modelStructs); err != nil {
		return err
		//} else if err := x.NewFile(props.Mappers+".go", func(file *os.File) error {
		//	return s.Golang.Create(file, goModel.CreateReq{
		//		Pkg: "mapper",
		//	}, goModel.Mapper{
		//		Val: goModel.Struct{},
		//		Key:   goModel.Struct{},
		//	})
		//}); err != nil {
		//	return err
	} else if err := x.NewFile(props.Handler+".go", func(file *os.File) error {
		return s.Golang.Create(file, model.CreateReq{
			Pkg: "adaptor",
		}, goModel.Service{
			Struct: goModel.Struct{
				Ident: goModel.Ident{
					Name: "Handler",
				},
			},
			Methods: lo.Map(providerInterface.Methods, func(item goModel.FuncType, index int) goModel.Func {
				return goModel.Func{
					Receiver: &goModel.Ident{Name: "Handler", Ptr: true},
					FuncType: item,
				}
			}),
		})
	}); err != nil {
		return err
	} else if err := x.NewFile(props.Requester+".go", func(file *os.File) error {
		return s.Golang.Create(file, model.CreateReq{
			Pkg: "adaptor",
		}, goModel.Service{
			Struct: goModel.Struct{
				Ident: goModel.Ident{
					Name: "Requester",
				},
			},
			Methods: lo.Map(providerInterface.Methods, func(item goModel.FuncType, index int) goModel.Func {
				return goModel.Func{
					Receiver: &goModel.Ident{Name: "Requester", Ptr: true},
					FuncType: item,
				}
			}),
		})
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
