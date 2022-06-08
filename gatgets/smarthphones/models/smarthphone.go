package smartphones

// Smartphone structure for smarthphones
type Smarthphone struct {
	Id            int64
	Name          string
	Price         int
	CountryOrigin string
	Os            string
}

//CreateSmartphoneCMD
type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	Os            string `json:"os"`
}
