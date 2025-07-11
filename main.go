package main

import (
	"log/slog"
	"time"
	"vidi/data_classifier"
	"vidi/data_destination"
	"vidi/data_source"
)

func main() {
	source := data_source.NewWebcam(0)
	destination := data_destination.NewFileImage("images/test/raw")
	classifierOne := data_classifier.NewCascadeClassifier(
		"examples/data/haarcascade_frontalface_default.xml",
		source,
	)
	classifierDestination := data_destination.NewFileImage("images/test/classified")

	slog.Info("Setting up source")
	err := source.Setup()
	if err != nil {
		panic(err)
	}

	slog.Info("Setting up destination")
	err = destination.Setup()
	if err != nil {
		panic(err)
	}

	slog.Info("Setting up classifier")
	err = classifierOne.Setup()
	if err != nil {
		panic(err)
	}

	slog.Info("Setting up classifier destination")
	err = classifierDestination.Setup()
	if err != nil {
		panic(err)
	}

	err = source.EventBus().Subscribe(source.TopicId(), destination.Process)
	if err != nil {
		panic(err)
	}

	err = source.EventBus().Subscribe(source.TopicId(), classifierOne.Process)
	if err != nil {
		panic(err)
	}

	err = classifierOne.EventBus().Subscribe(classifierOne.TopicId(), classifierDestination.Process)
	if err != nil {
		panic(err)
	}

	slog.Info("Starting webcam")
	go source.Start()

	time.Sleep(time.Second * 3)
	slog.Info("Stopping webcam")
	err = source.Close()
	if err != nil {
		panic(err)
	}
}
