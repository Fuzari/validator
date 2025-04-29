package text

import "errors"

func Length(t string, min, max int) (bool, error) {
	l := len(t)

	if l < min {
		return false, errors.New("not enough characters")
	} else if l > max {
		return false, errors.New("too much characters")
	} else {
		return true, nil
	}
}
