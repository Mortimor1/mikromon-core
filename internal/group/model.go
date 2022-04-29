package group

type Group struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name,omitempty"`
	State bool   `json:"state" bson:"state,omitempty"`
	Group string `json:"group" bson:"group,omitempty"` //id group
}
