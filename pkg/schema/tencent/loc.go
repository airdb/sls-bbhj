package tencent

type Location struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

type AddressComponents struct {
	Province     string `json:"province,omitempty"`
	City         string `json:"city,omitempty"`
	District     string `json:"district,omitempty"`
	Street       string `json:"street,omitempty"`
	StreetNumber string `json:"street_number,omitempty"`
}

type AdInfo struct {
	Adcode string `json:"adcode,omitempty"`
}

type ResultLoc struct {
	Title             string            `json:"title,omitempty"`
	Location          Location          `json:"location,omitempty"`
	AddressComponents AddressComponents `json:"address_components,omitempty"`
	AdInfo            AdInfo            `json:"ad_info,omitempty"`
	Reliability       int               `json:"reliability,omitempty"`
	Level             int               `json:"level,omitempty"`
}
type RespLoc struct {
	Status  int       `json:"status,omitempty"`
	Message string    `json:"message,omitempty"`
	Result  ResultLoc `json:"result,omitempty"`
}
