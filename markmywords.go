package main

import (
	"fmt"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func main() {
	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
}
