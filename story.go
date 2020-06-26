package markmywords

import (
	"fmt"
)

type screen struct {
	typer
	id    string
	title string
	els   []StoryElement
}

func (sc *screen) String() string {
	var body string
	for _, el := range sc.els {
		body += el.String()
	}
	return fmt.Sprintf("## %s {#%s .screen}\n\n%s\n\n", sc.title, sc.id, body)
}

type story struct {
	Title     string
	Author    string
	URL       string
	License   string
	ShortDesc string
	LongDesc  string

	screenOrder []string
	screens     map[string]*screen
	vars        map[string]interface{}
}

func New(title, author, url, license, shortDesc, longDesc string) *story {
	Style("", "underline", false)
	return &story{
		Title:     title,
		Author:    author,
		URL:       url,
		License:   license,
		ShortDesc: shortDesc,
		LongDesc:  longDesc,
		screens:   map[string]*screen{},
	}
}

func (s *story) Screen(title, id string, els ...StoryElement) {
	sc := &screen{
		id:    id,
		title: title,
		els:   els,
	}
	sc.elementType = "screen"
	s.screens[id] = sc
	s.screenOrder = append(s.screenOrder, id)
}

func (s *story) String() string {
	var body string
	for _, id := range s.screenOrder {
		body += s.screens[id].String()
	}
	return fmt.Sprintf(`# %s by %s {#markmywords .story}

%s

-----

%s

-----

URL: <%s>

License: %s

%s`, s.Title, s.Author, s.ShortDesc, s.LongDesc, s.URL, s.License, body)
}
