package todo

import "errors"

var ErrTaskAlreadyExists = errors.New("task already exists")
var ErrTaskNotFound = errors.New("task not found")
