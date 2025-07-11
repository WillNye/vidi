package data_source

import (
	"github.com/asaskevich/EventBus"
	"gocv.io/x/gocv"
)

type DataSource interface {
	SourceType() string
	SourceId() string
	TopicId() string
	EventBus() EventBus.Bus
	Setup() error
	Start()
	Stop() error
	Close() error
}

type SourceContent struct {
	SourceId string
	// Id is an incrementing value used to determine order
	Id    uint64
	Image gocv.Mat
	Err   error
}

func (c *SourceContent) GetId() uint64 {
	return c.Id
}

func (c *SourceContent) GetSourceId() string {
	return c.SourceId
}

func (c *SourceContent) GetImage() gocv.Mat {
	return c.Image
}

func (c *SourceContent) GetErr() error {
	return c.Err
}
