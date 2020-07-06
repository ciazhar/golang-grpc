package util

type Error struct {
	Error string `json:"error"`
}

func NewError(error string) Error {
	return Error{
		Error: error,
	}
}
