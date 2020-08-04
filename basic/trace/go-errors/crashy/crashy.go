package crashy

import "github.com/go-errors/errors"

func Crash() error {
	return errors.Errorf("this function is supposed to crash")
}
