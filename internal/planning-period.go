package internal

import (
	"time"

	"github.com/google/uuid"
)

type WorkType rune

const (
	DailyTask   WorkType = 'V'
	Meeting     WorkType = 'M'
	PlannedWork WorkType = 'A'
	Planning    WorkType = 'P'
	Events      WorkType = 'E'

	WorkBreak   WorkType = 'B'
	ClosingTime WorkType = 'X'
)

type PlanningPeriod struct {
	Id       uuid.UUID     `json:"id"`
	Start    time.Time     `json:"firstDay"`
	Profile  Profile       `json:"profile"`
	Expected ExpectedHours `json:"expectedHours"`
	Days     []float32     `json:"days"`
	Tasks    []WorkEntry   `json:"tasks"`
}

type WorkEntry struct {
	Id          uuid.UUID  `json:"id"`
	Start       time.Time  `json:"startTime"`
	End         *time.Time `json:"endTime,omitempty"`
	Type        WorkType   `json:"workType"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
}

type Profile struct {
	DailyHours float32 `json:"dailyHours"`
}

type ExpectedHours struct {
	DailyTaskHours   float32 `json:"dailyTasks"`
	MeetingHours     float32 `json:"meetings"`
	PlannedTaskHours float32 `json:"plannedTasks"`

	PreparationHours float32 `json:"preparation"`
	TeamEventHours   float32 `json:"teamEvents"`
}

func NewPlanningPeriod() *PlanningPeriod {
	now := time.Now()

	return &PlanningPeriod{
		Id:       uuid.New(),
		Start:    toTodayZeroOClock(&now),
		Days:     make([]float32, 0),
		Tasks:    make([]WorkEntry, 0),
		Profile:  Profile{},
		Expected: ExpectedHours{},
	}
}

func (pp *PlanningPeriod) AddWorkEntry(entryType WorkType, title string) uuid.UUID {
	now := time.Now().Local()

	workEntry := WorkEntry{
		Id:    uuid.New(),
		Start: now,
		End:   nil,
		Type:  entryType,
		Title: title,
	}

	pp.Tasks = append(pp.Tasks, workEntry)

	return workEntry.Id
}
