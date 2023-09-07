package models

//Circle : "Holds single state data"
type Circle struct {
	Name      string    `json:"name" bson:"name,omitempty"`
	Code      string    `json:"code"  bson:"code,omitempty"`
	ZoneCode  string    `json:"zoneCode,omitempty"  bson:"zoneCode,omitempty"`
	Status    string    `json:"status"  bson:"status,omitempty"`
	Created   Created   `json:"created"  bson:"created,omitempty"`
	Languages []string  `json:"languages"  bson:"languages,omitempty"`
	Updated   []Updated `json:"updated"  bson:"updated,omitempty"`
}

//RefCircle : "Circle with refrence data such as language..."
type RefCircle struct {
	Circle `bson:",inline"`
	Ref    struct {
		State    *State    `json:"state,omitempty" bson:"state,omitempty"`
		District *District `json:"district,omitempty" bson:"district,omitempty"`
		Village  *District `json:"village,omitempty" bson:"village,omitempty"`
		Zone     *Zone     `json:"zone,omitempty" bson:"zone,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//CircleFilter : "Used for constructing filter query"
type CircleFilter struct {
	Codes     []string `json:"codes,omitempty" bson:"codes,omitempty"`
	ZoneCodes []string `json:"zoneCodes,omitempty" bson:"zoneCodes,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}
