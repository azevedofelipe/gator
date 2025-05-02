package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Username is required")
	}

	username := cmd.Args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Couldnt set current user: %w", err)
	}

	fmt.Println("User switched succesfully")
	return nil
}
