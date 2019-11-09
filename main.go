/*
Package emojify provides a way to convert emoji shortnames (:beer:) to unicode emojis 🍺 specifically to render in text based terminal applications and UIs.

It is designed to be imported into your codebase, as opposed to a standalone binary:

import (
	"fmt"

	"github.com/haukened/emojify
)

func main() {
	fmt.Println(emojify.Render(":beer:"))
}

It also provides the Resolve(string) function to resolve aliases to root emojis that have unicode representations. Running Resolve(":cheeseburger:") results in ":hamburger:".

In the event emojify cannot render the unicode string (such as multiple emojis combined with zero width joiners (Families, etc...) the original string is returned as to not modify the intent of the message.

*/
package emojify

import (
	"regexp"

	"github.com/haukened/emojify/internal/emojiMap"
)


// Render accepts a string, and if the emoji exists in the map of emojis, the unicode representation is returned, if not (or if it does and cannot be rendered to a single unicode rune) the original string is returned.
func Render(s string) string {
	re := regexp.MustCompile(`:(\w+|([+-]\d)):`)
	return re.ReplaceAllStringFunc(s, emojiMap.RenderUnicodeEmoji)
}

// Resolve accepts a string, and returns the root emoji, in the event that the emoji is an alias.  If the emoji is already a root, is not within the set, or is not an emoji the original string is returned.
func Resolve(s string) string {
	re := regexp.MustCompile(`:(\w+|([+-]\d)):`)
	return re.ReplaceAllStringFunc(s, emojiMap.EmojiRootLookup)
}
