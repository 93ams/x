package x

import "os"

func NewFile(name string, fn func(*os.File) error) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	return fn(f)
}
