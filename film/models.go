package film

// FilmResponse is the response struct
type FilmResponse struct {
	Status string      `json:"status"`
	Film   *FilmStruct `json:"film"`
}

// FilmStruct stores the film data
type FilmStruct struct {
	Href    string  `json:"href"`
	Poster  string  `json:"poster"`
	Id      string  `json:"id"`
	IdParts IdSplit `json:"idParts"`
}

// IdSplit stores the film ID parts
type IdSplit struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	IdNum int    `json:"idNum"`
}
