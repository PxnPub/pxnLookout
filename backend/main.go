package main;

import(
	OS    "os"
	Time  "time"
	TrapC "github.com/PxnPub/pxnGoUtils/trapc"
	App   "github.com/PxnPub/pxnLookout/LookoutBackend/app"
);



func main() {
	print("\n");
	trapc := TrapC.New();
	App.Main(trapc);
	Time.Sleep(Time.Duration(250) * Time.Millisecond);
	print("\n"); trapc.Wait();
	print(" ~end~ \n");
	print("\n"); OS.Exit(0);
}
