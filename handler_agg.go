package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Time arg required")
	}
	interval := cmd.Args[0]

	timeBetweenRequests, err := time.ParseDuration(interval)
	if err != nil {
		return fmt.Errorf("Error converting time: %v", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
