package device

type Device struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	IpAddress string `json:"ipaddress" bson:"ipaddress,omitempty"`
	Name      string `json:"name" bson:"name,omitempty"`
	State     bool   `json:"state" bson:"state,omitempty"`
	Status    string `json:"status" bson:"status,omitempty"`
	Group     string `json:"group" bson:"group,omitempty"` //id group
}
