package film 

type FilmResponse struct {
  Status string `json:"status"`
  Film *FilmStruct `json:"film"`
}

type FilmStruct struct {
  Href string `json:"href"`
  Poster string `json:"poster"`
  Id string `json:"id"`
	IdParts IdSplit `json:"idParts"`
}

type IdSplit struct {
	Type string `json:"type"`
	Name string `json:"name"`
	IdNum int `json:"idNum"`
}
