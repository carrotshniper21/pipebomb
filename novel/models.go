package novel

type NovelSearch struct {
    Title string `json:"title"`
    Href  string `json:"href"`
    Image string `json:"image"`
    Author      string `json:"author"`
    Genres      string `json:"genres"`
    Status      string `json:"status"`
    Views       string `json:"views"`
    Description string `json:"description"`
}

