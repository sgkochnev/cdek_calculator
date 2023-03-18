package cdekcalculator

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Request struct {
	Type     string  `json:"type"`
	DateTime string  `json:"date_time"`
	State    string  `json:"state"`
	Errors   []Error `json:"errors,omitempty"`
}

type ErrorResp struct {
	Requests []Request `jspn:"requests,omitempty"`
}

type CDEKResponse struct {
	PriceSending []PriceSending `json:"tariff_codes"`
}

type PriceSending struct {
	TariffCode        int     `json:"tariff_code"`
	TariffName        string  `json:"tariff_name"`
	TariffDescription string  `json:"tariff_description"`
	DeliveryMode      int     `json:"delivery_mode"`
	DeliverySum       float64 `json:"delivery_sum"`
	PeriodMin         int     `json:"period_min"`
	PeriodMax         int     `json:"period_max"`
	CalendarMin       int     `json:"calendar_min,omitempty"`
	CalendarMax       int     `json:"calendar_max,omitempty"`
}

type Location struct {
	Code        int    `json:"code,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address,omitempty"`
}

type Size struct {
	Weight int `json:"weight"`
	Length int `json:"length,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type CDEKRequest struct {
	Date         string   `json:"date,omitempty"`
	Type         int      `json:"type,omitempty"`
	Currency     int      `json:"currency,omitempty"`
	Lang         string   `json:"lang,omitempty"`
	FromLocation Location `json:"from_location"`
	ToLocation   Location `json:"to_location"`
	Packages     []Size   `json:"packages"`
}
