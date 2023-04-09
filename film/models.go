package film

// FilmResponse is the response struct
// @Description is the response struct
type FilmResponse struct {
    Status string      `json:"status" example:"success"`
    Film   *FilmStruct `json:"film"`
}

// FilmStruct stores the film data
// @Description stores the film data
type FilmStruct struct {
    Href    string  `json:"href" example:"https://example.com/film/1"`
    Poster  string  `json:"poster" example:"https://example.com/poster/1.jpg"`
		Title string `json:"title" example:"https://example.com/2.jpg"`
		Description string `json:"description" example:"https://example.com/description"`
		Duration string `json:"duration" example:"https://example.com/duration"`
		Country []string `json:"country" example:"https://example.com/country"`
		Production []string `json:"production" example:"https://example.com/production"`
    Id      string  `json:"id" example:"film-1"`
    IdParts IdSplit `json:"idParts"`
		Cast []string `json:"casts" example:"film-2"`
		Genres []string `json:"genres" example:"film-3"`
		Released string `json:"released" example:"https://example.com/released"`
}

// IdSplit stores the film ID parts
// @Description stores the film ID parts
type IdSplit struct {
    Type  string `json:"type" example:"film"`
    Name  string `json:"name" example:"Film 1"`
    IdNum int    `json:"idNum" example:1`
}

