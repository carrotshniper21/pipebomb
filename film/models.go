package film 

type FilmResponse struct {
  Status string `json:"status"`
  Film *FilmStruct `json:"film,omitempty"`
}

type FilmStruct struct {
    Href string `json:"href"`
    Poster string `json:"poster"`
}
