package internal

import (
	"time"
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

type WorkEntry struct {
	Start       time.Time
	Type        WorkType
	Title       string
	Description string
}

type DayContent struct {
	Date             time.Time
	PlannedWorkHours int
	SpecialWorkHours int
	WorkList         []WorkEntry
}

type DayList map[string]DayContent

type PlanningPeriod struct {
	Days DayList
}

func (pp PlanningPeriod) AddWorkEntry(entryType WorkType) (entryKey string, isNewEntry bool) {
	now := time.Now().Local()

	entryKey = ToMapKey(&now)
	isNewEntry = false

	dayEntry, ok := pp.Days[entryKey]
	if !ok {
		today := toTodayZeroOClock(&now)
		dayEntry = DayContent{
			Date:             today,
			PlannedWorkHours: 0,
			SpecialWorkHours: 0,
			WorkList:         make([]WorkEntry, 0, 4),
		}
		isNewEntry = true
	}

	dayEntry.WorkList = append(dayEntry.WorkList, WorkEntry{
		Start: now,
		Type:  entryType,
	})

	pp.Days[entryKey] = dayEntry
	return
}

func ToMapKey(sourceTime *time.Time) string {
	return sourceTime.Local().Format("2026-01-31")
}

func toTodayZeroOClock(sourceTime *time.Time) time.Time {
	return time.Date(sourceTime.Local().Year(), sourceTime.Local().Month(), sourceTime.Local().Day(), 0, 0, 0, 0, time.Local)
}
