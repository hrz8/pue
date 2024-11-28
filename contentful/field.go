package contentful

type RoomField struct {
	RoomType    string     `json:"roomType,omitempty"`
	Hotel       SysEntry   `json:"hotel,omitempty"`
	BasePrice   []Price    `json:"basePrice,omitempty"`
	MaxAdults   int        `json:"maxAdults,omitempty"`
	MaxChildren int        `json:"maxChildren,omitempty"`
	Description RichText   `json:"description,omitempty"`
	Area        int        `json:"area,omitempty"`
	Images      []SysEntry `json:"images,omitempty"`
}

type RoomFieldMultiLang struct {
	RoomType    map[string]string     `json:"roomType,omitempty"`
	Hotel       map[string]SysEntry   `json:"hotel,omitempty"`
	BasePrice   map[string][]Price    `json:"basePrice,omitempty"`
	MaxAdults   map[string]int        `json:"maxAdults,omitempty"`
	MaxChildren map[string]int        `json:"maxChildren,omitempty"`
	Description map[string]RichText   `json:"description,omitempty"`
	Area        map[string]int        `json:"area,omitempty"`
	Images      map[string][]SysEntry `json:"images,omitempty"`
}

type BrandField struct {
	Slug        string   `json:"slug,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description RichText `json:"description,omitempty"`
}

type BrandFieldMultiLang struct {
	Slug        map[string]string   `json:"slug,omitempty"`
	Name        map[string]string   `json:"name,omitempty"`
	Description map[string]RichText `json:"description,omitempty"`
}

type HotelField struct {
	Slug  string   `json:"slug,omitempty"`
	Name  string   `json:"name,omitempty"`
	Brand SysEntry `json:"brand,omitempty"`
	City  string   `json:"city,omitempty"`
	Geo   struct {
		Lat float64 `json:"lat,omitempty"`
		Lon float64 `json:"lon,omitempty"`
	} `json:"geo,omitempty"`
	Address     RichText `json:"address,omitempty"`
	Website     string   `json:"website,omitempty"`
	Description RichText `json:"description,omitempty"`
}

type HotelFieldMultiLang struct {
	Slug  map[string]string   `json:"slug,omitempty"`
	Name  map[string]string   `json:"name,omitempty"`
	Brand map[string]SysEntry `json:"brand,omitempty"`
	City  map[string]string   `json:"city,omitempty"`
	Geo   map[string]struct {
		Lat float64 `json:"lat,omitempty"`
		Lon float64 `json:"lon,omitempty"`
	} `json:"geo,omitempty"`
	Address     map[string]RichText `json:"address,omitempty"`
	Website     map[string]string   `json:"website,omitempty"`
	Description map[string]RichText `json:"description,omitempty"`
}

type ReferenceField struct {
	Slug        string   `json:"slug,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description RichText `json:"description,omitempty"`
	Brand       SysEntry `json:"brand,omitempty"`
	City        string   `json:"city,omitempty"`
	Geo         struct {
		Lat float64 `json:"lat,omitempty"`
		Lon float64 `json:"lon,omitempty"`
	} `json:"geo,omitempty"`
	Address RichText `json:"address,omitempty"`
	Website string   `json:"website,omitempty"`
}

type ReferenceFieldMultiLang struct {
	Slug        map[string]string   `json:"slug,omitempty"`
	Name        map[string]string   `json:"name,omitempty"`
	Description map[string]RichText `json:"description,omitempty"`
	Brand       map[string]SysEntry `json:"brand,omitempty"`
	City        map[string]string   `json:"city,omitempty"`
	Geo         map[string]struct {
		Lat float64 `json:"lat,omitempty"`
		Lon float64 `json:"lon,omitempty"`
	}
	Address map[string]RichText `json:"address,omitempty"`
	Website map[string]string   `json:"website,omitempty"`
}

type AssetImage struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	File        ImageFile `json:"file,omitempty"`
}

type AssetImageMultiLang struct {
	Title       map[string]string    `json:"title,omitempty"`
	Description map[string]string    `json:"description,omitempty"`
	File        map[string]ImageFile `json:"file,omitempty"`
}
