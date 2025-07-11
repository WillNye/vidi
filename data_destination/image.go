package data_destination

import (
	"awesomeProject/core"
	"gocv.io/x/gocv"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
)

type FileImage struct {
	Path string
}

func (f *FileImage) DestinationType() string {
	return "Path"
}

func (f *FileImage) DestinationId() string {
	return f.Path
}

func (f *FileImage) ContentType() string {
	return "jpeg"
}

func (f *FileImage) Setup() error {
	return nil
}

func (f *FileImage) Process(content core.Content) error {
	destinationPath := filepath.Join(
		f.Path,
		normalizeSourceId(content.GetSourceId()),
	)
	err := os.MkdirAll(destinationPath, 0755)
	if err != nil {
		slog.Error(
			err.Error(),
			slog.String("destination_type", f.DestinationType()),
			slog.String("destination_id", f.DestinationId()),
			slog.String("source_id", content.GetSourceId()),
		)
		return err
	}

	ok := gocv.IMWrite(
		filepath.Join(destinationPath, strconv.Itoa(int(content.GetId()))+".jpeg"),
		content.GetImage(),
	)
	if !ok {
		slog.Warn(
			"Unable to write image",
			slog.String("destination_type", f.DestinationType()),
			slog.String("destination_id", f.DestinationId()),
			slog.String("source_id", content.GetSourceId()),
			slog.Uint64("content_id", content.GetId()),
		)
	}

	return nil
}

func NewFileImage(path string) *FileImage {
	return &FileImage{path}
}
