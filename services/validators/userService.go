package validators

import (
	"example/errors"
	"example/graph/model"
	"fmt"
)

type userService struct {
	Error error
}

func (u userService) PasswordAndConfirmEqual(c model.CreateUserDto) userService {
	if c.Password != c.PasswordConfirm {
		u.Error = errors.ValidationError{Reason: fmt.Sprintf("password '%s' does not equal '%s'", c.Password, c.PasswordConfirm)}
	}

	return u
}

func (u userService) Get() error {
	return u.Error
}
