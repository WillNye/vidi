package data_destination

import (
	"strings"
	"vidi/core"
)

type Destination interface {
	DestinationType() string
	DestinationId() string
	ContentType() string
	Setup() error
	Process(content core.Content) error
}

func normalizeSourceId(sourceId string) string {
	sourceId = strings.ReplaceAll(sourceId, "/", "_")
	sourceId = strings.ReplaceAll(sourceId, ".", "_")
	return sourceId
}
