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
func ReadFiles(fn func(*os.File) error, files ...string) error {
	for _, file := range files {
		if err := ReadFile(file, fn); err != nil {
			return err
		}
	}
	return nil
}
func ReadFile(name string, fn func(*os.File) error) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return fn(file)
}
