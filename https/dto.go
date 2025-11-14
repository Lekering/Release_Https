package https

import (
	"encoding/json"
	"errors"
	"time"
)

type CompleteDTO struct {
	Complete bool
}

type TaskDTO struct {
	Title       string
	Description string
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func NewErrorDTO(err error) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
}

func (t TaskDTO) TaskDTOValidateToCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Description == "" {
		return errors.New("descriprion is emty")
	}
	return nil
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}
