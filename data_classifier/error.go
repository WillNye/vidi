package data_classifier

type LoadClassifierError struct {
	ClassifierType string
	msg            string
}

func (e *LoadClassifierError) Error() string {
	return e.msg
}
