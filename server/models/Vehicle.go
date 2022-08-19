package models

type Vehicle struct {
	ID     string `bson:"-"`
	Plate  string `bson:"plate"`
	Make   string `bson:"make"`
	Model  int64  `bson:model`
	Series string `bson:series`
	Color  string `bson:color`
}
