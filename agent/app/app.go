package app;
// pxnLookout Agent

import(
	Flag    "flag"
	Service "github.com/PxnPub/pxnGoCommon/service"
);



type AppLookAgent struct {
	service *Service.Service
//	config  *Configs.CfgLookAgent
}



func New() Service.AppFace {
	return &AppLookAgent{};
}

func (app *AppLookAgent) Main() {
	app.service = Service.New();
	app.service.Start();
	app.flags_and_configs(DefaultConfigFile);
	// start things
	app.service.WaitUntilEnd();
}



func (app *AppLookAgent) flags_and_configs(file string) {
	Flag.Parse();
}
