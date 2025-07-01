package pages;
// pxnLookout

import(
	TPL     "html/template"
	Gorilla "github.com/gorilla/mux"
	HTML    "github.com/PoiXson/pxnGoCommon/utils/html"
	PxnNet  "github.com/PoiXson/pxnGoCommon/net"
	WebLink "github.com/PoiXson/pxnLookout/frontend/weblink"
);



type Pages struct {
	weblink  *WebLink.WebLink
	tpl_home *TPL.Template
}



func New(weblink *WebLink.WebLink, router *Gorilla.Router) *Pages {
	pages := Pages{
		weblink: weblink,
	};
	PxnNet.AddStaticRoute(router);
	router.HandleFunc("/", pages.PageWeb_Home);
	router.HandleFunc("/favicon.ico",
		PxnNet.NewRedirect("/static/hub.ico"));
	pages.PageInit_Home();
	return &pages;
}



func (pages *Pages) GetBuilder() *HTML.Builder {
	return HTML.NewBuilder().
		WithBootstrap().
		WithBootstrapIcons().
		WithBootstrapTooltips().
		SetFavIcon("/static/hub.ico").
		AddCSS("/static/hub.css").
		SetTitle("Hub");
}
