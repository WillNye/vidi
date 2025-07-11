package data_classifier

import (
	"awesomeProject/core"
	"github.com/asaskevich/EventBus"
	"gocv.io/x/gocv"
)

type DataClassifier interface {
	ClassifierType() string
	TopicId() string
	SourceId() string
	SourceType() string
	EventBus() EventBus.Bus
	Setup() error
	Process(content core.Content) error
	Close() error
}

type ClassifierContent struct {
	ClassifierType string
	ClassifierId   string
	SourceId       string
	Id             uint64
	Image          gocv.Mat
	Err            error
}

func (c *ClassifierContent) GetId() uint64 {
	return c.Id
}

func (c *ClassifierContent) GetSourceId() string {
	return c.SourceId
}

func (c *ClassifierContent) GetImage() gocv.Mat {
	return c.Image
}

func (c *ClassifierContent) GetErr() error {
	return c.Err
}
