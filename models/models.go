package models

// Basics defines response from base get
type Basics struct {
	Name string `json:"name"`
	Order int `json:"order"`
	Species Species `json:"species"`
	Stats []Stats `json:"stats"`
}

// Species .
type Species struct {
	Name string `json:"base_stat"`
	URL string `json:"url"`
}

// Stats .
type Stats struct {
	BaseStat int `json:"base_stat"`
	Effort int `json:"effort"`
	Stat Stat `json:"stat"`
}

// Stat .
type Stat struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

// Flavor .
type Flavor struct {
	FlavorTextEntries []FlavorTextEntries `json:"flavor_text_entries"`
}

// FlavorTextEntries .
type FlavorTextEntries struct {
	FlavorText string `json:"flavor_text"`
	Language Language `json:"language"`
}

// Language .
type Language struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

// SpriteURLs .
type SpriteURLs struct {
	Sprites Sprites `json:"sprites"`
}

// Sprites .
type Sprites struct {
	BackDefault string `json:"back_default"`
	BackFemale string `json:"back_female"`
	BackShiny string `json:"back_shiny"`
	BackShinyFemale string `json:"back_shiny_female"`
	FrontDefault string `json:"front_default"`
	FrontFemale string `json:"front_female"`
	FrontShiny string `json:"front:shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}
