package issue

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/crossbone-magister/timewlib"
)

const AnnotationPrefix = "gitlab:"

type Issue struct {
	Iid         string
	Project     string
	Duration    time.Duration
	Day         int
	Month       time.Month
	Year        int
	StartHour   int
	StartMinute int
	EndHour     int
	EndMinute   int
}

func (i *Issue) Details() string {
	return fmt.Sprintf("Project: %s, IID: %s, Duration: %s, Date: %d-%02d-%02d", i.Project, i.Iid, i.Duration.String(), i.Year, i.Month, i.Day)
}

func (i *Issue) Fingerprint() string {
	return fmt.Sprintf("%s#%s#%d-%02d-%02dT%02d:%02d-%02d:%02d",
		i.Project, i.Iid, i.Year, i.Month, i.Day,
		i.StartHour, i.StartMinute, i.EndHour, i.EndMinute)
}

func NewIssue(interval timewlib.Interval) (*Issue, error) {
	var project = ""
	var iid = ""
	for _, tag := range interval.Tags {
		if ref, ok := strings.CutPrefix(tag, AnnotationPrefix); ok {
			split := strings.Split(ref, "#")
			if len(split) == 2 {
				project = split[0]
				iid = split[1]
			}
		}
	}
	if project != "" && iid != "" {
		var year, month, day = interval.StartDate()
		return &Issue{
			Project:     project,
			Iid:         iid,
			Duration:    interval.Duration(),
			Day:         day,
			Month:       month,
			Year:        year,
			StartHour:   interval.StartHour(),
			StartMinute: interval.StartMinute(),
			EndHour:     interval.EndHour(),
			EndMinute:   interval.EndMinute(),
		}, nil
	} else {
		return nil, errors.New("No gitlab issue found in interval")
	}
}
