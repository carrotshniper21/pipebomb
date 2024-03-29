// pipebomb/anime/animesources.go
package anime

import (
	"encoding/json"
	"log"

	"github.com/gocolly/colly"
)

func animeSources(showId, translationType, episodeString string) (*AnimeSource, error) {
	var anime AnimeSource
	c := colly.NewCollector()

	searchGql := `
  query ($showId: String!, $translationType: VaildTranslationTypeEnumType!, $episodeString: String!) {
      episode(
          showId: $showId
          translationType: $translationType
          episodeString: $episodeString
      ) {
          episodeString sourceUrls
  	  }
  }
  `

	variables := Variables{
 	    ShowId: showId,
			TranslationType: translationType,
			EpisodeString: episodeString,
 	}

	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		log.Fatalf("Error marshalling variables: %s", err)
	}

	apiReq := AssignUrlValues(c, searchGql, variablesJSON)

	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &anime)
		if err != nil {
			log.Fatalf("Error unmarshalling variables: %s", err)
		}
	})

	if err := c.Visit(root + apiReq); err != nil {
		return nil, err
	}

	return &anime, nil
}

func ProcessSources(showId, translationType, episodeString string) (interface{}, error) {
	var results []*AnimeSource

	anime, err := animeSources(showId, translationType, episodeString)
	if err != nil {
		results = append(results, anime)
	}

	return results, nil
}
