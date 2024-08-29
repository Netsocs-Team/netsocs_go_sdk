package sdk_errors

import (
	"fmt"
	"strings"
)

type NOT_FOUND_ITEMS struct {
	Items   []string
	Domain  string
	Message string
}

func NewNotFoundItemsError(domain string, items []string) *NOT_FOUND_ITEMS {
	return &NOT_FOUND_ITEMS{
		Items:  items,
		Domain: domain,
	}
}

func (e *NOT_FOUND_ITEMS) Error() string {
	items := strings.Join(e.Items, ", ")
	return fmt.Sprintf("not found %s in %s", items, e.Domain)
}
