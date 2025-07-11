package core

import "gocv.io/x/gocv"

type Content interface {
	GetId() uint64
	GetSourceId() string
	GetImage() gocv.Mat
	GetErr() error
}
