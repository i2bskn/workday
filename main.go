package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/najeira/jpholiday"
)

const (
	exitCodeSuccess    = 0
	exitCodeNotWorkday = 1
	workdayStr         = "work"
)

var (
	outputStr = flag.Bool("s", false, "output result as string")
)

func main() {
	flag.Parse()

	now := time.Now()
	// now, _ := time.Parse("2006-01-02", "2017-05-08")
	if now.Weekday() == time.Sunday || now.Weekday() == time.Saturday {
		if *outputStr {
			os.Exit(exitCodeSuccess)
		}
		os.Exit(exitCodeNotWorkday)
	}

	if len(jpholiday.Name(now)) > 0 {
		if *outputStr {
			os.Exit(exitCodeSuccess)
		}
		os.Exit(exitCodeNotWorkday)
	}

	if *outputStr {
		fmt.Println(workdayStr)
	}
	os.Exit(exitCodeSuccess)
}
