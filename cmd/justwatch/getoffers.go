package justwatchcmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Jisin0/filmigo/justwatch"
	"github.com/spf13/cobra"
)

var (
	getOffersCmd = &cobra.Command{
		Use:   "getoffers",
		Short: "Get Offers For a Movie or Show",
		Long:  `Get All Available Offers of a Movie or Show Using it's ID.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runGetOffers,
	}
)

func init() {
	getOffersCmd.Flags().IntVar(&maxEpisodes, "maxepisodes", 0, "maximum episodes of a show to return")
}

func runGetOffers(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	input := args[0]

	client := justwatch.NewClient()

	result, err := client.GetTitleOffers(input, &justwatch.GetTitleOptions{Country: country, Language: language, EpisodeMaxLimit: maxEpisodes})
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
