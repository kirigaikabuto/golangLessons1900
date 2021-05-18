package start

import "errors"

var (
	ErrUserAlreadyExist      = errors.New("User with that username already exist")
	ErrUsernamePasswordEmpty = errors.New("Fields username or password is empty")
)
