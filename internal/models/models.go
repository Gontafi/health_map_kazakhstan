package models

type Region struct {
	ID   int16
	Name string
}

type Statistics struct {
	RegionID int16  `json:"region_id,omitempty"`
	Sick     int    `json:"sick"`
	Dead     int    `json:"dead"`
	Cured    int    `json:"cured"`
	SickName string `json:"sick_name"`
}

type Sick struct {
	RegionID int16
	SickName string
	Count    int
	TypeID   int
}
