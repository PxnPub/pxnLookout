package protocols;

import(
	Fmt       "fmt"
	Time      "time"
	TamDB     "github.com/PxnPub/TamDB"
	MapStruct "github.com/go-viper/mapstructure/v2"
	Configs   "github.com/PxnPub/pxnLookout/LookoutBackend/configs"
);



type MonProtocol interface {
	Loop(timestamp int64)
	Close()
}



func New(name string, protokey string, protoraw interface{},
		tam *TamDB.TamDB, interval *Time.Duration) (MonProtocol, error){
	switch protokey {
		case "Ping": {
			var cfg Configs.ConfigMonProto_Ping;
			raw := protoraw.(map[string]interface{});
			if err := MapStruct.Decode(raw, &cfg); err != nil { return nil, err; }
			return NewMonProtocol_Ping(name, &cfg, interval, tam);
		}
		case "HTTP": {
			var cfg Configs.ConfigMonProto_HTTP;
			raw := protoraw.(map[string]interface{});
			if err := MapStruct.Decode(raw, &cfg); err != nil { return nil, err; }
			return NewMonProtocol_HTTP(name, &cfg, interval, tam);
		}
		case "SNMP": {
			var cfg Configs.ConfigMonProto_SNMP;
			raw := protoraw.(map[string]interface{});
			if err := MapStruct.Decode(raw, &cfg); err != nil { return nil, err; }
			return NewMonProtocol_SNMP(name, &cfg, interval, tam);
		}
	}
	return nil, Fmt.Errorf("Invalid protocol: %s", protokey);
}
