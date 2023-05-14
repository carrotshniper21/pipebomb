// pipebomb/anime/animesearch.go 
package anime

import ( 
	"bytes"
	"encoding/json"
	"log"
	"io/ioutil"
	"net/url" 

	"github.com/gocolly/colly"
)

const root = "https://api.allanime.to/allanimeapi?"

func animeSearcher(animeQuery string) (*AnimeSearch, error) {
	var anime AnimeSearch
	c := colly.NewCollector()

	searchGql := `
	query(
		$search: SearchInput
		$limit: Int
		$page: Int
		$translationType: VaildTranslationTypeEnumType
    $countryOrigin: VaildCountryOriginEnumType
	) {
		shows(
			search: $search
			limit: $limit
			page: $page
			translationType: $translationType
			countryOrigin: $countryOrigin
		) {
			edges {
				_id name englishName nativeName thumbnail type airedStart episodeDuration episodeCount lastUpdateEnd availableEpisodes __typename
			}
    }
	}
	`

	variables := Variables{
		Search: SearchInput{
			AllowAdult: false,
			AllowUnknown: false,
			Query: animeQuery,
		},
		Limit: 40,
		Page: 1,
		TranslationType: "dub",
		CountryOrigin: "ALL",
	}

	variablesJSON, err := json.Marshal(variables)
	if err != nil {
		log.Fatalf("Error marshalling variables: %s", err)
	}
	
	apiReq := AssignUrlValues(c, searchGql, variablesJSON)

	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &anime)
		if err != nil {
			log.Fatalf("Error deserializing response: %s", err)
		}
	})

	if err := c.Visit(root + apiReq); err != nil {
		return nil, err
	}

	return &anime, nil
}

func AssignUrlValues(c *colly.Collector, searchGql string, variablesJSON []byte) string {
	reqBody := bytes.NewBufferString(url.Values{
			"query":     []string{searchGql},
			"variables": []string{string(variablesJSON)},
	}.Encode())

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")
		r.Headers.Set("Referer", "https://allanime.to")
		r.Body = ioutil.NopCloser(reqBody)
	})

	body := ioutil.NopCloser(reqBody)
	bodyBytes, _ := ioutil.ReadAll(body)
	return string(bodyBytes)
}

// ProcessQuery processes a query and returns a AnimeStruct
func ProcessQuery(query string) *AnimeSearch {
	anime, err := animeSearcher(query)
	if err != nil {
		return nil
	}

	return anime
}

