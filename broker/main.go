package main;

import App "github.com/PoiXson/pxnLookout/broker/app";

const Version = "{{{VERSION}}}";

func main() {
	app := App.New();
	app.Main();
}
