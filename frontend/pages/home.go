package pages;

import(
	HTTP     "net/http"
	Template "html/template"
	JSON     "encoding/json"
	TamAPI   "github.com/PxnPub/TamDB/client"
);



type PageVars_Home struct {
	Title string
}



func Page_Home(w HTTP.ResponseWriter, r *HTTP.Request) {
	tpl, err := Template.ParseFiles(
		"html/main.tpl",
		"html/home.tpl",
	);
	if err != nil {
		HTTP.Error(w, err.Error(), HTTP.StatusInternalServerError);
		return;
	}
	vars := PageVars_Home{
		Title: "pxnLookout",
	};
	err = tpl.Execute(w, vars);
	if err != nil {
		HTTP.Error(w, err.Error(), HTTP.StatusInternalServerError);
		return;
	}
}



func API_Home(w HTTP.ResponseWriter, r *HTTP.Request) {
//TODO
name := "router";
	api, err := TamAPI.New("tcp", "127.0.0.1:9999");
	if err != nil { panic(err); }
	data_in  := api.Query(name, name, "WAN-in");
	data_out := api.Query(name, name, "WAN-out");
	fig := map[string]interface{}{
		"max": 50000000,
		"data": []interface{}{
			map[string]interface{}{
				"type": "scatter",
				"x": data_in.DataX,
				"y": data_in.DataY,
			},
			map[string]interface{}{
				"type": "scatter",
				"x": data_out.DataX,
				"y": data_out.DataY,
			},
		},
	};
	figjson, err := JSON.MarshalIndent(fig, "", "\t");
	if err != nil {
		HTTP.Error(w, err.Error(), HTTP.StatusInternalServerError);
		return;
	}
	w.Header().Set("Content-Type", "application/json");
	w.Write([]byte(figjson));
}
/*
//	api := ApiClient.New("tcp", "127.0.0.1:9999");
//	if err != nil { panic(err); }
//	defer api.Close();
//TODO: name parameter
//	name := "router";
//	traces := make([]Typely.Trace, 0);
//	api.Query(traces, name, name, "WAN-in");
//	api.Query(traces, name, name, "WAN-out");

//dataX := []int64{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20 };
//dataY := []int64{ 5, 2, 3, 5, 6, 7, 9, 8, 4,  5,  3,  5,  6,  5,  8,  6,  5,  3,  1,  4 };
//	dataX = make([]int64, 0);
//	dataY = make([]int64, 0);
//	request_in  := ApiQuery{ Name: "router", Key: "WAN-in"  }
//	request_out := ApiQuery{ Name: "router", Key: "WAN-out" }
//	var reply_in  ApiData;
//	var reply_out ApiData;
//TODO
//	if err := client.Call("TamAPI.Query", request_in,  &reply_in ); err != nil { panic(err); }
//	if err := client.Call("TamAPI.Query", request_out, &reply_out); err != nil { panic(err); }
//	w.Write([]byte(reply));
	// plotly

//	traces = append(traces,
//		&Plotly.Scatter{
//			X: Typely.DataArray(reply_in.DataX),
//			Y: Typely.DataArray(reply_in.DataY),
//		},
//		&Plotly.Scatter{
//			X: Typely.DataArray(reply_out.DataX),
//			Y: Typely.DataArray(reply_out.DataY),
//		},
//	);
//	var maxval int64 = 0;
//	if v := max(reply_in.DataY ); maxval < v { maxval = v; }
//	if v := max(reply_out.DataY); maxval < v { maxval = v; }
/ *
	fig := &Plotly.Fig{
		Data: traces,
		Layout: &Plotly.Layout{
			Title: &Plotly.LayoutTitle{
				Text: "WAN",
			},
			Xaxis: &Plotly.LayoutXaxis{
				Uirevision: "constant",
				Rangeslider: &Plotly.LayoutXaxisRangeslider{
					Visible: Typely.True,
					Autorange: Typely.True,
				},
			},
			Yaxis: &Plotly.LayoutYaxis{
				Range: Typely.DataArray([]int64{0, maxval+1}),
			},
			Showlegend: Typely.False,
		},
	};
	figjson, err := json.MarshalIndent(fig, "", "\t");
*/



func max(data []int64) int64 {
	highest := data[0];
	for _, val := range data {
		if highest < val {
			highest = val;
		}
	}
	return highest;
}

var count = 0;
func rotate(array []float64, k int) {
	k = k % len(array);
	temp := append(array[k:], array[0:k]...);
	copy(array, temp);
}
