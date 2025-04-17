package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

	gatorerrs "github.com/benskia/Gator/internal/gatorErrs"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerAgg(s *state, cmd command) error {
	errWrap := gatorerrs.NewErrWrapper("handlerAgg")

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return errWrap("failed to fetch feed", err)
	}
	fmt.Printf("%+v", feed)

	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	errWrap := gatorerrs.NewErrWrapper("fetchFeed")

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, errWrap("failed to create request", err)
	}
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errWrap("failed to do request", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || 299 < res.StatusCode {
		return nil, errWrap("", fmt.Errorf("bad HTTP status code: %v", res.StatusCode))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errWrap("failed to read response body", err)
	}

	feed := &RSSFeed{}
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, errWrap("failed to unmarshal data", err)
	}
	escapeFeedStrings(feed)

	return feed, nil
}

func escapeFeedStrings(feed *RSSFeed) {
	feed.Channel.Title = html.EscapeString(feed.Channel.Title)
	feed.Channel.Description = html.EscapeString(feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		item.Title = html.EscapeString(item.Title)
		item.Description = html.EscapeString(item.Description)
	}
}
