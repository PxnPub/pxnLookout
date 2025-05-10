package protocols;

import(
	Time    "time"
	TamDB   "github.com/PxnPub/TamDB"
	Configs "github.com/PxnPub/pxnLookout/LookoutBackend/configs"
);



type MonProtocol_Ping struct {
	Name     string
	Interval *Time.Duration
	Tam      *TamDB.TamDB
}



func NewMonProtocol_Ping(name string, cfg *Configs.ConfigMonProto_Ping,
		interval *Time.Duration, tam *TamDB.TamDB) (*MonProtocol_Ping, error) {
	return &MonProtocol_Ping{
		Name:     name,
		Interval: interval,
		Tam:      tam,
	}, nil;
}

func (mon *MonProtocol_Ping) Close() {
}



func (mon *MonProtocol_Ping) Loop(timestamp int64) {
	print("LOOP Ping\n");
}
