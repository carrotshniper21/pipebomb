// pipebomb/profiles/models.go
package profiles

// swagger:model
type User struct {
	Username string  `json:"username"`
	Avatar   string  `json:"avatar"`
	Profile  Profile `json:"profile"`
}

type Profile struct {
	Name         string        `json:"name"`
	Image        string        `json:"image"`
	Bio          string        `json:"bio"`
	Philosophy   string        `json:"philosophy"`
	Achievements []Achievement `json:"achievements"`
	SocialLinks  []Social      `json:"socialLinks"`
}

type Social struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	URL  string `json:"url"`
}

type Achievement struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Users []User
