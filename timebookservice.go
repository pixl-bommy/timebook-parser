package main

import (
	"context"
	"timebook/utils"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type TimebookService struct{}

func (t *TimebookService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

// Read a file and parse its content to a map of task short to total duration in minutes
func (*TimebookService) ParseFile(filePath string) (map[string]int, error) {
	lines, err := utils.LoadFileToStringArray(filePath)
	if err != nil {
		return nil, err
	}

	taskDurationMap := make(map[string]int)

	filteredLines := utils.FilterAndTrimLines(lines)
	for _, line := range filteredLines {
		rawTask, ok := utils.ParseLine(line)
		if !ok {
			continue
		}

		parsedTask, ok := utils.ConvertRawToParsed(rawTask)
		if !ok {
			continue
		}

		taskDurationMap[parsedTask.TaskShort] += parsedTask.DurationMins
	}

	return taskDurationMap, nil
}
