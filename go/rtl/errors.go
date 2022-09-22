package rtl

type ResponseError struct {
	StatusCode int
	Body       []byte
	Err        error
}

func (e ResponseError) Error() string {
	return e.Err.Error()
}
