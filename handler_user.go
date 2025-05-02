package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/azevedofelipe/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Username is required")
	}

	username := cmd.Args[0]

	if _, err := s.db.GetUser(context.Background(), username); err != nil {
		return fmt.Errorf("User doesnt exist")
	}

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("Couldnt set current user: %w", err)
	}

	fmt.Println("User switched succesfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("name is required")
	}

	name := cmd.Args[0]

	args := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	_, err := s.db.CreateUser(context.Background(), args)
	if err != nil {
		return err
	}

	s.cfg.SetUser(name)

	fmt.Printf("User %s was created succesfully\n", name)
	log.Printf("User %s was created\n", name)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("User table reset")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Couldnt list users: %w", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.User {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}

		fmt.Printf("* %s", user.Name)
	}
	return nil
}
