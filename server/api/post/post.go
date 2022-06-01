package post

import (
	"fmt"
	"strings"
)

type Post struct {
	Category string `json:"category"`
	Topic    string `json:"topic"`
	Content  []struct {
		Text string `json:"text"`
		Type string `json:"type"`
		Path string `json:"path"`
	} `json:"content"`
	Summary string `json:"summary"`
}

func (m Post) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\nCategory: %v\n", m.Category))
	sb.WriteString(fmt.Sprintf("Title: %v\n", m.Topic))
	for _, cont := range m.Content {
		sb.WriteString(fmt.Sprintf("	Text: %v\n", cont.Text))
		sb.WriteString(fmt.Sprintf("	Type: %v\n", cont.Type))
		sb.WriteString(fmt.Sprintf("	Path: %v\n", cont.Path))
	}
	sb.WriteString(fmt.Sprintf("Summary: %v\n", m.Summary))
	return sb.String()
}