package markmywords

import (
	"fmt"
	"strings"

	"github.com/makyo/ansigo"
)

type StoryElement interface {
	String() string

	Type() string
}

type typer struct {
	elementType string
}

func (t *typer) Type() string {
	return t.elementType
}

type textElement struct {
	typer
	text string
}

func (t *textElement) String() string {
	return fmt.Sprintf("%s", ansigo.MaybeApplyWithReset(currentStyle.text, t.text))
}

func Text(content string) *textElement {
	el := &textElement{
		text: content,
	}
	el.elementType = "text"
	return el
}

type autoreadElement struct {
	typer
	wordsPerSecond int
	minDuration    int
	newlines       bool
	parts          []string
}

func (a *autoreadElement) String() string {
	if a.newlines {
		return strings.Join(a.parts, "...\n\n")
	} else {
		return strings.Join(a.parts, "...")
	}
}

func Autoread(wordsPerSecond, minDuration int, newlines bool, parts ...string) *autoreadElement {
	el := &autoreadElement{
		wordsPerSecond: wordsPerSecond,
		minDuration:    minDuration,
		newlines:       newlines,
		parts:          parts,
	}
	el.elementType = "text"
	return el
}

type headerElement textElement

func (h *headerElement) String() string {
	return fmt.Sprintf("### %s\n\n", h.text)
}

func Header(content string) *headerElement {
	el := &headerElement{
		text: content,
	}
	el.elementType = "header"
	return el
}

type replaceElement struct {
	typer
	from  string
	toEls []StoryElement
}

func (r *replaceElement) String() string {
	var body string
	for _, el := range r.toEls {
		body += el.String()
	}
	if strings.Contains(body, "\n") {
		return fmt.Sprintf("~{%s}{\n%s\n}<!-- /%s -->", r.from, body, r.from)
	} else {
		return fmt.Sprintf("~{%s}{%s}", r.from, body)
	}
}

func Replace(from string, toEls ...StoryElement) *replaceElement {
	el := &replaceElement{
		from:  from,
		toEls: toEls,
	}
	el.elementType = "action"
	return el
}

type linkElement struct {
	typer
	text string
	to   string
}

func Link(text, to string) *linkElement {
	el := &linkElement{
		text: text,
		to:   to,
	}
	el.elementType = "action"
	return el
}

func (l *linkElement) String() string {
	return fmt.Sprintf("[%s](%s)", l.text, l.to)
}
