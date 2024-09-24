package imdbcmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Jisin0/filmigo/imdb"
	"github.com/Jisin0/filmigo/internal/id"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get Full Data of a Movie or Person",
		Long:  `Get Movie Using It's imdb/justwatch id or URL.`,
		Args:  cobra.ExactArgs(1),
		RunE:  runGet,
	}
)

func runGet(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("not enough arguments")
	}

	input := args[0]

	idSource, idType := id.RecognizeID(input)
	if idSource != id.IDSourceImdb {
		return errors.New("id or url invalid")
	}

	client := imdb.NewClient(imdb.ImdbClientOpts{DisableCaching: noCache, CacheExpiration: time.Minute * time.Duration(cacheExpiration)})

	switch idType {
	case id.IDTypeTitle:
		movie, err := client.GetMovie(input)
		if err != nil {
			return err
		}

		if outputJson {
			bytes, err := json.MarshalIndent(movie, "", "   ")
			if err != nil {
				return err
			}

			fmt.Println(string(bytes))
			break
		}

		movie.PrettyPrint()
	case id.IDTypePerson:
		person, err := client.GetPerson(input)
		if err != nil {
			return err
		}

		bytes, err := json.MarshalIndent(person, "", "   ")
		if err != nil {
			return err
		}

		fmt.Println(string(bytes))
	default:
		return errors.ErrUnsupported
	}

	return nil
}
