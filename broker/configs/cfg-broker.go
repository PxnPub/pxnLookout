package configs;
// pxnLookout Broker - config

import(
	IOUtils "io/ioutil"
	YAML    "gopkg.in/yaml.v2"
);



type CfgBroker struct {
	BindRPC        string `yaml:"Bind-RPC"`
	ChecksumBase   uint16 `yaml:"Checksum-Base"`
	ListenInterval string `yaml:"Listen-Interval"`
	SyncInterval   string `yaml:"Sync-Interval"`
	BatchInterval  string `yaml:"Batch-Interval"`
	RateLimit CfgRateLimit   `yaml:"Rate-Limit"`
	Users map[string]CfgUser `yaml:"Users"`
	NumShards      uint8
}

type CfgRateLimit struct {
	TokenInterval string `yaml:"Token-Interval"`
	TokensPerHit  uint16 `yaml:"Tokens-Per-Hit"`
	TokensThresh  uint16 `yaml:"Tokens-Thresh"`
	TokensCap     uint16 `yaml:"Tokens-Cap"`
}

type CfgUser struct {
	Desc         string   `yaml:"Desc"`
	PermitIPs    []string `yaml:"Permit-IPs"`
	PermitWeb    bool     `yaml:"Permit-Web"`
	PermitShards []uint8  `yaml:"Permit-Shards"`
}



func LoadConfig(file string) (*CfgBroker, error) {
	data, err := IOUtils.ReadFile(file);
	if err != nil { return nil, err; }
	var config CfgBroker;
	if err := YAML.Unmarshal(data, &config); err != nil {
		return nil, err; }
	return &config, nil;
}
