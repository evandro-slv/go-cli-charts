package main

import (
	tm "github.com/buger/goterm"
	"github.com/evandro-slv/go-cli-charts/bar"
)

func main() {
	data := make(map[string]float64)

	data["Copper"] = 8.94
	data["Silver"] = 10.49
	data["Gold"] = 19.30
	data["Platinum"] = 21.45
	data["Other"] = 1.5

	graph := bar.Draw(data, bar.Options{
		Chart: bar.Chart{
			Height: 14,
			Margin: bar.Margin{
				Top:    1,
				Bottom: 1,
				Left:   1,
			},
		},
		Bars: bar.Bars{
			Width: 6,
			Margin: bar.Margin{
				Left:  3,
				Right: 3,
			},
		},
		Precision: 2,
	})

	graph2 := bar.Draw(data, bar.Options{
		Chart: bar.Chart{
			Height: 14,
			Margin: bar.Margin{
				Top:    1,
				Bottom: 1,
				Left:   1,
			},
		},
		Bars: bar.Bars{
			Width: 1,
			Margin: bar.Margin{
				Left:  4,
				Right: 4,
			},
		},
		Precision: 2,
		UI: bar.UI{
			FullValue: '-',
			HalfValue: '_',
			YLabel: bar.YLabel{
				Hide: true,
			},
		},
	})

	graph3 := bar.Draw(data, bar.Options{
		Chart: bar.Chart{
			Height: 14,
			Margin: bar.Margin{
				Top:    1,
				Bottom: 1,
				Left:   1,
			},
		},
		Bars: bar.Bars{
			Width: 1,
			Margin: bar.Margin{
				Left:  4,
				Right: 4,
			},
		},
		Precision: 2,
		UI: bar.UI{
			XBar: ' ',
			YBar: ' ',
			YLabel: bar.YLabel{
				Spacing: 2,
			},
		},
	})

	tm.Println("Simple chart")
	tm.Println(graph)

	tm.Println("Custom UI")
	tm.Println(graph2)

	tm.Println("Y Label Spacing")
	tm.Println(graph3)
	tm.Flush()
}
