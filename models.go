package main

import (
	"time"
)

type shortenRequestURL struct {
	URL          string `json: "url" binding:"required"`
	CustomSlug   string `json: "custom_slug"`
}


type Link struct {
	OriginalURL  string `json: "original_url"`
	ShortCode    string `json: "short_code"`
	CreatedAt    string  `json: :"created_at"`
}

