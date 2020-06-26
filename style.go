package markmywords

import (
	"fmt"

	"github.com/makyo/ansigo"
)

type style struct {
	text          string
	action        string
	displayHeader bool
}

func Style(text, action string, displayHeader bool) *style {
	currentStyle = &style{
		text:          text,
		action:        action,
		displayHeader: displayHeader,
	}
	return currentStyle
}

func (s *style) String() string {
	return fmt.Sprintf("---\ntext: %s\naction: %s\ndisplayHeader: %t\n---\n\n", s.text, s.action, s.displayHeader)
}

func (s *style) Type() string {
	return "style"
}

func (s *style) Text(text string) string {
	return ansigo.MaybeApplyWithReset(s.text, text)
}

func (s *style) Action(text string) string {
	return ansigo.MaybeApplyWithReset(s.action, text)
}

var currentStyle = &style{}
