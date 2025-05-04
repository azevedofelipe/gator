package main

import (
	"context"
	"fmt"
	"time"

	"github.com/azevedofelipe/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("url is required")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting feed by URL: %v", err)
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return fmt.Errorf("Error following feed: %v", err)
	}

	fmt.Printf("Feed: %s - User: %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil

}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error getting users followed feeds: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Printf("No followed feeds for user %s\n", user.Name)
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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("url is required")
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting feed by URL: %v", err)
	}

	args := database.RemoveFeedFollowForUserParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.RemoveFeedFollowForUser(context.Background(), args)
	if err != nil {
		return fmt.Errorf("Error unfollowing feed: %v", err)
	}

	fmt.Printf("Unfollowed %s for user %s\n", url, user.Name)
	return nil
}
