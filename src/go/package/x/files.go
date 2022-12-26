package x

import (
	"os"
	"path/filepath"
	"sync"
)

func NewFile(path string, fn func(*os.File) error) error {
	if err := NewDir(filepath.Dir(path)); err != nil {
		return err
	} else if f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666); err != nil {
		return err
	} else if err != nil {
		return err
	} else {
		defer f.Close()
		return fn(f)
	}
}
func ListFiles(path string, recursive bool) (ret []string, err error) {
	items, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	lock := sync.Mutex{}
	return ret, ParallelTry(func(t os.DirEntry) (err error) {
		var values []string
		if !t.IsDir() {
			values = append(values, filepath.Join(path, t.Name()))
		} else if !recursive {
			return nil
		} else if values, err = ListFiles(filepath.Join(path, t.Name()), recursive); err != nil {
			return err
		}
		if len(values) > 0 {
			lock.Lock()
			ret = append(ret, values...)
			lock.Unlock()
		}
		return nil
	}, items)
}
func ReadFiles(fn func(*os.File) error, paths ...string) error {
	for _, file := range paths {
		if err := ReadFile(file, fn); err != nil {
			return err
		}
	}
	return nil
}
func ReadFile(path string, fn func(*os.File) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return fn(file)
}
func DeleteFile(path string) error      { return os.Remove(path) }
func NewDir(path string) error          { return os.MkdirAll(path, 0666) }
func DeleteDir(path string) error       { return os.RemoveAll(path) }
func NewDirs(dirs ...string) error      { return ParallelTry(NewDir, dirs) }
func DeleteDirs(dirs ...string) error   { return ParallelTry(DeleteDir, dirs) }
func DeleteFiles(files ...string) error { return ParallelTry(DeleteFile, files) }
