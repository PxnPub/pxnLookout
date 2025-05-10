
{{define "page-content"}}
<div class="chartbox">

	<div id="plot"></div>
	<script src="https://cdn.plot.ly/plotly-3.0.1.min.js"></script>
	<script>
	var layout = {
//		title: {
//			text: "WAN",
//		},
		xaxis: {
			uirevision: "constant",
			rangeslider: {
				visible: true,
				autorange: true,
			},
		},
		yaxis: {
//			range: [0, 12000000],
		},
		showlegend: false,
	};
	async function FetchAndUpdateChart() {
		const resp = await fetch('/api');
		const json = await resp.json();
		layout.yaxis.range = [ 0, json.max ];
		Plotly.react('plot', json.data, layout);
	}
	FetchAndUpdateChart();
	setInterval(FetchAndUpdateChart, 5000);
	</script>

</div>
{{end}}
