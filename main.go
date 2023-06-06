package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jfajardo5/dolorem"
)

var help = flag.Bool("help", false, "View usage")

var paragraph = flag.Bool("paragraph", false, "Generate paragraphs. Add extra numeric parameters to override default:\n* X (Number of paragraphs)\n* Y (Number of sentences per paragraph)\n* Z (Number of words per sentence)")

var sentence = flag.Bool("sentence", false, "Generate a sentence. Add an extra numeric parameter to override default:\n* X (Number of words in the sentece)")

var word = flag.Bool("word", false, "Generate a word.")

func main() {
	args := os.Args[1:]
	lorem := dolorem.Ipsum()
	flag.Parse()

	if len(args) == 0 {
		fmt.Println(lorem.Paragraph())
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}

}
