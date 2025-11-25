package issue

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/crossbone-magister/timewlib"
)

const ProjectIssue = "issue:"

type Issue struct {
	Iid      string
	Project  string
	Duration time.Duration
	Day      int
	Month    time.Month
	Year     int
}

func (i *Issue) Details() string {
	return fmt.Sprintf("Project: %s, IID: %s, Duration: %s, Date: %d-%02d-%02d", i.Project, i.Iid, i.Duration.String(), i.Year, i.Month, i.Day)
}

func NewIssue(interval timewlib.Interval) (*Issue, error) {
	var project = ""
	var iid = ""
	for _, tag := range interval.Tags {
		if issue, ok := strings.CutPrefix(tag, ProjectIssue); ok {
			split := strings.Split(issue, "#")
			project = split[0]
			iid = split[1]
		}
	}
	if project != "" && iid != "" {
		var year, month, day = interval.StartDate()
		return &Issue{Project: project, Iid: iid, Duration: interval.Duration(), Day: day, Month: month, Year: year}, nil
	} else {
		return nil, errors.New("No gitlab issue found in interval")
	}
}
