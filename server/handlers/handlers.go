package handlers

type H map[string]any

type MyErr struct {
	Msg  string
	Errs []error
}

func Ok(data any) (int, any) {
	return 0, data
}

func Err(msg string, errs ...error) (int, MyErr) {
	return 1, MyErr{msg, errs}
}

func ErrParam(errs ...error) (int, MyErr) {
	return 1, MyErr{"参数不正确", errs}
}
