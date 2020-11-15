package maps


// Basics parse basic pokemon data
type Basics struct {
	Name string `json:"name"`
	Order int `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort int `json:"effort"`
		Stat struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
}


// Flavor parse flavor text for pokemon
type Flavor struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
}


// SpriteURL parse sprite URLs
type SpriteURL struct {
	Sprites struct {
		BackDefault string `json:"back_default"`
		BackFemale string `json:"back_female"`
		BackShiny string `json:"back_shiny"`
		BackShinyFemale string `json:"back_shiny_female"`
		FrontDefault string `json:"front_default"`
		FrontFemale string `json:"front_female"`
		FrontShiny string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
	} `json:"sprites"`
}


// CharPixel is converted pixel ascii
type CharPixel struct {
	Char byte
	R uint8
	G uint8
	B uint8
	A uint8
}


// ASCIIOptions convert pixel to raw char
type ASCIIOptions struct {
	Pixels []byte
	Reversed bool
	Colored bool
}


// ConvertOptions to convert image to ascii
type ConvertOptions struct {
	Ratio float64
	FixedWidth int
	FixedHeight int
	Colored bool // may only work in terminal?
	Reversed bool
}
