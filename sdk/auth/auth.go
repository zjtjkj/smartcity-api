package auth

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidValue = errors.New("invalid value")

//IsRoot auth user privilege is admin or other
func IsRoot(root string) (bool, error) {
	switch root {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, ErrInvalidValue
	}
}

// Region auth user's region privilege
func Region(regions string, target uint64) (bool, error) {
	tg := fmt.Sprintf(" %d", target)
	if strings.Contains(regions, tg) {
		return true, nil
	} else {
		return false, nil
	}
}

// Feature auth user's feature privilege
func Feature(features string, target uint64) (bool, error) {
	tg := fmt.Sprintf(" %d", target)
	if strings.Contains(features, tg) {
		return true, nil
	} else {
		return false, nil
	}
}
