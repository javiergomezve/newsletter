package entities

import "io"

type Media struct {
	ID          string       `json:"id"`
	FileName    string       `json:"file_name"`
	ContentType string       `json:"content_type"`
	Location    string       `json:"location"`
	Content     io.Reader    `json:"-"`
	Newsletters []Newsletter `json:"-"`
}
