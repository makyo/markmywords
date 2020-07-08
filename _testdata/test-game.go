package main

import (
	"fmt"

	m "github.com/makyo/markmywords"
)

func main() {
	s := m.New(
		"Test game",
		"Maddy Paws",
		"https://github.com/makyo/markmywords/tree/master/_testdata/test-game.go",
		"CC0",
		"A test game for MarkMyWords",
		"Just a game testing various features",
	)
	s.Screen(
		"Introduction",
		"intro",
		m.Text("First there was a "),
		m.Autoread(3, 1, false, "very ", "delightful "),
		m.Link("fox.", "#fox"),
	)
	s.Screen(
		"A fox",
		"fox",
		m.Style("", "underline+red", true),
		m.Header("About the fox"),
		m.Text("It was a very pretty fox\n\n"),
		m.Replace("But also...",
			m.Style("", "underline+red", false),
			m.Text("It was a "),
			m.Replace("smart", m.Text("clever")),
			m.Replace(" fox.", m.Text(" fox. Perhaps too clever.")),
		),
	)

	fmt.Println(s.String())
}
