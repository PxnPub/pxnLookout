package app;
// pxnLookout Frontend

import(
	Log     "log"
	Flag    "flag"
	Flagz   "github.com/PxnPub/pxnGoCommon/utils/flagz"
	Utils   "github.com/PxnPub/pxnGoCommon/utils"
	UtilsFS "github.com/PxnPub/pxnGoCommon/utils/fs"
	PxnNet  "github.com/PxnPub/pxnGoCommon/utils/net"
	Service "github.com/PxnPub/pxnGoCommon/service"
	WebLink "github.com/PxnPub/pxnLookout/frontend/weblink"
	Configs "github.com/PxnPub/pxnLookout/frontend/configs"
	Pages   "github.com/PxnPub/pxnLookout/frontend/pages"
);



type AppFront struct {
	service *Service.Service
	websvr  *PxnNet.WebServer
	pages   *Pages.Pages
	link    *WebLink.WebLink
	config  *Configs.CfgFront
}



func New() Service.AppFace {
	return &AppFront{};
}

func (app *AppFront) Main() {
	app.service = Service.New();
	app.service.Start();
	app.flags_and_configs(DefaultConfigFile);
	// web server
	app.websvr = PxnNet.NewWebServer(
		app.service,
		app.config.BindWeb,
		app.config.Proxied,
	);
	router := app.websvr.WithGorilla();
	app.pages = Pages.New(app.link, router);
	// start things
	if err := app.websvr.Start(); err != nil { Log.Panic(err); }
	app.service.WaitUntilEnd();
}



func (app *AppFront) flags_and_configs(file string) {
	var flag_bindweb string;
	var flag_proxied bool;
	Flagz.String(&flag_bindweb, "bind", "");
	Flagz.Bool  (&flag_proxied, "proxied" );
	Flag.Parse();
	// load config
	cfg, err := UtilsFS.LoadConfig[Configs.CfgFront](file);
	if err != nil { Log.Panicf("%s, when loading config %s", err, file); }
	// bind web
	if flag_bindweb != "" { cfg.BindWeb = flag_bindweb;          }
	if cfg.BindWeb  == "" { cfg.BindWeb = PxnNet.DefaultBindWeb; }
	if flag_proxied       { app.config.Proxied = true;           }
	app.config = cfg;
}
