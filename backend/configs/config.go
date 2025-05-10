package configs;



type CfgMon struct {
	Enable     *bool                  `json:"Enable,omitempty"`
	ProtosRaw  map[string]interface{} `json:"Protocols"`
}



type ConfigMonProto_Ping struct {
	Hosts  []string  `json:"Hosts,omitempty"`
}

type ConfigMonProto_HTTP struct {
	Hosts  []string  `json:"Hosts,omitempty"`
}

type ConfigMonProto_SNMP struct {
	Host       string             `json:"Host,omitempty"`
	Community  string             `json:"Community,omitempty"`
	Nodes      map[string]string  `json:"Nodes,omitempty"`
}
