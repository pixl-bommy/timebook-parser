package main

// Summary of timebook entries including total minutes
type TimebookSummary struct {
	Entries   []SummaryEntry
	TotalMins int
}

// A summary entry for a specific task
type SummaryEntry struct {
	// The task short code (e.g. "A" for planned work)
	TaskShort TaskShort
	// The full name of the task (e.g. "Planned Work")
	TaskName string
	// The category short code (e.g. "M" for meetings)
	CategoryShort CategoryShort
	// The full name of the category (e.g. "Meetings")
	CategoryName string
	// Number of tasks for this entry
	CountTasks int

	// Minutes expected for this task
	// If zero, no expectation is set
	ExpectedMinutes int
	// Minutes actually received for this task
	ReceivedMinutes int

	// Factor of received minutes to expected minutes
	// If ExpectedMinutes is zero, this will also be zero.
	// NOTE: This is a factor between 0 and 1, not a percentage.
	FactorOfExpected float64
	// Factor of received minutes to total minutes
	// NOTE: This is a factor calculated over all tasks in the timebook.
	// NOTE: This is a factor between 0 and 1, not a percentage.
	FactorOfTotal float64
}

// A task short code (e.g. "A" for planned work)
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
