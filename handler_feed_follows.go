package main

import (
	"context"
	"fmt"
	"time"

	"github.com/azevedofelipe/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("url is required")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting feed by URL: %v", err)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.User)
	if err != nil {
		return fmt.Errorf("Error getting current user: %v", err)
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("Error following feed: %v", err)
	}

	fmt.Printf("Feed: %s - User: %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil

}

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.User)
	if err != nil {
		return fmt.Errorf("Error getting current user: %v", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("Error getting users followed feeds: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Printf("No followed feeds for user %s\n", currentUser.Name)
		return nil
	}

	fmt.Println("Followed Feeds:")
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
