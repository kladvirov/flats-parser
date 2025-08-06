package realt

import (
	"encoding/json"
	"flats-parser/constants"
	"time"
)

type Flat struct {
	UUID             string    `json:"uuid"`
	Code             int       `json:"code"`
	Title            string    `json:"title"`
	Price            int       `json:"price"`
	PriceCurrency    int       `json:"priceCurrency"`
	AreaTotal        float64   `json:"areaTotal"`
	AreaLiving       float64   `json:"areaLiving"`
	Rooms            int       `json:"rooms"`
	Storey           int       `json:"storey"`
	Storeys          int       `json:"storeys"`
	Images           []string  `json:"images"`
	CreatedAt        time.Time `json:"createdAt"`
	Address          string    `json:"address"`
	DirectionName    string    `json:"directionName"`
	Headline         string    `json:"headline"`
	MetroStationName string    `json:"metroStationName"`
}

func (f Flat) GetID() int {
	return f.Code
}

type SearchBody struct {
	Results []Flat `json:"results"`
}

type SearchObjects struct {
	Body SearchBody `json:"body"`
}

type GraphQLData struct {
	SearchObjects SearchObjects `json:"searchObjects"`
}

type GraphQLResponse struct {
	Data GraphQLData `json:"data"`
}

func BuildFlatsBody() []byte {
	payload := map[string]any{
		"operationName": "searchObjects",
		"variables": map[string]any{
			"data": map[string]any{
				"where": map[string]any{
					"addressV2": []map[string]string{
						{"townUuid": "4cb07174-7b00-11eb-8943-0cc47adabd66"},
					},
					"category": 2,
				},
				"pagination": map[string]int{
					"page":     1,
					"pageSize": 90,
				},
				"sort": []map[string]string{
					{"by": "createdAt", "order": "DESC"},
				},
				"extraFields":       []string{"minPriceAggregation"},
				"isReactAdaptiveUA": false,
			},
		},
		"query": constants.GQL_QUERY,
	}
	b, _ := json.Marshal(payload)
	return b
}
