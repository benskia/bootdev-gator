package main

import (
	"context"
	"errors"
	"time"

	"github.com/benskia/Gator/internal/database"
	gatorerrs "github.com/benskia/Gator/internal/gatorErrs"
	"github.com/google/uuid"
)

// handlerFollow() creates a feed_follow relationship between the current user
// and the feed matching the provided URL.
func handlerFollow(s *state, cmd command) error {
	errWrap := gatorerrs.NewErrWrapper("handlerFollow")

	if len(cmd.Args) == 0 {
		return errors.New("usage: follow <url>")
	}

	user, err := s.db.GetUserByName(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errWrap("failed GetUserByName query", err)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return errWrap("failed GetFeedByURL", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return errWrap("failed CreateFeed query", err)
	}

	return nil
}

// handlerFollowing() prints the names of all feeds the current user is following.
func handlerFollowing(s *state, cmd command) error {
	return nil
}
