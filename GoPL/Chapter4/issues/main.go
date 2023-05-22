package main

import (
	"Chapter4/github/github"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	mFlag, yCurFlag, yPastFlag := false, false, false
	var m, yearCur, yearPast []string
	for _, item := range result.Items {
		if !mFlag && time.Now().Year()-item.CreatedAt.Year() < 1 &&
			(time.Now().Month()-item.CreatedAt.Month() < 1 ||
				time.Now().Month()-item.CreatedAt.Month() == 1 &&
					time.Now().Day()-item.CreatedAt.Day() < 1) {
			if !mFlag {
				m = append(m, "\n#Created less than a month ago\n")
				mFlag = true
			}
			m = append(m, fmt.Sprintf("%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title))
		} else if time.Now().Year()-item.CreatedAt.Year() < 1 &&
			(time.Now().Month()-item.CreatedAt.Month() > 1 ||
				time.Now().Month()-item.CreatedAt.Month() == 1 &&
					time.Now().Day()-item.CreatedAt.Day() > 0) {
			if !yCurFlag {
				yCurFlag = true
				yearCur = append(yearCur, "\n#Created less than a year ago\n")
			}
			yearCur = append(yearCur, fmt.Sprintf("%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title))
		} else {
			if !yPastFlag {
				yPastFlag = true
				yearPast = append(yearPast, "\n#Created at least a year ago\n")
			}
			yearPast = append(yearPast, fmt.Sprintf("%-5d %9.9s %55s\n", item.Number, item.User.Login, item.Title))
		}
	}
	for i := 0; i < len(m); i++ {
		fmt.Printf("%s", m[i])
	}
	for _, item := range yearCur {
		fmt.Printf("%s", item)
	}
	for _, item := range yearPast {
		fmt.Printf("%s", item)
	}
}
