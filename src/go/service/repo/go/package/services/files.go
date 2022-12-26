package services

import (
	"context"
	"github.com/samber/lo"
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/package/adaptor/driver"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"github.com/tilau2328/x/src/go/service/repo/go/package/services/provider"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type Service struct {
}

var _ provider.GolangProvider = &Service{}

func NewService() *Service {
	return &Service{}
}

func (d *Service) Read(_ context.Context, in ...model.ReadReq) (ret map[string]*model.File, err error) {
	return driver.ReadFiles(lo.Map(in, func(item model.ReadReq, _ int) string { return filepath.Join(item.Dir, ensureDotGo(item.Name)) })...)
}
func (d *Service) Search(ctx context.Context, in ...model.SearchReq) (map[string][]model.Node, error) {
	ret := map[string][]model.Node{}
	lock := sync.Mutex{}
	if err := x.ParallelTry(func(v model.SearchReq) error {
		_, err := driver.SearchFiles(func(node model.Node) bool {
			switch n := node.(type) {
			case *model.Func:
				if v.Func != nil && v.Func.Filter(n) {
					lock.Lock()
					ret[v.Id] = append(ret[v.Id], n.Clone())
					lock.Unlock()
				}
			case *model.File:
				if v.Package {
					lock.Lock()
					ret[v.Id] = append(ret[v.Id], n.Name.Clone())
					lock.Unlock()
				}
			case *model.Type:
				if v.Type != nil && v.Type.Filter(n) {
					lock.Lock()
					ret[v.Id] = append(ret[v.Id], n.Clone())
					lock.Unlock()
				}
			}
			return false
		}, filepath.Join(v.Dir, ensureDotGo(v.Name)))
		return err
	}, in); err != nil {
		return nil, err
	}
	return ret, nil
}
func (d *Service) Create(_ context.Context, in ...model.CreateReq) error {
	return driver.CreateFiles(lo.SliceToMap(in, func(item model.CreateReq) (string, *model.File) {
		return filepath.Join(item.Dir, ensureDotGo(item.Name)), item.File
	}))
}
func (d *Service) Modify(_ context.Context, in ...model.ModifyReq) error {
	return driver.ModifyFiles(lo.SliceToMap(in, func(item model.ModifyReq) (string, driver.FileMod) {
		return filepath.Join(item.Dir, ensureDotGo(item.Name)), func(file *model.File) (*model.File, error) {
			return file, nil
		}
	}))
}
func (d *Service) Delete(_ context.Context, in ...model.DeleteReq) error {
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
func ensureDotGo(in string) string {
	if !strings.HasSuffix(in, ".go") {
		in += ".go"
	}
	return in
}
