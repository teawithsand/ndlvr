package ndlvr

import (
	"fmt"
)

type ErrorBag struct {
	Errors []error
}

func (e *ErrorBag) IsEmpty() bool {
	return e == nil || len(e.Errors) == 0
}

func (e *ErrorBag) Error() string {
	if e == nil {
		return "<nil>"
	}

	return fmt.Sprintf("livr: error bag with %d errors", len(e.Errors))
}

func MakeLIVRError(msg, code string) error {
	return fmt.Errorf("livr: [%s] %s", code, msg)
}
