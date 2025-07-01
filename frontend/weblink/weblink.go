package weblink;
// pxnLookout Frontend -> Broker

import(
	UtilsRPC  "github.com/PxnPub/pxnGoCommon/rpc"
	Service   "github.com/PxnPub/pxnGoCommon/service"
	API_Front "github.com/PxnPub/pxnLookout/api/front"
);



type WebLink struct {
	service *Service.Service
	rpc     *UtilsRPC.ClientRPC
	API     API_Front.ServiceFrontendAPIClient
}



func New(service *Service.Service, addr string) *WebLink {
	rpc := UtilsRPC.NewClientRPC(service, addr);
//TODO
//	rpc.UseTLS = true;
	return &WebLink{
		service: service,
		rpc:     rpc,
	};
}



func (link *WebLink) Start() error {
	if err := link.rpc.Start(); err != nil { return err; }
	link.API = API_Front.NewServiceFrontendAPIClient(link.rpc.GetClientGRPC());
	return nil;
}

func (link *WebLink) Close() {
	link.rpc.Close();
}
