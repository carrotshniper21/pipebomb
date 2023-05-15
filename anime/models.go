// pipebomb/anime/models.go
package anime

type SearchInput struct {
	AllowAdult 	 bool	  `json:"allowAdult"`
	AllowUnknown bool   `json:"allowUnknown"`
	Query        string `json:"query"`
}

type AnimeSource struct {
    Data struct {
        Episode struct {
            EpisodeString string `json:"episodeString"`
            SourceUrls    []struct {
                SourceURL  string  `json:"sourceUrl"`
                Priority   float64 `json:"priority"`
                SourceName string  `json:"sourceName"`
                Type       string  `json:"type"`
                ClassName  string  `json:"className"`
                StreamerID string  `json:"streamerId"`
                Downloads  struct {
                    SourceName  string `json:"sourceName"`
                    DownloadURL string `json:"downloadUrl"`
                } `json:"downloads,omitempty"`
                Mobile struct {
                    SourceName  string `json:"sourceName"`
                    DownloadURL string `json:"downloadUrl"`
                } `json:"mobile,omitempty"`
                Sandbox string `json:"sandbox,omitempty"`
            } `json:"sourceUrls"`
        } `json:"episode"`
    } `json:"data"`
}

type Variables struct {
  ShowId string `json:"showId"`
  EpisodeString string `json:"episodeString"`
	Search       	  SearchInput `json:"search"`
	Limit 			 		int `json:"limit"`
	Page 				 		int `json:"page"`
	TranslationType string `json:"translationType"`
	CountryOrigin 	string `json:"countryOrigin"`
}

type AnimeSearch struct {
	Data struct {
		Shows struct {
			Edges []struct {
				ID          string      `json:"_id"`
				Name        string      `json:"name"`
				EnglishName interface{} `json:"englishName"`
				NativeName  interface{} `json:"nativeName"`
				Thumbnail   string      `json:"thumbnail"`
				AlternateThumbnails []string `json:"alternateThumbnails"`
				Type        interface{} `json:"type"`
				Description string `json:"description"`
				AiredStart  struct {
				} `json:"airedStart"`
				Status string `json:"status"`
				Genres []string `json:"genres"`
				Tags []string `json:"tags"`
				Country string `json:"country"`
				Rating  string `json:"rating"`
				Studios []string `json:"studios"`
				EpisodeDuration   interface{} `json:"episodeDuration"`
				EpisodeCount      interface{} `json:"episodeCount"`
				LastUpdateEnd     string `json:"lastUpdateEnd"`
				AvailableEpisodes struct {
					Sub int `json:"sub"`
					Dub int `json:"dub"`
					Raw int `json:"raw"`
				} `json:"availableEpisodes"`
				Typename string `json:"__typename"`
			} `json:"edges"`
		} `json:"shows"`
	} `json:"data"`
}
