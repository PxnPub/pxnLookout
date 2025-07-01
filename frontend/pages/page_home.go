package pages;

import(
	HTTP "net/http"
	TPL  "html/template"
	HTML "github.com/PoiXson/pxnGoCommon/utils/html"
);



func (pages *Pages) PageInit_Home() {
	pages.tpl_home = TPL.Must(TPL.ParseFiles(
		"html/main.tpl",
		"html/three-column.tpl",
		"html/pages/home.tpl",
	));
}



func (pages *Pages) PageWeb_Home(out HTTP.ResponseWriter, in *HTTP.Request) {
	HTML.SetContentType(out, "html");
	build := pages.GetBuilder();
	vars := struct {
		Page  string
		Title string
	}{
		Page:  "home",
		Title: "Hub",
	};
	out.Write([]byte(build.RenderTop()));
	pages.tpl_home.ExecuteTemplate(out, "main.tpl", vars);
//TODO: three-column
	pages.tpl_home.ExecuteTemplate(out, "home.tpl", vars);
	out.Write([]byte(build.RenderBottom()));








	out.Write([]byte("OK"));
}
