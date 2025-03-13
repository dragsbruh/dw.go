# dw.go

> NOTE: This is a fork of [bensch777/discord-webhook-golang](https://github.com/bensch777/discord-webhook-golang) with some changes that i felt would make it better.

## Installation

```go
go get github.com/dragsbruh/dw.go
```

## Code Example

```go
package main

import (
	"fmt"
	"time"

	dw "github.com/dragsbruh/dw.go"
)

func main() {
	webhookUrl := "https://discord.com/api/webhooks/....."

	timestamp := time.Now()

	embed := dw.Embed{
		Title:     "Example Webhook",
		Color:     0xffffff,
		Url:       "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
		Timestamp: &timestamp,
		Thumbnail: &dw.Thumbnail{
			Url: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
		},
		Author: &dw.Author{
			Name:    "Author Name",
			IconUrl: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
		},
		Fields: []dw.Field{
			{
				Name:   "Field 1",
				Value:  "Field Value 1",
				Inline: true,
			},
			{
				Name:   "Field 2",
				Value:  "Field Value 2",
				Inline: true,
			},
			{
				Name:   "Field 3",
				Value:  "Field Value 3",
				Inline: false,
			},
		},
		Footer: &dw.Footer{
			Text:    "Footer Text",
			IconUrl: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
		},
	}

	hook := dw.Hook{
		Username:  "Captain Hook",
		AvatarUrl: "https://avatars.githubusercontent.com/u/6016509?s=48&v=4",
		Content:   "Message",
		Embeds:    []dw.Embed{embed},
	}

	err := dw.ExecuteWebhook(webhookUrl, &hook)

	if _, ok := err.(*dw.RateLimitError); ok {
		fmt.Println("rate limit")
	} else if err != nil {
		panic(err)
	}
}
```

## What are the changes?

1. Made many fields optional, such as timestamp, etc.
2. Added a `RateLimitError` type that is returned when the rate limit is hit.
3. Make the `ExecuteWebhook` function take a `Hook` instead.

## Future plans?

Not entirely sure but:

- [ ] Message builder
- [ ] Also get rate limit information
