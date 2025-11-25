package output

import (
	"fmt"
	"io"

	"github.com/crossbone-magister/timewlib"
)

func PrintReport(output io.Writer, intervals []timewlib.Interval, successes int, errors int) error {
	_, err := fmt.Fprintln(output, "Intervals processed: ", len(intervals))
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(output, "Intervals registered: ", successes)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(output, "Intervals registration failed: ", errors)
	if err != nil {
		return err
	}
	return nil
}
