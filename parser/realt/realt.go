package realt

type Flat struct {
	Uuid       string   `json:"uuid"`
	Code       int      `json:"code"`
	Headline   string   `json:"headline"`
	Price      int      `json:"price"`
	Images     []string `json:"images"`
	Address    string   `json:"address"`
	Metro      string   `json:"metroStationName"`
	AreaTotal  float64  `json:"areaTotal"`
	AreaLiving float64  `json:"areaLiving"`
	Floor      int      `json:"storey"`
}

func (f Flat) GetID() int {
	return f.Code
}

type ObjectsListing struct {
	IsPending bool   `json:"isPending"`
	Objects   []Flat `json:"objects"`
}

type InitialState struct {
	ObjectsListing ObjectsListing `json:"objectsListing"`
}

type PageProps struct {
	InitialState InitialState `json:"initialState"`
}

type Response struct {
	PageProps PageProps `json:"pageProps"`
}
