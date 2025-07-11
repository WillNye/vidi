package data_source

import (
	"github.com/asaskevich/EventBus"
	"gocv.io/x/gocv"
	"log/slog"
	"strconv"
)

type WebcamConfig struct {
	Source int
}

type Webcam struct {
	conn     *gocv.VideoCapture
	config   WebcamConfig
	eventBus EventBus.Bus
	connOpen bool
}

func (cam *Webcam) SourceType() string {
	return "Webcam"
}

func (cam *Webcam) SourceId() string {
	return strconv.Itoa(cam.config.Source)
}

func (cam *Webcam) TopicId() string {
	return "datastream:" + cam.SourceType() + ":" + cam.SourceId()
}

func (cam *Webcam) Setup() error {
	return nil
}

func (cam *Webcam) EventBus() EventBus.Bus {
	return cam.eventBus
}

func (cam *Webcam) Start() {
	cam.connOpen = true

	conn, err := gocv.OpenVideoCapture(cam.config.Source)
	if err != nil {
		slog.Error(
			err.Error(),
			slog.String("source_type", cam.SourceType()),
			slog.String("source_id", cam.SourceId()),
		)
	}
	cam.conn = conn
	defer cam.Close()

	img := gocv.NewMat()
	defer img.Close()

	var id uint64 = 1
	retryCount := 0
	maxRetryCount := 5
	for {
		if !cam.connOpen {
			return
		}

		cam.conn.Read(&img)
		if img.Empty() {
			if retryCount < maxRetryCount {
				retryCount++
				continue
			} else {
				slog.Error(
					"Failed to read image",
					slog.String("source", cam.SourceId()),
				)
				return
			}
		}
		go cam.eventBus.Publish(
			cam.TopicId(),
			&SourceContent{SourceId: cam.SourceId(), Id: id, Image: img, Err: nil},
		)

		id++
		retryCount = 0
	}
}

func (cam *Webcam) Stop() error {
	cam.connOpen = false
	return nil
}

func (cam *Webcam) Close() error {
	return cam.Stop()
}

func NewWebcam(source int) *Webcam {
	return &Webcam{config: WebcamConfig{Source: source}, eventBus: EventBus.New(), connOpen: true}
}
