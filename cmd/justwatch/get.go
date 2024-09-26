package justwatchcmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Jisin0/filmigo/internal/id"
	"github.com/Jisin0/filmigo/justwatch"
	"github.com/spf13/cobra"
)

var (
	maxEpisodes int
	getCmd      = &cobra.Command{
		Use:   "get",
		Short: "Get Full Data of a Movie or Show",
		Long:  `Get Movie Using It's Justwatch id or URL.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runGet,
	}
)

func init() {
	getCmd.Flags().IntVar(&maxEpisodes, "maxepisodes", 0, "maximum episodes of a show to return")
}

func runGet(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	input := args[0]

	client := justwatch.NewClient()

	idSource, idType := id.RecognizeID(input)
	if idSource != id.IDSourceJW {
		return errors.ErrUnsupported
	}

	var (
		result *justwatch.Title
		err    error
	)

	switch idType {
	case id.IDTypeTitle:
		result, err = client.GetTitle(input, &justwatch.GetTitleOptions{Country: country, Language: language, EpisodeMaxLimit: maxEpisodes})
	case id.IDTypeURL:
		var r *justwatch.URLDetails

		r, err = client.GetTitleFromURL(input, &justwatch.GetTitleOptions{Country: country, Language: language, EpisodeMaxLimit: maxEpisodes})
		result = r.Data
	}

	if err != nil {
		return err
	}

	if outputJson {
		bytes, err := json.MarshalIndent(result, "", "   ")
		if err != nil {
			return err
		}

		fmt.Println(string(bytes))
	} else {
		result.PrettyPrint()
	}

	return nil
}
