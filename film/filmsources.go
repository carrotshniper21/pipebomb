// pipebomb/film/filmsources.go
package film

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"pipebomb/util"

	"github.com/gocolly/colly"
)

func GetFilmSources(serverId string) (*FilmSourcesDecrypted, error) {
	c := colly.NewCollector()
	var filmSources *FilmSourcesEncrypted

	c.OnResponse(func(r *colly.Response) {
		var source FilmSource
		err := json.Unmarshal(r.Body, &source)
		if err != nil {
			return
		}
		response, _ := getStream(source.Link)

		filmSources = response
	})

	err := c.Visit(root + "/ajax/sources/" + serverId)
	if err != nil {
		return nil, err
	}

	decryptedUrl := util.Dechiper(filmSources.Sources)
	var source []Source
	if err = json.Unmarshal(decryptedUrl, &source); err != nil {
		return nil, err
	}

	return &FilmSourcesDecrypted{source, filmSources.Tracks, filmSources.Server}, nil
}

func getStream(url string) (*FilmSourcesEncrypted, error) {
	providerLinkRegex, _ := regexp.Compile(`(https?://[^\s/]+)`)
	embedRegex, _ := regexp.Compile(`embed-(\d+)/([\w-]+)\??`)
	providerLink := providerLinkRegex.FindString(url)
	embedMatches := embedRegex.FindStringSubmatch(url)
	embedType := embedMatches[1]
	embedId := embedMatches[2]

	reqURL := fmt.Sprintf("%s/ajax/embed-%s/getsources?id=%s", providerLink, embedType, embedId)

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body: ", err)
		}
	}(resp.Body)

	var response FilmSourcesEncrypted
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
