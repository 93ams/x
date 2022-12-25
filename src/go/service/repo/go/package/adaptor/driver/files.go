package driver

import (
	"github.com/tilau2328/x/src/go/package/x"
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
	"go/parser"
	"go/token"
	"os"
)

type (
	Filter  func(model.Node) bool
	FileMod func(*model.File) (*model.File, error)
)

func ReadFile(file string) (*model.File, error) {
	return ParseFile(token.NewFileSet(), file, nil, parser.AllErrors)
}
func ReadFiles(files ...string) (map[string]*model.File, error) {
	ret := map[string]*model.File{}
	if err := x.ParallelTry(func(name string) (err error) {
		ret[name], err = ReadFile(name)
		return
	}, files); err != nil {
		return nil, err
	}
	return ret, nil
}
func CreateFiles(files map[string]*model.File) error { return x.ParallelMapTry(CreateFile, files) }
func CreateFile(file string, node *model.File) error {
	return x.NewFile(file, func(file *os.File) error { return Write(file, node) })
}
func ModifyFiles(files map[string]FileMod) error { return x.ParallelMapTry(ModifyFile, files) }
func ModifyFile(file string, fn FileMod) error {
	if readFile, err := ReadFile(file); err != nil {
		return err
	} else if writeFile, err := fn(readFile); err != nil {
		return err
	} else if err := CreateFile(file, writeFile); err != nil {
		return err
	}
	return nil
}
func SearchFiles(filter Filter, files ...string) (map[string][]model.Node, error) {
	ret := map[string][]model.Node{}
	if err := x.ParallelTry(func(name string) error {
		if file, err := ReadFile(name); err != nil {
			return err
		} else {
			ret[name] = SearchFile(file, filter)
		}
		return nil
	}, files); err != nil {
		return nil, err
	}
	return ret, nil
}
func SearchFile(node *model.File, filter Filter) (ret []model.Node) {
	model.Inspect(node, func(node model.Node) bool {
		if filter(node) {
			ret = append(ret, node)
		}
		return true
	})
	return
}
