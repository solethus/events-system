package api

import "time"

type Events struct {
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

type Slideshow struct {
	Slideshow SlideshowContent `json:"slideshow"`
}

type SlideshowContent struct {
	Author string  `json:"author"`
	Date   string  `json:"date"`
	Slides []Slide `json:"slides"`
	Title  string  `json:"title"`
}

type Slide struct {
	Title string   `json:"title"`
	Type  string   `json:"type"`
	Items []string `json:"items,omitempty"`
}

type Response struct {
	MessageID string
}

type SignupEvent struct {
	AuthorName string
}
