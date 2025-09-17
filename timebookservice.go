package main

import (
	"context"
	"timebook/utils"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type TaskShort string

const (
	PlannedWork   TaskShort = "A"
	UnplannedWork TaskShort = "O"
	Deployments   TaskShort = "D"
	Meetings      TaskShort = "M"
	Support       TaskShort = "S"
	Maintenance   TaskShort = "W"
	Miscellaneous TaskShort = "V"
)

func toTaskShort(input string) TaskShort {
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

func (t TaskShort) FullName() string {
	switch t {
	case PlannedWork:
		return "Geplante Arbeiten"
	case UnplannedWork:
		return "Ungeplante Arbeiten"
	case Deployments:
		return "Deployments"
	case Meetings:
		return "Meetings"
	case Support:
		return "Support"
	case Maintenance:
		return "Wartung"
	case Miscellaneous:
		return "Verschiedenes"
	default:
		return "Unbekannt"
	}
}

func (t TaskShort) Category() CategoryShort {
	switch t {
	case PlannedWork:
		return PlannedWorkCategory
	case UnplannedWork:
		return UnplannedWorkCategory
	case Deployments:
		return MaintenanceCategory
	case Meetings:
		return MeetingsCategory
	case Support:
		return MaintenanceCategory
	case Maintenance:
		return MaintenanceCategory
	case Miscellaneous:
		return MiscellaneousCategory
	default:
		return MiscellaneousCategory
	}
}

type CategoryShort string

const (
	PlannedWorkCategory   CategoryShort = "A"
	UnplannedWorkCategory CategoryShort = "O"
	MeetingsCategory      CategoryShort = "M"
	MaintenanceCategory   CategoryShort = "W"
	MiscellaneousCategory CategoryShort = "V"
)

func (t CategoryShort) FullName() string {
	switch t {
	case PlannedWorkCategory:
		return "Geplante Arbeiten"
	case UnplannedWorkCategory:
		return "Ungeplante Arbeiten"
	case MeetingsCategory:
		return "Meetings"
	case MaintenanceCategory:
		return "Wartung"
	case MiscellaneousCategory:
		return "Verschiedenes"
	default:
		return "Unbekannt"
	}
}

type SummaryEntry struct {
	TaskShort       TaskShort
	TaskName        string
	CategoryShort   CategoryShort
	CategoryName    string
	TotalMinutes    int
	ExpectedMinutes int
	CountTasks      int
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

type TimebookSummary struct {
	Entries   []SummaryEntry
	TotalMins int
}

type TimebookService struct{}

func (t *TimebookService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

// Read a file and parse its content to a map of task short to total duration in minutes
func (t *TimebookService) ParseFile(filePath string) (TimebookSummary, error) {
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

		taskShort := toTaskShort(parsedExpection.TaskShort)

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

		taskShort := toTaskShort(parsedTask.TaskShort)

		// increment total minutes
		totalMins += parsedTask.DurationMins

		// update existing entry
		if entry, exists := taskDurationMap[taskShort]; exists {
			entry.TotalMinutes += parsedTask.DurationMins
			entry.CountTasks++
			taskDurationMap[taskShort] = entry
			continue
		}

		// otherwise create new entry
		newTask := newSummaryEntry(taskShort)
		newTask.TotalMinutes = parsedTask.DurationMins
		newTask.CountTasks = 1
		taskDurationMap[taskShort] = newTask
	}

	// convert map to slice
	entries := make([]SummaryEntry, 0, len(taskDurationMap))
	for _, entry := range taskDurationMap {
		entries = append(entries, entry)
	}

	return TimebookSummary{
		Entries:   entries,
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
