package internal

type DBError struct {
	msg string
}

func (err *DBError) Error() string {
	return err.msg
}

func NewError(text string) error {
	// construct a new error object, return
	// pointer to the struct initialized with
	// our error mesage
	return &DBError{text}
}
