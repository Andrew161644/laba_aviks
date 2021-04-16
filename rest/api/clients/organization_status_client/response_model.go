package organization_status_client

type OrgStatusResponseModel struct {
	Kabsl  float64 `json:"kabsl"`
	Kcurrl float64 `json:"kcurrl"`
	Kfastl float64 `json:"kfastl"`
	Kfin   float64 `json:"kfin"`
	Kfu    float64 `json:"kfu"`
	Kk     float64 `json:"kk"`
	Kn     float64 `json:"kn"`
	Report string  `json:"report"`
}
