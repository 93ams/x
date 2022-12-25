package x

import (
	"errors"
	"fmt"
	"github.com/samber/lo"
	"strings"
)

func MultiErrf(e []error, m string, a ...any) error { return MultiErr(e, fmt.Sprintf(m, a...)) }
func MultiErr(errs []error, msg ...string) error {
	if err := strings.Join(lo.Map(lo.Filter(errs, func(item error, _ int) bool { return item != nil }),
		func(item error, _ int) string { return item.Error() }), "\n"); err != "" {
		return errors.New(strings.Join(append(msg, err), "\n"))
	}
	return nil
}
