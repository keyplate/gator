package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
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
    PubDate     time.Time `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
    req, err := http.NewRequestWithContext(context.Background(), "GET", feedUrl, nil)
    if err != nil {
        return &RSSFeed{}, err
    }
    req.Header.Add("User-Agen", "gator/1.0")
    
    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return &RSSFeed{}, err
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return &RSSFeed{}, err
    }
    
    rssFeed := &RSSFeed{}
    err = xml.Unmarshal(body, rssFeed)
    if err != nil {
        return &RSSFeed{}, err
    }

    return rssFeed, nil
}

func unescapeHTML(feed RSSFeed) RSSFeed {
    feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
    feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
    for i, _ := range feed.Channel.Item {
        feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
        feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
    }
    return feed
}
