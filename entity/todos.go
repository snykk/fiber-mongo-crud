package entity

type Todo struct {
	Id       string `bson:"_id" json:"id"`
	Activity string `bson:"activity" json:"activity"`
	Priority string `bson:"priority" json:"priority"`
	IsDone   bool   `bson:"is_done" json:"is_done"`
}
