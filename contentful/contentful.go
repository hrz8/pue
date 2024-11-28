package contentful

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ContentfulResponse[T any, U any, V any] struct {
	Sys   Sys `json:"sys,omitempty"`
	Total int `json:"total,omitempty"`
	Skip  int `json:"skip,omitempty"`
	Limit int `json:"limit,omitempty"`
	Items []struct {
		Fields T `json:"fields,omitempty"`
		Sys    struct {
			ID string `json:"id"`
		} `json:"sys,omitempty"`
	} `json:"items,omitempty"`
	Includes Include[U, V] `json:"includes,omitempty"`
}

type Sys struct {
	Type     string `json:"type,omitempty"`
	LinkType string `json:"linkType,omitempty"`
	ID       string `json:"id,omitempty"`
}

type Price struct {
	Price    int    `json:"price,omitempty"`
	Symbol   string `json:"symbol,omitempty"`
	Currency string `json:"currency,omitempty"`
}

type SysEntry struct {
	Sys Sys `json:"sys,omitempty"`
}

type Content struct {
	Data     struct{}       `json:"data,omitempty"`
	Content  []ContentChild `json:"content,omitempty"`
	NodeType string         `json:"nodeType,omitempty"`
}

type ContentChild struct {
	Data     struct{} `json:"data,omitempty"`
	Marks    []any    `json:"marks,omitempty"`
	Value    string   `json:"value,omitempty"`
	NodeType string   `json:"nodeType,omitempty"`
}

type RichText struct {
	Data     struct{}  `json:"data,omitempty"`
	Content  []Content `json:"content,omitempty"`
	NodeType string    `json:"nodeType,omitempty"`
}

type Metadata struct {
	Tags     []any `json:"tags,omitempty"`
	Concepts []any `json:"concepts,omitempty"`
}

type SysFull struct {
	Space            SysEntry  `json:"space,omitempty"`
	ID               string    `json:"id,omitempty"`
	Type             string    `json:"type,omitempty"`
	CreatedAt        time.Time `json:"createdAt,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
	Environment      SysEntry  `json:"environment,omitempty"`
	PublishedVersion int       `json:"publishedVersion,omitempty"`
	Revision         int       `json:"revision,omitempty"`
	ContentType      SysEntry  `json:"contentType,omitempty"`
}

type Entry[T any] struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Sys      SysFull  `json:"sys,omitempty"`
	Fields   T        `json:"fields,omitempty"`
}

type ImageFile struct {
	URL     string `json:"url,omitempty"`
	Details struct {
		Size  int `json:"size,omitempty"`
		Image struct {
			Width  int `json:"width,omitempty"`
			Height int `json:"height,omitempty"`
		} `json:"image,omitempty"`
	} `json:"details,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
}

type Asset[T any] struct {
	Metadata Metadata `json:"metadata,omitempty"`
	Sys      SysFull  `json:"sys,omitempty"`
	Fields   T        `json:"fields,omitempty"`
}

type Include[T any, U any] struct {
	Entry []Entry[T] `json:"Entry,omitempty"`
	Asset []Asset[U] `json:"Asset,omitempty"`
}

type Room struct {
	ID          string   `json:"id,omitempty"`
	RoomType    string   `json:"roomType,omitempty"`
	Hotel       Hotel    `json:"hotel,omitempty"`
	Brand       Brand    `json:"brand,omitempty"`
	City        string   `json:"city,omitempty"`
	Prices      []Price  `json:"prices,omitempty"`
	MaxAdults   int      `json:"maxAdults,omitempty"`
	MaxChildren int      `json:"maxChildren,omitempty"`
	Images      []string `json:"images,omitempty"`
	Area        int      `json:"area,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Brand struct {
	Slug        string `json:"slug,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type GeoCoordinate struct {
	Lat float64 `json:"lat,omitempty"`
	Lon float64 `json:"lon,omitempty"`
}

type Hotel struct {
	Slug        string        `json:"slug,omitempty"`
	Name        string        `json:"name,omitempty"`
	Brand       Brand         `json:"brand,omitempty"`
	City        string        `json:"city,omitempty"`
	Geo         GeoCoordinate `json:"geo,omitempty"`
	Address     string        `json:"address,omitempty"`
	Website     string        `json:"website,omitempty"`
	Description string        `json:"description,omitempty"`
}

type Contentful struct {
	cache map[string][]byte
}

func NewContentful() *Contentful {
	return &Contentful{
		cache: make(map[string][]byte),
	}
}

func (c *Contentful) FetchRoomsByHotelID(hotelID, lang string) ([]byte, error) {
	cacheKey := hotelID + ":" + lang
	if cache, exists := c.cache[cacheKey]; exists {
		return cache, nil
	}

	var body []byte
	url := fmt.Sprintf("%s/spaces/%s/environments/%s/entries?content_type=room&select=fields,sys.id&locale=%s&include=10&fields.hotel.sys.id=%s", ContentfulBaseURL, SpaceID, EnvID, lang, hotelID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return body, fmt.Errorf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return body, fmt.Errorf("Error reading response: %v", err)
	}

	c.cache[cacheKey] = body
	return body, nil
}

func (c *Contentful) ConvertRawToRoom(raw []byte) ([]Room, error) {
	rooms := make([]Room, 0)

	var rawResp ContentfulResponse[RoomField, ReferenceField, AssetImage]
	if err := json.Unmarshal(raw, &rawResp); err != nil {
		return rooms, fmt.Errorf("Error parsing response: %v", err)
	}

	brandMap := make(map[string]Brand)
	for _, entry := range rawResp.Includes.Entry {
		if entry.Sys.ContentType.Sys.ID == "brand" {
			brandMap[entry.Sys.ID] = Brand{
				Slug:        entry.Fields.Slug,
				Name:        entry.Fields.Name,
				Description: entry.Fields.Description.Content[0].Content[0].Value,
			}
		}
	}

	hotelMap := make(map[string]Hotel)
	for _, entry := range rawResp.Includes.Entry {
		if entry.Sys.ContentType.Sys.ID == "hotel" {
			hotelMap[entry.Sys.ID] = Hotel{
				Slug:  entry.Fields.Slug,
				Name:  entry.Fields.Name,
				Brand: brandMap[entry.Fields.Brand.Sys.ID],
				City:  entry.Fields.City,
				Geo: GeoCoordinate{
					Lat: entry.Fields.Geo.Lat,
					Lon: entry.Fields.Geo.Lon,
				},
				Address:     entry.Fields.Address.Content[0].Content[0].Value,
				Website:     entry.Fields.Website,
				Description: entry.Fields.Description.Content[0].Content[0].Value,
			}
		}
	}

	assetMap := make(map[string]string)
	for _, asset := range rawResp.Includes.Asset {
		assetMap[asset.Sys.ID] = asset.Fields.File.URL
	}

	for _, room := range rawResp.Items {
		hotelID := room.Fields.Hotel.Sys.ID
		hotelData, exists := hotelMap[hotelID]
		if !exists {
			continue
		}

		imageUrls := make([]string, 0)
		for _, image := range room.Fields.Images {
			if url, ok := assetMap[image.Sys.ID]; ok {
				imageUrls = append(imageUrls, "https:"+url)
			}
		}

		rooms = append(rooms, Room{
			ID:          room.Sys.ID,
			RoomType:    room.Fields.RoomType,
			Hotel:       hotelData,
			Brand:       hotelData.Brand,
			City:        hotelData.City,
			Prices:      room.Fields.BasePrice,
			MaxAdults:   room.Fields.MaxAdults,
			MaxChildren: room.Fields.MaxChildren,
			Images:      imageUrls,
			Area:        room.Fields.Area,
			Description: room.Fields.Description.Content[0].Content[0].Value,
		})
	}

	return rooms, nil
}
