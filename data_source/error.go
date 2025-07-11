package data_source

type InvalidSource struct {
	Source string
	msg    string
}

func (e *InvalidSource) Error() string {
	return e.Source + ": " + e.msg
}

type ReadContentError struct {
	Source string
	msg    string
}

func (e *ReadContentError) Error() string {
	return e.Source + ": " + e.msg
}

type SourceConnectionError struct {
	Source string
	msg    string
}

func (e *SourceConnectionError) Error() string {
	return e.Source + ": " + e.msg
}
