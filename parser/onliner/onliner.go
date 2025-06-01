package onliner

type Apartment struct {
	ID       int    `json:"id"`
	Price    Price  `json:"price"`
	RentType string `json:"rent_type"`
	Location struct {
		Address     string  `json:"address"`
		UserAddress string  `json:"user_address"`
		Latitude    float64 `json:"latitude"`
		Longitude   float64 `json:"longitude"`
	} `json:"location"`
	Photo   string `json:"photo"`
	Contact struct {
		Owner bool `json:"owner"`
	} `json:"contact"`
	CreatedAt  string `json:"created_at"`
	LastTimeUp string `json:"last_time_up"`
	URL        string `json:"url"`
}

func (a Apartment) GetID() int {
	return a.ID
}

type Price struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	Converted struct {
		BYN struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"BYN"`
		USD struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"USD"`
	} `json:"converted"`
}

type Response struct {
	Apartments []Apartment `json:"apartments"`
}
