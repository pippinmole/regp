package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"regexp"
)

func main() {
	app := &cli.App{
		Name: "regp",
		Commands: []*cli.Command{
			{
				Name:    "match",
				Aliases: []string{"m"},
				Usage:   "Match a pattern against a list of strings",
				Action: func(context *cli.Context) error {
					if context.Args().Len() < 1 {
						return errors.New("no arguments provided")
					}

					pattern := context.Args().Get(0)
					others := context.Args().Slice()[1:]

					if len(pattern) == 0 {
						return errors.New("no pattern to match against")
					}

					if len(others) == 0 {
						return errors.New("no strings to match against")
					}

					return match(pattern, others)
				},
			},
		},
		Action: func(cCtx *cli.Context) error {
			// Print help stuff

			fmt.Println("Usage: regp [command] [arguments]")

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func match(pattern string, others []string) error {
	re, err := regexp.Compile(pattern)

	if err != nil {
		return err
	} else {
		for _, other := range others {
			isMatch := re.MatchString(other)

			if isMatch {
				fmt.Printf("✅  Match: '%s'\n", other)
			} else {
				fmt.Printf("❌  No match: '%s'\n", other)
			}
		}
	}

	return nil
}
