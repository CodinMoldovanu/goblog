package models

import (
	"encoding/xml"
	"html/template"
	"strings"
	"time"
)

//Post to be used throughout
type Post struct {
	ID      int
	Author  int
	Posted  time.Time
	Updated *time.Time
	Title   string
	Content template.HTML
	Status  string
	URL     string
	Excerpt template.HTML
	Meta    *xml.CharData
}

//CreateExcerpt from blog post
func (p Post) CreateExcerpt() string {
	content := string(p.Content)
	fc := strings.Index(content, "\n")
	excerpt := content[0:fc]
	return excerpt
}
