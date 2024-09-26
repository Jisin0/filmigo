package justwatchcmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Jisin0/filmigo/justwatch"
	"github.com/spf13/cobra"
)

var (
	noTitlesWithoutURL bool
	limit              int
	searchCmd          = &cobra.Command{
		Use:   "search",
		Short: "Search For a Movie or Show on Justwatch",
		Long:  `Search Movies and Shows Using Their Name or ID.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runSearch,
	}
)

func init() {
	searchCmd.Flags().BoolVar(&noTitlesWithoutURL, "no-titles-without-url", false, "dont return titles without a url")
	searchCmd.Flags().IntVar(&limit, "limit", 0, "maxmimum results to return")
}

func runSearch(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	input := args[0]

	client := justwatch.NewClient()

	result, err := client.SearchTitle(input, &justwatch.SearchOptions{Limit: limit, Country: country, Language: language, NoTitlesWithoutURL: noTitlesWithoutURL})
	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}
