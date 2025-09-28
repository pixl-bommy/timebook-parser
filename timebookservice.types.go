package main

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
	TaskShort     TaskShort
	TaskName      string
	CategoryShort CategoryShort
	CategoryName  string
	CountTasks    int

	ExpectedMinutes int
	ReceivedMinutes int

	FactorOfExpected float64
	FactorOfTotal    float64
}

type TimebookSummary struct {
	Entries   []SummaryEntry
	TotalMins int
}
