package data_classifier

import (
	"github.com/asaskevich/EventBus"
	"gocv.io/x/gocv"
	"image/color"
	"log/slog"
	"vidi/core"
	"vidi/data_source"
)

type CascadeClassifier struct {
	classifier         gocv.CascadeClassifier
	classifierDataPath string
	eventBus           EventBus.Bus
	sourceId           string
	sourceType         string
}

func (c *CascadeClassifier) ClassifierType() string {
	return "cascade"
}

func (c *CascadeClassifier) ClassifierId() string {
	return c.classifierDataPath
}

func (c *CascadeClassifier) TopicId() string {
	return "classifier:" + c.SourceType() + ":" + c.SourceId() + ":" + c.ClassifierType()
}

func (c *CascadeClassifier) SourceId() string {
	return c.sourceId
}

func (c *CascadeClassifier) SourceType() string {
	return c.sourceType
}

func (c *CascadeClassifier) EventBus() EventBus.Bus {
	return c.eventBus
}

func (c *CascadeClassifier) Setup() error {
	if !c.classifier.Load(c.ClassifierId()) {
		errMsg := "Error reading cascade file"
		slog.Error(errMsg, slog.String("classifier_type", c.ClassifierType()))
		return &LoadClassifierError{c.ClassifierType(), errMsg}
	}

	return nil
}

func (c *CascadeClassifier) Process(content core.Content) error {

	img := content.GetImage()
	blue := color.RGBA{0, 0, 255, 0}
	rects := c.classifier.DetectMultiScale(img)

	// draw a rectangle around each face on the original image
	for _, r := range rects {
		err := gocv.Rectangle(&img, r, blue, 3)
		if err != nil {
			slog.Error(
				err.Error(),
				slog.String("classifier_type", c.ClassifierType()),
				slog.String("classifier_id", c.ClassifierId()),
				slog.String("source_id", c.SourceId()),
				slog.String("source_type", c.SourceType()),
				slog.Uint64("content_id", content.GetId()),
			)
			return err
		}
	}

	go c.eventBus.Publish(
		c.TopicId(),
		&ClassifierContent{c.ClassifierType(),
			c.ClassifierId(),
			c.SourceId(),
			content.GetId(),
			img,
			nil,
		},
	)

	return nil
}

func (c *CascadeClassifier) Close() error {
	return c.classifier.Close()
}

func NewCascadeClassifier(dataPath string, source data_source.DataSource) *CascadeClassifier {
	return &CascadeClassifier{
		classifier:         gocv.NewCascadeClassifier(),
		classifierDataPath: dataPath,
		eventBus:           EventBus.New(),
		sourceId:           source.SourceId(),
		sourceType:         source.SourceType(),
	}
}
