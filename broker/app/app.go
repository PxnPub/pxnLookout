package app;
// pxnLookout Broker

import(
	Flag    "flag"
	Service "github.com/PxnPub/pxnGoCommon/service"
);



type AppBroker struct {
	service *Service.Service
//	config  *Configs.CfgBroker
}



func New() Service.AppFace {
	return &AppBroker{};
}

func (app *AppBroker) Main() {
	app.service = Service.New();
	app.service.Start();
	app.flags_and_configs(DefaultConfigFile);
	// databases
	// start things
//TODO: db
//	if err := app.heart.Start(); err != nil {
//		Log.Panicf("%s, when starting heartbeat", err); }
//	if err := app.link.Start(); err != nil {
//		Log.Panicf("%s, when starting uplink",    err); }
	app.service.WaitUntilEnd();
}



func (app *AppBroker) flags_and_configs(file string) {
	Flag.Parse();
}
