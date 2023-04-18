package show

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

type ShowSourcesDecrypted struct {
    Sources []Source `json:"sources"`
    Tracks  []Track `json:"tracks"`
    Server int `json:"server"`
}

type ShowSourcesEncrypted struct {
    Sources string `json:"sources"`
    Tracks  []Track `json:"tracks"`
    Server int `json:"server"`
}

type ShowSource struct {
	Link string `json:"link"`
	Type string `json:"type"`
}

type ShowServer struct {
	ServerName string `json:"serverName"`
	LinkID     string `json:"linkID"`
}

type Episode struct {
	Title string `json:"title"`
	EpisodeID string `json:"episodeID"`
}

type ShowSeason struct {
	SeasonName string `json:"serverName"`
	SeasonID   string `json:"serverID"`
	Episodes []Episode `json:"episodes"`
}

// ShowResponse is the response struct
// @Description is the response struct
type ShowSearchResponse struct {
	Status string      `json:"status" example:"success"`
	Show   *ShowSearch `json:"show"`
}

// ShowStruct stores the show data
// @Description stores the show data
type ShowSearch struct {
	Href        string   `json:"href" example:"https://example.com/show/1"`
	Poster      string   `json:"poster" example:"https://example.com/poster/1.jpg"`
	Title       string   `json:"title" example:"Show"`
	Description string   `json:"description" example:"Description"`
	Duration    string   `json:"duration" example:"0 min"`
	Country     []string `json:"country" example:"country"`
	Production  []string `json:"production" example:"production"`
	Id          string   `json:"id" example:"show/episode-1"`
	IdParts     IdSplit  `json:"idParts"`
	Cast        []string `json:"casts" example:"cast"`
	Genres      []string `json:"genres" example:"genre"`
	Released    string   `json:"released" example:"2000"`
}

// IdSplit stores the show ID parts
// @Description stores the show ID parts
type IdSplit struct {
	Type  string `json:"type" example:"show"`
	Name  string `json:"name" example:"show"`
	IdNum int    `json:"idNum" example:"1"`
}
