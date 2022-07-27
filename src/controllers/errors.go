package controllers

type Error struct {
	Message string `json:"error"`
}

func getError(err error) *Error {
	return &Error{err.Error()}
}
