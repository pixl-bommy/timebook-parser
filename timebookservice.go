package main

import (
	"context"
	"timebook/utils"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type TimebookSummary struct {
	// Map of task short to total duration in minutes
	Entries map[string]int
	// Total duration in minutes
	TotalMins int
}

type TimebookService struct{}

func (t *TimebookService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

// Read a file and parse its content to a map of task short to total duration in minutes
func (*TimebookService) ParseFile(filePath string) (TimebookSummary, error) {
	lines, err := utils.LoadFileToStringArray(filePath)
	if err != nil {
		return TimebookSummary{}, err
	}

	taskDurationMap := make(map[string]int)
	totalMins := 0

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
		totalMins += parsedTask.DurationMins
	}

	return TimebookSummary{
		Entries:   taskDurationMap,
		TotalMins: totalMins,
	}, nil
}
