package hosts

type Host struct {
	Hostname      string   `json:"hostname"`
	HWAddress     string   `json:"hw-address"`
	ClientClasses []string `json:"client-classes,omitempty"`
}
