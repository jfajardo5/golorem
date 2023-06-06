package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jfajardo5/dolorem"
)

var help = flag.Bool("help", false, "View usage")

var paragraph = flag.Bool("p", false, "Generate paragraphs. Add extra numeric parameters to override default:\n* X (Number of paragraphs)\n* Y (Number of sentences per paragraph)\n* Z (Number of words per sentence)")

var sentence = flag.Bool("s", false, "Generate a sentence. Add an extra numeric parameter to override default:\n* X (Number of words in the sentece)")

var word = flag.Bool("w", false, "Generate a word.")

func main() {
	lorem := dolorem.Ipsum()
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.CommandLine.Args()

	if *paragraph {
		paragraphRouter(lorem, args)
		os.Exit(0)
	}
	if *sentence {
		sentenceRouter(lorem, args)
		os.Exit(0)
	}
	if *word {
		fmt.Println(lorem.Word())
		os.Exit(0)
	}

	fmt.Println(lorem.Paragraph())

}

func sentenceRouter(lorem dolorem.Dolorem, args []string) {
	switch len(args) {
	case 1:
		x, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Arguments must be numeric!")
		}
		fmt.Println(lorem.Sentence(x))
	default:
		fmt.Println(lorem.Sentence())
	}
}

func paragraphRouter(lorem dolorem.Dolorem, args []string) {
	switch len(args) {
	case 1:
		x, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Arguments must be numeric!")
		}
		fmt.Println(lorem.Paragraph(x))
	case 2:
		x, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Arguments must be numeric!")
		}
		y, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Arguments must be numeric!")
		}
		fmt.Println(lorem.Paragraph(x, y))
	case 3:
		x, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Arguments must be numeric!")
		}
		y, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Arguments must be numeric!")
		}
		z, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal("Arguments must be numeric!")
		}
		fmt.Println(lorem.Paragraph(x, y, z))
	default:
		fmt.Println(lorem.Paragraph())
	}
}
