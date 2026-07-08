package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/phasesoftware/slug"
)

// ANSI coloured output modes
const (
	ansiReset          = "\033[0m"
	ansiFaint          = "\033[2m"
	ansiResetFaint     = "\033[22m"
	ansiBrightRed      = "\033[91m"
	ansiBrightGreen    = "\033[92m"
	ansiBrightYellow   = "\033[93m"
	ansiBrightRedFaint = "\033[91;2m"
)

func main() {
	german := flag.Bool("de", false, "transliterate umlauts and ß as digraphs (ä→ae, ö→oe, ü→ue, ß→ss)")
	flag.Parse()

	style := slug.Default
	if *german {
		style = slug.German
	}

	inputs := flag.Args()

	total := len(inputs)
	renamed := 0

	// Iterate over each input and treat as file path
	for _, input := range inputs {
		// Check if valid file or directory
		info, err := os.Stat(input)
		if err != nil {
			reason := errors.Unwrap(err)
			fmt.Printf(" %sx%s %s %s(%v)%s\n", ansiBrightRed, ansiReset, input, ansiBrightRedFaint, reason, ansiReset)
			continue
		}

		// Extract base name and extension, then slugify
		extension := path.Ext(input)
		pathTo, name := path.Split(input)
		name = strings.TrimSuffix(name, extension)

		result := style.Ify(name)
		if !info.IsDir() {
			result = result + extension
		}

		// Rename file or directory
		err = os.Rename(input, pathTo+result)

		if err == nil {
			fmt.Printf(" %s✓%s %s %s(%s)%s\n", ansiBrightGreen, ansiReset, result, ansiFaint, input, ansiReset)
			renamed += 1
			continue
		}

		fmt.Printf(" %sx%s %s %s(%s)%s\n", ansiBrightRed, ansiReset, input, ansiBrightRedFaint, err.Error(), ansiReset)
	}

	fmt.Printf("%sRenamed %d/%d.%s\n", ansiBrightYellow, renamed, total, ansiReset)
}
