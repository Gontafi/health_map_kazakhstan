package models

type Region struct {
	ID   int16
	Name string
}

type Statistics struct {
	RegionID int16
	Sick     int
	Dead     int
	Cured    int
	SickName string
}

type Sick struct {
	RegionID int16
	SickName string
	Count    int
	TypeID   int
}
