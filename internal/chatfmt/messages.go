package chatfmt

import (
	"fmt"
	"strings"

	"github.com/danielmiessler/fabric/internal/chat"
)

func FormatMessages(msgs []*chat.ChatCompletionMessage) string {
	var builder strings.Builder

	for _, msg := range msgs {
		builder.WriteString(FormatMessage(msg))
	}

	return builder.String()
}

func FormatMessage(msg *chat.ChatCompletionMessage) string {
	var builder strings.Builder
	header := roleHeader(msg.Role)

	if len(msg.MultiContent) > 0 {
		builder.WriteString(fmt.Sprintf("%s:\n", header))
		for _, part := range msg.MultiContent {
			builder.WriteString(fmt.Sprintf("  - Type: %s\n", part.Type))
			if part.Type == chat.ChatMessagePartTypeImageURL && part.ImageURL != nil {
				builder.WriteString(fmt.Sprintf("    Image URL: %s\n", part.ImageURL.URL))
				continue
			}
			builder.WriteString(fmt.Sprintf("    Text: %s\n", part.Text))
		}
		builder.WriteString("\n")
		return builder.String()
	}

	builder.WriteString(fmt.Sprintf("%s:\n%s\n\n", header, msg.Content))
	return builder.String()
}

// roleHeader changes the first letter of the role to uppercase, for example "user" to "User".
func roleHeader(role string) string {
	if role == "" {
		return role
	}
	return strings.ToUpper(role[:1]) + role[1:]
}
