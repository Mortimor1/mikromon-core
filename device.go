package core

type Device struct {
	Id        string `json:"id"`
	IpAddress string `json:"ipaddress"`
	Name      string `json:"name"`
	State     bool   `json:"state"`
	Status    string `json:"status"`
	Group     string `json:"group"` //id group
}
