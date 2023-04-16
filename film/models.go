package film

type FilmSource struct {
	Link string `json:"link"`
	Type string `json:"type"`
}

type FilmServer struct {
	ServerName string `json:"serverName"`
	LinkID     string `json:"linkID"`
}

type Source struct {
	File    string `json:"file"`
	Type    string `json:"type"`
}

type Track struct {
    File    string `json:"file"`
    Label   string `json:"label,omitempty"`
    Kind    string `json:"kind"`
    Default bool   `json:"default,omitempty"`
}

type FilmSourcesDecrypted struct {
	  Sources []Source `json:"sources"`
    Tracks  []Track `json:"tracks"`
    Server int `json:"server"`
}

type FilmSourcesEncrypted struct {
    Sources string `json:"sources"`
    Tracks  []Track `json:"tracks"`
    Server int `json:"server"`
}

// FilmResponse is the response struct
// @Description is the response struct
type FilmSearchResponse struct {
	Status string      `json:"status" example:"success"`
	Film   *FilmSearch `json:"film"`
}

// FilmStruct stores the film data
// @Description stores the film data
type FilmSearch struct {
	Href        string   `json:"href" example:"https://example.com/film/1"`
	Poster      string   `json:"poster" example:"https://example.com/poster/1.jpg"`
	Title       string   `json:"title" example:"Film"`
	Description string   `json:"description" example:"Description"`
	Duration    string   `json:"duration" example:"0 min"`
	Country     []string `json:"country" example:"country"`
	Production  []string `json:"production" example:"production"`
	Id          string   `json:"id" example:"movie/film-1"`
	IdParts     IdSplit  `json:"idParts"`
	Cast        []string `json:"casts" example:"cast"`
	Genres      []string `json:"genres" example:"genre"`
	Released    string   `json:"released" example:"2000"`
}

// IdSplit stores the film ID parts
// @Description stores the film ID parts
type IdSplit struct {
	Type  string `json:"type" example:"movie"`
	Name  string `json:"name" example:"film"`
	IdNum int    `json:"idNum" example:"1"`
}
