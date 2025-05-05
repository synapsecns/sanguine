package signoz

import (
	"fmt"
	"strconv"
	"time"
)

// TimePreferenceType is the type of time preference.
type TimePreferenceType string

const (
	// Last5Min is the last 5 minutes.
	Last5Min TimePreferenceType = "LAST_5_MIN"
	// Last30Min is the last 30 minutes.
	Last30Min TimePreferenceType = "LAST_30_MIN"
	// Last1Hr is the last 1 hour.
	Last1Hr TimePreferenceType = "LAST_1_HR"
	// Last3Hr is the last 3 hour.
	Last3Hr TimePreferenceType = "LAST_3_HR"
	// Last15Min is the last 15 minutes.
	Last15Min TimePreferenceType = "LAST_15_MIN"
	// Last6Hr is the last 6 hours.
	Last6Hr TimePreferenceType = "LAST_6_HR"
	// Last1Day is the last 1 day.
	Last1Day TimePreferenceType = "LAST_1_DAY"
	// Last3Days is the last 3 days.
	Last3Days TimePreferenceType = "LAST_3_DAYS"
	// Last1Week is the last 1 week.
	Last1Week TimePreferenceType = "LAST_1_WEEK"
)

// StartEnd is the start and end time.
type StartEnd struct {
	Start string
	End   string
}

// StartEndInt is the start and end time.
type StartEndInt struct {
	Start int64
	End   int64
}

func getMicroSeconds(t time.Time) string {
	return fmt.Sprintf("%d", t.UnixNano())
}

func getMinAgo(minutes int) time.Time {
	return time.Now().Add(-time.Duration(minutes) * time.Minute)
}

func calculateStartAndEndTime(minutes int, endString string) StartEnd {
	agoDate := getMinAgo(minutes)
	agoString := getMicroSeconds(agoDate)

	return StartEnd{
		Start: agoString,
		End:   endString,
	}
}

// GetStartAndEndTimeInt returns the start and end time as integers.
func GetStartAndEndTimeInt(preference TimePreferenceType) (StartEndInt, error) {
	res := GetStartAndEndTime(preference)
	start, err := strconv.Atoi(res.Start)
	if err != nil {
		return StartEndInt{}, fmt.Errorf("error converting start time to int: %w", err)
	}

	end, err := strconv.Atoi(res.End)
	if err != nil {
		return StartEndInt{}, fmt.Errorf("error converting end time to int: %w", err)
	}

	return StartEndInt{
		Start: int64(start),
		End:   int64(end),
	}, nil
}

// GetStartAndEndTime returns the start and end time.
// nolint: cyclop
func GetStartAndEndTime(preference TimePreferenceType) StartEnd {
	end := time.Now()
	endString := getMicroSeconds(end)

	switch preference {
	case Last5Min:
		return calculateStartAndEndTime(5, endString)
	case Last30Min:
		return calculateStartAndEndTime(30, endString)
	case Last1Hr:
		return calculateStartAndEndTime(60, endString)
	case Last15Min:
		return calculateStartAndEndTime(15, endString)
	case Last3Hr:
		return calculateStartAndEndTime(3*60, endString)
	case Last6Hr:
		return calculateStartAndEndTime(6*60, endString)
	case Last1Day:
		return calculateStartAndEndTime(24*60, endString)
	case Last3Days:
		return calculateStartAndEndTime(24*60*3, endString)
	case Last1Week:
		return calculateStartAndEndTime(24*60*7, endString)
	default:
		// TODO: log error
		return calculateStartAndEndTime(30, endString)
	}
}
