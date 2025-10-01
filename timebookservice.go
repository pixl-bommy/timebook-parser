package main

import (
	"context"
	"timebook/utils"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type TimebookService struct {
	// TODO: this is used to cache the last parsed file, to avoid re-parsing it
	// for multiple future interpretations (e.g. use as is, sum per categroy).
	currentTimebookSummary *TimebookSummary
}

func (t *TimebookService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

func (t *TimebookService) LoadFile(filePath string) (TimebookSummary, error) {
	t.currentTimebookSummary = nil
	timebookSummary, err := t.parseFile(filePath)

	if err == nil {
		t.currentTimebookSummary = &timebookSummary
	}
	return timebookSummary, err
}

// Read a file and parse its content to a map of task short to total duration in minutes
func (*TimebookService) parseFile(filePath string) (TimebookSummary, error) {
	lines, err := utils.LoadFileToStringArray(filePath)
	if err != nil {
		return TimebookSummary{}, err
	}

	taskDurationMap := make(map[TaskShort]SummaryEntry)
	totalMins := 0

	// parse each line for expected task information
	for _, line := range lines {
		parsedExpection, ok := utils.ParseExpectionLine(line)
		if !ok {
			continue
		}

		taskShort := newTaskShortFromInput(parsedExpection.TaskShort)

		// update existing entry
		if entry, exists := taskDurationMap[taskShort]; exists {
			entry.ExpectedMinutes += parsedExpection.DurationMins
			taskDurationMap[taskShort] = entry
			continue
		}

		// otherwise create new entry
		newTask := newSummaryEntry(taskShort)
		newTask.ExpectedMinutes = parsedExpection.DurationMins
		taskDurationMap[taskShort] = newTask
	}

	// parse each line for task information
	filteredLines := utils.FilterAndTrimLines(lines)
	for _, line := range filteredLines {
		rawTask, ok := utils.ParseTaskLine(line)
		if !ok {
			continue
		}

		parsedTask, ok := utils.ConvertRawToParsed(rawTask)
		if !ok {
			continue
		}

		taskShort := newTaskShortFromInput(parsedTask.TaskShort)

		// increment total minutes
		totalMins += parsedTask.DurationMins

		// update existing entry
		if entry, exists := taskDurationMap[taskShort]; exists {
			entry.ReceivedMinutes += parsedTask.DurationMins
			entry.CountTasks++
			taskDurationMap[taskShort] = entry
			continue
		}

		// otherwise create new entry
		newTask := newSummaryEntry(taskShort)
		newTask.ReceivedMinutes = parsedTask.DurationMins
		newTask.CountTasks = 1
		taskDurationMap[taskShort] = newTask
	}

	// calculate percentages
	for taskShort, entry := range taskDurationMap {
		if entry.ExpectedMinutes > 0 {
			entry.FactorOfExpected = float64(entry.ReceivedMinutes) / float64(entry.ExpectedMinutes)
		}

		if totalMins > 0 {
			entry.FactorOfTotal = float64(entry.ReceivedMinutes) / float64(totalMins)
		}

		taskDurationMap[taskShort] = entry
	}

	// convert map to slice
	entries := make([]SummaryEntry, 0, len(taskDurationMap))
	for _, entry := range taskDurationMap {
		entries = append(entries, entry)
	}

	timebookSummary := TimebookSummary{
		Entries:   entries,
		TotalMins: totalMins,
	}
	return timebookSummary, nil
}

func newTaskShortFromInput(input string) TaskShort {
	switch input {
	case "A":
		return PlannedWork
	case "O":
		return UnplannedWork
	case "D":
		return Deployments
	case "M":
		return Meetings
	case "S":
		return Support
	case "W":
		return Maintenance
	default:
		return Miscellaneous
	}
}

func newSummaryEntry(taskShort TaskShort) SummaryEntry {
	category := taskShort.Category()

	return SummaryEntry{
		TaskShort:     taskShort,
		TaskName:      taskShort.FullName(),
		CategoryShort: category,
		CategoryName:  category.FullName(),
	}
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
