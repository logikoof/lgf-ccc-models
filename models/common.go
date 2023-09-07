package models

import "time"

//Created : "Used To store created On and created by details"
type Created struct {
	On       *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	By       string     `json:"by,omitempty" form:"by" bson:"by,omitempty"`
	Scenario string     `json:"scenario" bson:"scenario,omitempty"`
}

//Updated : ""
type Updated struct {
	On       *time.Time `json:"on" bson:"updatedOn,omitempty"`
	By       string     `json:"by" bson:"updatedBy,omitempty"`
	Scenario string     `json:"scenario" bson:"scenario,omitempty"`
	ByType   string     `json:"bytype,omitempty" form:"bytype" bson:"bytype,omitempty"`
}
type CreatedV2 struct {
	On      *time.Time `json:"on" form:"on" bson:"on,omitempty"`
	By      string     `json:"by,omitempty" form:"by" bson:"by,omitempty"`
	ByType  string     `json:"bytype,omitempty" form:"bytype" bson:"bytype,omitempty"`
	Remarks string     `json:"remarks" bson:"remarks,omitempty"`
}
type Countstruct struct {
	Count int64 `json:"count" form:"count" bson:"count,omitempty"`
}
