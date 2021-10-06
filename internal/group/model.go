package group

type Group struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	State bool   `json:"state"`
	Group string `json:"group"` //id group
}
