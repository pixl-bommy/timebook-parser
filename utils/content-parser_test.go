package utils

import (
	"reflect"
	"testing"
)

func TestFilterAndTrimLines(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Empty input",
			input:    []string{},
			expected: []string{},
		},
		{
			name: "All empty lines",
			input: []string{
				"",
				"   ",
				"\t\t",
			},
			expected: []string{},
		},
		{
			name: "Lines without expected prefix",
			input: []string{
				"Hello world",
				"(test)",
			},
			expected: []string{},
		},
		{
			name: "Lines with expected prefix and spaces",
			input: []string{
				"   - (Task 1) something",
				"\t- (Task 2) another",
				" - (Task 3) ",
				"- (Task 4)done",
			},
			expected: []string{
				"- (Task 1) something",
				"- (Task 2) another",
				"- (Task 3)",
				"- (Task 4)done",
			},
		},
		{
			name: "Mixed valid and invalid lines",
			input: []string{
				"   - (Valid 1)",
				"invalid",
				"",
				"\t- (Valid 2) ",
				"  not valid",
				"- (Valid 3)",
			},
			expected: []string{
				"- (Valid 1)",
				"- (Valid 2)",
				"- (Valid 3)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterAndTrimLines(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FilterAndTrimLines(%v) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected *RawTask
		ok       bool
	}{
		{
			name:  "Valid line with V short",
			input: "- (V 1:23 - 4:56) Task description",
			expected: &RawTask{
				Line:      "- (V 1:23 - 4:56) Task description",
				TaskShort: "V",
				StartTime: "1:23",
				EndTime:   "4:56",
			},
			ok: true,
		},
		{
			name:  "Valid line with Mm short and double digit times",
			input: "- (Mm 11:23 - 14:56) Task description",
			expected: &RawTask{
				Line:      "- (Mm 11:23 - 14:56) Task description",
				TaskShort: "Mm",
				StartTime: "11:23",
				EndTime:   "14:56",
			},
			ok: true,
		},
		{
			name:  "Valid line with leading zeros",
			input: "- (V 01:23 - 04:56) Task description",
			expected: &RawTask{
				Line:      "- (V 01:23 - 04:56) Task description",
				TaskShort: "V",
				StartTime: "01:23",
				EndTime:   "04:56",
			},
			ok: true,
		},
		{
			name:     "Missing closing parenthesis",
			input:    "- (V 1:23 - 4:56 Task description",
			expected: nil,
			ok:       false,
		},
		{
			name:     "Too few fields in paren",
			input:    "- (V 1:23) Task description",
			expected: nil,
			ok:       false,
		},
		{
			name:     "No paren at all",
			input:    "V 1:23 - 4:56 Task description",
			expected: nil,
			ok:       false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: nil,
			ok:       false,
		},
		{
			name:     "Paren content too short",
			input:    "- (V - ) Task description",
			expected: nil,
			ok:       false,
		},
		{
			name:  "Valid line with extra spaces",
			input: "- (V   1:23   -   4:56 ) Task description",
			expected: &RawTask{
				Line:      "- (V   1:23   -   4:56 ) Task description",
				TaskShort: "V",
				StartTime: "1:23",
				EndTime:   "4:56",
			},
			ok: true,
		},
		{
			name:  "Valid line with no description",
			input: "- (V 1:23 - 4:56)",
			expected: &RawTask{
				Line:      "- (V 1:23 - 4:56)",
				TaskShort: "V",
				StartTime: "1:23",
				EndTime:   "4:56",
			},
			ok: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := ParseTaskLine(tt.input)
			if ok != tt.ok {
				t.Errorf("ParseLine(%q) ok = %v; want %v", tt.input, ok, tt.ok)
			}
			if ok && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseLine(%q) = %+v; want %+v", tt.input, result, tt.expected)
			}
			if !ok && result != nil {
				t.Errorf("ParseLine(%q) = %+v; want nil", tt.input, result)
			}
		})
	}
}

func TestConvertRawToParsed(t *testing.T) {
	tests := []struct {
		name     string
		input    *RawTask
		expected *ParsedTask
		ok       bool
	}{
		{
			name: "Valid single letter TaskShort",
			input: &RawTask{
				Line:      "- (V 1:23 - 4:56) Task description",
				TaskShort: "V",
				StartTime: "1:23",
				EndTime:   "4:56",
			},
			expected: &ParsedTask{
				TaskShort:    "V",
				StartTime:    "1:23",
				EndTime:      "4:56",
				DurationMins: 213,
			},
			ok: true,
		},
		{
			name: "Valid multi-letter TaskShort (should use first letter uppercased)",
			input: &RawTask{
				Line:      "- (mm 11:23 - 14:56) Task description",
				TaskShort: "mm",
				StartTime: "11:23",
				EndTime:   "14:56",
			},
			expected: &ParsedTask{
				TaskShort:    "M",
				StartTime:    "11:23",
				EndTime:      "14:56",
				DurationMins: 213,
			},
			ok: true,
		},
		{
			name: "Valid with leading zeros",
			input: &RawTask{
				Line:      "- (V 01:23 - 04:56) Task description",
				TaskShort: "V",
				StartTime: "01:23",
				EndTime:   "04:56",
			},
			expected: &ParsedTask{
				TaskShort:    "V",
				StartTime:    "01:23",
				EndTime:      "04:56",
				DurationMins: 213,
			},
			ok: true,
		},
		{
			name: "End time before start time (should return positive duration)",
			input: &RawTask{
				Line:      "- (V 14:56 - 11:23) Task description",
				TaskShort: "V",
				StartTime: "14:56",
				EndTime:   "11:23",
			},
			expected: &ParsedTask{
				TaskShort:    "V",
				StartTime:    "14:56",
				EndTime:      "11:23",
				DurationMins: 213,
			},
			ok: true,
		},
		{
			name: "Invalid TaskShort (empty)",
			input: &RawTask{
				Line:      "- ( 1:23 - 4:56) Task description",
				TaskShort: "",
				StartTime: "1:23",
				EndTime:   "4:56",
			},
			expected: nil,
			ok:       false,
		},
		{
			name: "Invalid StartTime",
			input: &RawTask{
				Line:      "- (V xx:23 - 4:56) Task description",
				TaskShort: "V",
				StartTime: "xx:23",
				EndTime:   "4:56",
			},
			expected: nil,
			ok:       false,
		},
		{
			name: "Invalid EndTime",
			input: &RawTask{
				Line:      "- (V 1:23 - yy:56) Task description",
				TaskShort: "V",
				StartTime: "1:23",
				EndTime:   "yy:56",
			},
			expected: nil,
			ok:       false,
		},
		{
			name: "Negative hour or minute in StartTime",
			input: &RawTask{
				Line:      "- (V -1:23 - 4:56) Task description",
				TaskShort: "V",
				StartTime: "-1:23",
				EndTime:   "4:56",
			},
			expected: nil,
			ok:       false,
		},
		{
			name: "Negative hour or minute in EndTime",
			input: &RawTask{
				Line:      "- (V 1:23 - -4:56) Task description",
				TaskShort: "V",
				StartTime: "1:23",
				EndTime:   "-4:56",
			},
			expected: nil,
			ok:       false,
		},
		{
			name: "Malformed StartTime (missing colon)",
			input: &RawTask{
				Line:      "- (V 123 - 4:56) Task description",
				TaskShort: "V",
				StartTime: "123",
				EndTime:   "4:56",
			},
			expected: nil,
			ok:       false,
		},
		{
			name: "Malformed EndTime (missing colon)",
			input: &RawTask{
				Line:      "- (V 1:23 - 456) Task description",
				TaskShort: "V",
				StartTime: "1:23",
				EndTime:   "456",
			},
			expected: nil,
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, ok := ConvertRawToParsed(tt.input)
			if ok != tt.ok {
				t.Errorf("ConvertRawToParsed(%+v) ok = %v; want %v", tt.input, ok, tt.ok)
			}
			if ok && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ConvertRawToParsed(%+v) = %+v; want %+v", tt.input, result, tt.expected)
			}
			if !ok && result != nil {
				t.Errorf("ConvertRawToParsed(%+v) = %+v; want nil", tt.input, result)
			}
		})
	}
}
