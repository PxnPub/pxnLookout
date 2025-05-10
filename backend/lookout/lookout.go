package lookout;

import(
	Log       "log"
	Fmt       "fmt"
	Time      "time"
	Sync      "sync"
	TamDB     "github.com/PxnPub/TamDB"
	TrapC     "github.com/PxnPub/pxnGoUtils/trapc"
	Configs   "github.com/PxnPub/pxnLookout/LookoutBackend/configs"
	LookProto "github.com/PxnPub/pxnLookout/LookoutBackend/protocols"
);



const DefaultInterval = "1s";
const MicroInterval = Time.Duration(800) * Time.Millisecond;



type Lookout struct {
	Name      string
	MonCfg    *Configs.CfgMon
	StopChan  chan bool
	WaitGrp   *Sync.WaitGroup
	Interval  *Time.Duration
	Protocols []LookProto.MonProtocol
	Tam       *TamDB.TamDB
}



func New(tam *TamDB.TamDB, name string, cfg *Configs.CfgMon,
		trapc *TrapC.TrapC, interval *Time.Duration) (*Lookout, error) {
	look := &Lookout{
		Name:      name,
		MonCfg:    cfg,
		StopChan:  trapc.NewStopChan(),
		WaitGrp:   trapc.WaitGrp,
		Interval:  interval,
		Tam:       tam,
		Protocols: make([]LookProto.MonProtocol, 0),
	};
	// load protocols
	for protokey, protoraw := range cfg.ProtosRaw {
		Fmt.Printf("   > %s\n", protokey);
		proto, err := LookProto.New(name, protokey, protoraw, tam, interval);
		if err != nil { return nil, err; }
		look.Protocols = append(look.Protocols, proto);
	}
	return look, nil;
}



func (look *Lookout) Loop() {
	look.WaitGrp.Add(1);
	defer func() {
		look.Close();
		look.WaitGrp.Done();
	}();
	// monitor loop
	for loops := 1; ; loops++ {
		now := Time.Now();
		var sleep Time.Duration = *look.Interval - now.Sub(now.Truncate(*look.Interval));
		// first loop, 2 second minimum
		if loops == 1 {
			if sleep <= Time.Duration(2) * Time.Second {
				sleep += Time.Duration(2) * Time.Second;
			}
		}
		// sleep loop
		for sleep > 0 {
			// stop channel
			select {
			case stopping := <-look.StopChan:
				if stopping {
					Log.Printf(" [ %s ] Stopping monitor..\n", look.Name);
					return;
				}
			default:
			}
			if sleep > MicroInterval {
				Time.Sleep(MicroInterval);
				sleep -= MicroInterval;
			// final sleep
			} else {
				Time.Sleep(sleep);
				sleep = 0;
			}
		}
		// run protocols
		timestamp := int64(Time.Now().Truncate(*look.Interval).Unix());
		for _, proto := range look.Protocols {
			proto.Loop(timestamp);
		}
	}
}



func (look *Lookout) GetTamDB() *TamDB.TamDB {







	return look.Tam;
}



func (look *Lookout) Close() {
	for _, protocol := range look.Protocols {
		protocol.Close();
	}
}
