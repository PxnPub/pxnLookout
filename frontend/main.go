package main;

import(
	OS    "os"
	Log   "log"
	HTTP  "net/http"
	Pages "github.com/PxnPub/pxnLookout/LookoutFrontend/pages"
);



func main() {
	print("\n");
	// routes
	HTTP.HandleFunc("/",    Pages.Page_Home);
	HTTP.HandleFunc("/api", Pages.API_Home );
	// start listening
	if err := HTTP.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		Log.Panic(err);
	}
	print(" ~end~ \n");
	print("\n"); OS.Exit(0);
}
