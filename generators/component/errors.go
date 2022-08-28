package component

import "fmt"

type InvalidStyleError struct {
	Style Style
}

func (i *InvalidStyleError) Error() string {
	return fmt.Sprintf("invalid component style: %s", i.Style)
}
