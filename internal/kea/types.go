package kea

type Reservation struct {
	HWAddress string `json:"hw-address"`
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ip-address,omitempty"`
}

type ReservationAddRequest struct {
	Command   string      `json:"command"`
	Service   []string    `json:"service"`
	Arguments Reservation `json:"arguments"`
}
