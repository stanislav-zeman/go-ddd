package entity

import valueobject "github.com/stanislav-zeman/go-ddd/internal/domain/value_object"

type Todo struct {
	ID      int
	Content string
	State   valueobject.State
}
