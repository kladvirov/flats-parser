package kufar

type Response struct {
	Total int  `json:"total"`
	Ads   []Ad `json:"ads"`
}

type Ad struct {
	AccountID         string      `json:"account_id"`
	AccountParameters []Parameter `json:"account_parameters"`
	AdID              int         `json:"ad_id"`
	AdLink            string      `json:"ad_link"`
	AdParameters      []Parameter `json:"ad_parameters"`
	Body              *string     `json:"body"`
	BodyShort         string      `json:"body_short"`
	Category          string      `json:"category"`
	CompanyAd         bool        `json:"company_ad"`
	Currency          string      `json:"currency"`
	Images            []Image     `json:"images"`
	IsMine            bool        `json:"is_mine"`
	ListID            int         `json:"list_id"`
	ListTime          string      `json:"list_time"`
	MessageID         string      `json:"message_id"`
	PhoneHidden       bool        `json:"phone_hidden"`
	PriceBYN          string      `json:"price_byn"`
	PriceUSD          string      `json:"price_usd"`
	RemunerationType  string      `json:"remuneration_type"`
}

func (a Ad) GetID() int {
	return a.AdID
}

type Parameter struct {
	PL string      `json:"pl"`
	VL interface{} `json:"vl"`
	P  string      `json:"p"`
	V  interface{} `json:"v"`
	PU string      `json:"pu"`
	G  []GObject   `json:"g,omitempty"`
}

type GObject struct {
	GI int    `json:"gi"`
	GL string `json:"gl"`
	GO int    `json:"go"`
	PO int    `json:"po"`
}

type Image struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}
