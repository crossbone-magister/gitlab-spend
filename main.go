package main

import (
	"fmt"
	"gitlab-spend/config"
	"gitlab-spend/issue"
	"gitlab-spend/logic"
	"gitlab-spend/output"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/crossbone-magister/timewlib"
)

func main() {
	var client = &http.Client{}
	if raw, err := timewlib.Parse(os.Stdin); err == nil {
		var configuration timewlib.Configuration = raw.Configuration
		timewlib.ExitOnNoData(raw.Intervals, raw.Configuration)
		timewlib.SetupLogging(configuration)
		var extConfig, err = config.New(configuration)
		timewlib.ExitIfError(err)
		if intervals, err := timewlib.Process(raw.Intervals); err == nil {
			var successes = 0
			var errors = 0
			for _, interval := range intervals {
				toRegister, err := issue.NewIssue(interval)
				if err == nil {
					var response, err = logic.RegisterTimeSpent(*toRegister, client, extConfig)
					if err != nil {
						fmt.Println("Error: ", err, " for interval ", interval.String())
						errors++
						continue
					}
					if response.StatusCode != 201 {
						fmt.Println("Error response: ", response.Status, " for interval ", interval.String())
						errors++
						continue
					}
					successes++
					if configuration.IsDebug() {
						log.Println(response)
						all, err := io.ReadAll(response.Body)
						if err != nil {
							log.Println("Error reading response body:", err)
						} else {
							log.Println("Response body: ", string(all))
						}
					}
				} else {
					log.Printf("Skipping interval %s, no issue found\n\n", interval.String())
				}
			}
			err := output.PrintReport(os.Stdout, intervals, successes, errors)
			timewlib.ExitIfError(err)
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		}
	}
}
