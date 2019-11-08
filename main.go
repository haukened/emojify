package emojify

import (
	"regexp"

	"github.com/haukened/emojify/internal/emojiMap"
)

func Render(s string) string {
	re := regexp.MustCompile(`:(\w+|([+-]\d)):`)
	return re.ReplaceAllStringFunc(s, emojiMap.RenderUnicodeEmoji)
}

func Resolve(s string) string {
	re := regexp.MustCompile(`:(\w+|([+-]\d)):`)
	return re.ReplaceAllStringFunc(s, emojiMap.EmojiRootLookup)
}
