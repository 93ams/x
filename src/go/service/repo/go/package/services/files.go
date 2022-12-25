package services

import (
	"context"
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/provider"
	"path"
	"path/filepath"
)

type Service struct {
}

var _ provider.GolangProvider = &Service{}

func NewService() *Service {
	return &Service{}
}

func (d *Service) Read(ctx context.Context, in ...model.ReadReq) (ret map[string]*model.File, err error) {
	return driver.ReadFiles(lo.Map(in, func(item model.ReadReq, _ int) string { return filepath.Join(item.Dir, item.Name) })...)
}
func (d *Service) Search(ctx context.Context, in ...model.SearchReq) (map[string][]model.Node, error) {
	ret := map[string][]model.Node{}
	_, err := driver.SearchFiles(func(node model.Node) bool {
		switch n := node.(type) {
		case *model.Type:
			for _, v := range in {
				if v.Type.Filter(n) {
					path := filepath.Join(v.Dir, v.Name)
					ret[path] = append(ret[path], n)
				}
			}
		}
		return false
	}, lo.Map(in, func(item model.SearchReq, _ int) string { return filepath.Join(item.Dir, item.Name) })...)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
func (d *Service) Create(ctx context.Context, in ...model.CreateReq) error {
	return driver.CreateFiles(lo.SliceToMap(in, func(item model.CreateReq) (string, *model.File) {
		return item.Name, item.File
	}))
}
func (d *Service) Modify(ctx context.Context, in ...model.ModifyReq) error {
	return driver.ModifyFiles(lo.SliceToMap(in, func(item model.ModifyReq) (string, driver.FileMod) {
		return item.Name, func(file *model.File) (*model.File, error) {
			return file, nil
		}
	}))
}
func (d *Service) Delete(ctx context.Context, in ...model.DeleteReq) error {
	var files, dirs []string
	for _, v := range in {
		if v.Name != "" {
			files = append(files, path.Join(v.Dir, v.Name))
		} else if v.Dir != "" {
			dirs = append(dirs, v.Dir)
		}
	}
	return x.ConcurrentTry(
		func() error { return x.DeleteDirs(dirs...) },
		func() error { return x.DeleteFiles(files...) },
	)
}
