package app;

import(
	OS      "os"
	Log     "log"
	Fmt     "fmt"
	Time    "time"
	IOUtil  "io/ioutil"
	JSON    "encoding/json"
	TrapC   "github.com/PxnPub/pxnGoUtils/trapc"
	Utils   "github.com/PxnPub/pxnGoUtils"
	TamDB   "github.com/PxnPub/TamDB"
	TamAPI  "github.com/PxnPub/TamDB/server"
	Lookout "github.com/PxnPub/pxnLookout/LookoutBackend/lookout"
	Configs "github.com/PxnPub/pxnLookout/LookoutBackend/configs"
);



const BindSocket = "tcp://127.0.0.1:9999";
//const BindSocket = "unix://lookout.sock";



const ConfigFileMonitor = "monitors.json";



func Main(trapc *TrapC.TrapC) {
//TODO
interval, err := Time.ParseDuration(Lookout.DefaultInterval);
if err != nil { panic(err); }
//TODO
driver := TamDB.Driver_LibSQL;
//driver := TamDB.Driver_DuckDB;
	// find monitors.json
	file := Utils.FindFile(ConfigFileMonitor, Utils.DefaultConfigSearchPaths...);
	if file == "" { Log.Panic("Config file not found: ", ConfigFileMonitor); }
	// load monitors.json
	cfgmons := load(file);
	if len(cfgmons) == 0 {
		Log.Panic("No monitor targets configured");
	}
	// start monitoring
	Log.Print("Starting monitors..\n");
	Fmt.Printf("Interval: %s\n", interval);
	sleep, err := Time.ParseDuration("100ms");
	if err != nil { panic(err); }
	api := TamAPI.New(trapc, BindSocket);
	lookouts := make([]Lookout.Lookout, 0);
	for name, cfg := range cfgmons {
		if cfg.Enable == nil || *cfg.Enable {
			name_db, name_table := TamDB.SplitNameKey(name);
			Fmt.Printf(" [ %s ] starting..\n", name_db);
			// open database
			tam_rw, err := TamDB.New(driver, name_db, name_table, "db/", true);
			if err != nil { panic(err); }
			tam_ro, err := TamDB.New(driver, name_db, name_table, "db/", false);
			if err != nil { panic(err); }
			// monitor task
			look, err := Lookout.New(tam_rw, name_db, &cfg, trapc, &interval);
			if err != nil { panic(err); }
			lookouts = append(lookouts, *look);
			go look.Loop();
			Time.Sleep(sleep);
			// api
			api.AddTamDB(tam_ro, name_db);
		} else {
			Fmt.Printf(" [ %s ] disabled\n", name);
		}
	}
	// start api
	api.StartListening();
}



// load monitors.json
func load(file string) map[string]Configs.CfgMon {
	handle, err := OS.Open(file);
	if err != nil { panic(err); }
	defer handle.Close();
	data, err := IOUtil.ReadAll(handle);
	if err != nil { panic(err); }
	var cfgmons map[string]Configs.CfgMon;
	err = JSON.Unmarshal(data, &cfgmons);
	if err != nil { panic(err); }
	return cfgmons;
}
