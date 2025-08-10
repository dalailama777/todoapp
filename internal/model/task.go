package model

import "errors"

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t *Task) TaskValidation() error {
	if len(t.Title) == 0 {
		return errors.New("empty title")
	}
	return nil
}
