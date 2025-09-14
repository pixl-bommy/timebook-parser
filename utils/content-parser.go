package utils

import (
	"log"
	"strconv"
	"strings"
)

type RawTask struct {
	Line      string
	TaskShort string
	StartTime string
	EndTime   string
}

type ParsedTask struct {
	TaskShort    string
	StartTime    string
	EndTime      string
	DurationMins int
}

type ParsedExpection struct {
	Line         string
	TaskShort    string
	DurationMins int
}

// Filter and trim lines for expected content
func FilterAndTrimLines(lines []string) []string {
	filteredLines := make([]string, 0)

	for _, line := range lines {
		// Trim spaces and tabs
		trimmedLine := strings.Trim(line, string([]byte{' ', '\t'}))

		// Ignore empty lines
		if len(trimmedLine) == 0 {
			continue
		}

		// Expected content starts with "- ("
		if len(trimmedLine) < 3 || trimmedLine[0:3] != "- (" {
			continue
		}

		filteredLines = append(filteredLines, trimmedLine)
	}

	return filteredLines
}

// Parse expected line to extract expected task information
// Example line: "> - Task Long A: 178h"
// Example line: "> - Task Long W: 20h"
// Example line: "> - Another Task Long M: 20h"
func ParseExpectionLine(line string) (*ParsedExpection, bool) {
	// if line does not start with "> - ", ignore
	if len(line) < 4 || line[0:4] != "> - " {
		return nil, false
	}

	// Find the position of the colon
	colonIndex := strings.Index(line, ":")
	if colonIndex == -1 {
		return nil, false
	}

	// Task short is the last character before the colon
	taskPart := strings.TrimSpace(line[colonIndex-1 : colonIndex])
	if len(taskPart) == 0 {
		return nil, false
	}

	// Duration is the part after the colon, trim spaces and "h"
	durationPart := strings.TrimSpace(line[colonIndex+1:])
	durationPart = strings.TrimSuffix(durationPart, "h")
	if len(durationPart) == 0 {
		return nil, false
	}

	durationHours, err := strconv.Atoi(durationPart)
	if err != nil || durationHours < 0 {
		return nil, false
	}

	log.Printf("Parsed expection line: taskShort=%s, durationHours=%d", taskPart, durationHours)

	return &ParsedExpection{
		Line:         line,
		TaskShort:    strings.ToUpper(taskPart),
		DurationMins: durationHours * 60,
	}, true
}

// Parse task line to extract task short, start and end time
// Example line: "- (V 1:23 - 4:56) Task description"
// Example line: "- (Mm 1:23 - 4:56) Task description"
// Example line: "- (V 11:23 - 14:56) Task description"
// Example line: "- (V 01:23 - 04:56) Task description"
func ParseTaskLine(line string) (*RawTask, bool) {
	// Find the position of the closing parenthesis
	closeParenIndex := strings.Index(line, ")")
	if closeParenIndex == -1 {
		return nil, false
	}

	// Extract the content within the parentheses
	parenContent := line[3:closeParenIndex] // Skip the initial "- ("

	// Split the content by spaces
	// Should result in 4 parts: [TaskShort, StartTime, "-", EndTime]
	parts := strings.Fields(parenContent)
	if len(parts) < 4 {
		return nil, false
	}

	return &RawTask{
		Line:      line,
		TaskShort: parts[0],
		StartTime: parts[1],
		EndTime:   parts[3],
	}, true
}

// Convert RawTask to ParsedTask
func ConvertRawToParsed(raw *RawTask) (*ParsedTask, bool) {
	if len(raw.TaskShort) == 0 {
		return nil, false
	}
	taskShort := strings.ToUpper(raw.TaskShort)

	startMins, ok1 := parseTimeStringToMins(raw.StartTime)
	endMins, ok2 := parseTimeStringToMins(raw.EndTime)
	if !ok1 || !ok2 {
		return nil, false
	}

	durationMins := endMins - startMins
	if durationMins < 0 {
		durationMins = durationMins * -1
	}

	return &ParsedTask{
		TaskShort:    taskShort[:1],
		StartTime:    raw.StartTime,
		EndTime:      raw.EndTime,
		DurationMins: durationMins,
	}, true
}

func parseTimeStringToMins(timeStr string) (int, bool) {
	// Split the time string by ":"
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return 0, false
	}

	hours, err1 := strconv.Atoi(parts[0])
	minutes, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || hours < 0 || minutes < 0 {
		return 0, false
	}

	return hours*60 + minutes, true
}
