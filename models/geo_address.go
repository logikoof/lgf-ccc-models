package models

//Address : ""
type Address struct {
	No           string   `json:"no" bson:"no,omitempty"`
	StateCode    string   `json:"stateCode" bson:"stateCode,omitempty"`
	DistrictCode string   `json:"districtCode" bson:"districtCode,omitempty"`
	VillageCode  string   `json:"villageCode" bson:"villageCode,omitempty"`
	CircleCode   string   `json:"circleCode" bson:"circleCode,omitempty"`
	ZoneCode     string   `json:"zoneCode" bson:"zoneCode,omitempty"`
	SectorCode   string   `json:"sectorCode" bson:"sectorCode,omitempty"`
	WardCode     string   `json:"wardCode" bson:"wardCode,omitempty"`
	AL1          string   `json:"al1" bson:"al1,omitempty"`
	Al2          string   `json:"al2" bson:"al2,omitempty"`
	PlotNo       string   `json:"plotNo" bson:"plotNo,omitempty"`
	KhataNo      string   `json:"khataNo" bson:"khataNo,omitempty"`
	PostalCode   string   `json:"postalCode" bson:"postalCode,omitempty"`
	Landmark     string   `json:"landmark" bson:"landmark,omitempty"`
	Location     Location `json:"location" bson:"location,omitempty"`
}

//Location : ""
type Location struct {
	Type        string    `json:"type" bson:"type,omitempty"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates,omitempty"`
}

//RefAddress :""
type RefAddress struct {
	State    *State    `json:"state" bson:"state,omitempty"`
	District *District `json:"district" bson:"district,omitempty"`
	Village  *Village  `json:"village" bson:"village,omitempty"`
	Zone     *Zone     `json:"zone" bson:"zone,omitempty"`
	Ward     *Ward     `json:"ward" bson:"ward,omitempty"`
}

//AddressSearch : ""
type AddressSearch struct {
	StateCode    []string `json:"stateCode" bson:"stateCode,omitempty"`
	DistrictCode []string `json:"districtCode" bson:"districtCode,omitempty"`
	VillageCode  []string `json:"villageCode" bson:"villageCode,omitempty"`
	ZoneCode     []string `json:"zoneCode" bson:"zoneCode,omitempty"`
	WardCode     []string `json:"wardCode" bson:"wardCode,omitempty"`
	Country      []string `json:"country" bson:"country,omitempty"`
	Location     Location `json:"location" bson:"location,omitempty"`
	Type         string   `json:"type" bson:"type,omitempty"`
}
