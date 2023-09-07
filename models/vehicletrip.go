package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type VehicleTrip struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID      string             `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	StartTime     string             `json:"vehicleId" bson:"vehicleId,omitempty"`
	EndTime       string             `json:"tripId" bson:"tripId,omitempty"`
	Date          string             `json:"date" bson:"date,omitempty"`
	StartLocation Location           `json:"startLocation" bson:"startLocation,omitempty"`
	EndLocation   Location           `json:"endLocation" bson:"endLocation,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	Created       *Created           `json:"created"  bson:"created,omitempty"`
}

type RefVehicleTrip struct {
	VehicleTrip `bson:",inline"`
	Ref         struct {
		OrganisationId Organisation `json:"organisationId" bson:"organisationId,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

type FilterVehicleTrip struct {
	Status         []string `json:"status,omitempty" bson:"status,omitempty"`
	OrganisationId []string `json:"organisationId" bson:"organisationId,omitempty"`
	BranchId       []string `json:"branchId" bson:"branchId,omitempty"`
	Regex          struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
