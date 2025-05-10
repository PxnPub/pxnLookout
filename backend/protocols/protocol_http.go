package protocols;

import(
	Time    "time"
	TamDB   "github.com/PxnPub/TamDB"
	Configs "github.com/PxnPub/pxnLookout/LookoutBackend/configs"
);



type MonProtocol_HTTP struct {
	Name     string
	Interval *Time.Duration
	Tam      *TamDB.TamDB
}



func NewMonProtocol_HTTP(name string, cfg *Configs.ConfigMonProto_HTTP,
		interval *Time.Duration, tam *TamDB.TamDB) (*MonProtocol_HTTP, error) {
	return &MonProtocol_HTTP{
		Name:     name,
		Interval: interval,
		Tam:      tam,
	}, nil;
}

func (mon *MonProtocol_HTTP) Close() {
}



func (mon *MonProtocol_HTTP) Loop(timestamp int64) {
	print("LOOP http\n");
}
