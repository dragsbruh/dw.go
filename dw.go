package hook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Field struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

type Thumbnail struct {
	Url string `json:"url,omitempty"`
}

type Footer struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type Embed struct {
	Title       string     `json:"title,omitempty"`
	Url         string     `json:"url,omitempty"`
	Description string     `json:"description,omitempty"`
	Color       int        `json:"color,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Fields      []Field    `json:"fields,omitempty"`
	Timestamp   *time.Time `json:"timestamp,omitempty"`
	Author      *Author    `json:"author,omitempty"`
}

type Author struct {
	Name    string `json:"name,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
	Url     string `json:"url,omitempty"`
}

type Attachment struct {
	Id          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Filename    string `json:"filename,omitempty"`
}

type Hook struct {
	Username    string       `json:"username,omitempty"`
	AvatarUrl   string       `json:"avatar_url,omitempty"`
	Content     string       `json:"content,omitempty"`
	Embeds      []Embed      `json:"embeds,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type RateLimitError struct {
	Message string
}

func (e *RateLimitError) Error() string {
	return e.Message
}

func ExecuteWebhook(link string, hook *Hook) error {
	data, err := json.Marshal(hook)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", link, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("%s", string(bodyText))
	}
	if resp.StatusCode == 429 {
		return &RateLimitError{Message: "rate limit reached"}
	}

	return err
}
