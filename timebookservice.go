package main

import (
	"context"
	"timebook/utils"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type TimebookSummary struct {
	// Map of task short to total duration in minutes
	Entries map[string]struct {
		Value    int
		Expected int
	}
	// Total duration in minutes
	TotalMins int
}

// Map of task short to full task name
var taskTypes map[string]string = map[string]string{
	"A": "Geplante Arbeiten",
	"O": "Ungeplante Arbeiten",
	"D": "Deployments",
	"M": "Meetings",
	"S": "Support",
	"W": "Wartung",
	"V": "Verschiedenes",
}

// Merged map of task short to full task name, to cover company-specific task groups
var mergedTaskTypes map[string]string = map[string]string{
	"A": "Geplante Arbeiten",
	"O": "Ungeplante Arbeiten",
	"M": "Meetings",
	"V": "Verschiedenes",
	"D": "Wartung",
	"S": "Wartung",
	"W": "Wartung",
}

type TimebookService struct{}

func (t *TimebookService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

// Read a file and parse its content to a map of task short to total duration in minutes
func (t *TimebookService) ParseFile(filePath string) (TimebookSummary, error) {
	return t.ParseFileExtended(filePath, taskTypes)
}

// Read a file and parse its content to a map of task short to total duration in minutes, merging some task types
func (t *TimebookService) ParseFileMerged(filePath string) (TimebookSummary, error) {
	return t.ParseFileExtended(filePath, mergedTaskTypes)
}

// Read a file and parse its content to a map of task short to total duration in minutes
func (*TimebookService) ParseFileExtended(filePath string, taskMap map[string]string) (TimebookSummary, error) {
	lines, err := utils.LoadFileToStringArray(filePath)
	if err != nil {
		return TimebookSummary{}, err
	}

	taskDurationMap := make(map[string]struct {
		Value    int
		Expected int
	})
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

		sanitizedTaskShort := taskMap[parsedTask.TaskShort]
		if sanitizedTaskShort == "" {
			sanitizedTaskShort = taskMap["V"]
		}

		// add entry if not exists
		if _, exists := taskDurationMap[sanitizedTaskShort]; !exists {
			taskDurationMap[sanitizedTaskShort] = struct {
				Value    int
				Expected int
			}{
				Value:    parsedTask.DurationMins,
				Expected: 0, // TODO: parse and set
			}
		} else {
			// update existing entry
			entry := taskDurationMap[sanitizedTaskShort]
			entry.Value += parsedTask.DurationMins
			taskDurationMap[sanitizedTaskShort] = entry
		}

		totalMins += parsedTask.DurationMins
	}

	return TimebookSummary{
		Entries:   taskDurationMap,
		TotalMins: totalMins,
	}, nil
}

func (*TimebookService) SelectFile() (string, error) {
	dialog := application.OpenFileDialog()

	dialog.CanChooseFiles(true)
	dialog.CanChooseDirectories(false)
	dialog.ShowHiddenFiles(true)

	dialog.SetTitle("Select Timebook File")
	dialog.AddFilter("Timebook (*.md)", "*.md")
	dialog.AddFilter("All files", "*")

	return dialog.PromptForSingleSelection()
}
