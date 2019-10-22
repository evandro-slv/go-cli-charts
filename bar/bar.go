package bar

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Options struct {
	Chart     Chart // chart options
	Bars      Bars  // bars options
	UI        UI    // UI / theming options
	Precision int   // float64 values precision
}

type Chart struct {
	Margin Margin // The margin of the chart (in tiles)
	Height int    // The height of the chart, not counting margins (in tiles)
}

type Bars struct {
	Margin Margin // The margin of the bars (in tiles)
	Width  int    // The width of the bars (in tiles)
}

type Margin struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

type UI struct {
	XBar      int32  // X Axis Bar icon (default: '_')
	YBar      int32  // Y Axis Bar icon (default: '|')
	FullValue int32  // bar with full value icon (default: '█')
	HalfValue int32  // bar with half value icon (default: '▄')
	YLabel    YLabel // YAxis label options
}

type YLabel struct {
	Hide    bool // hides the Y Axis labels
	Spacing int  // sets the spacing on the labels (in tiles)
}

func Draw(data map[string]float64, opt Options) string {
	chart := ""
	high := -math.MaxFloat64
	low := math.MaxFloat64

	// define UI
	if opt.UI.XBar == 0 {
		opt.UI.XBar = '_'
	}

	if opt.UI.YBar == 0 {
		opt.UI.YBar = '|'
	}

	if opt.UI.FullValue == 0 {
		opt.UI.FullValue = '█'
	}

	if opt.UI.HalfValue == 0 {
		opt.UI.HalfValue = '▄'
	}

	if opt.UI.YLabel.Spacing == 0 {
		opt.UI.YLabel.Spacing = 1
	}

	if opt.Precision == 0 {
		opt.Precision = 2
	}

	// get data map keys
	keys := make([]string, 0)

	for k, _ := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// get highest and lowest chart number
	for _, k := range keys {
		v := data[k]

		if v > high {
			high = v
		}

		if v < low {
			low = v
		}
	}

	steps := high / float64(opt.Chart.Height)
	YLabelSize := len(strconv.FormatFloat(high, 'f', opt.Precision, 64))

	// draw chart margin
	for i := 0; i < opt.Chart.Margin.Top; i++ {
		chart += "\n"
	}

	// draw Y data
	for i := opt.Chart.Height; i > 0; i-- {
		r := float64(i) * steps
		chart += strings.Repeat(" ", opt.Chart.Margin.Left)

		if !opt.UI.YLabel.Hide && i%opt.UI.YLabel.Spacing == 0 {
			chart += drawYLabel(r, YLabelSize, opt.Precision, string(opt.UI.YBar))
		} else if !opt.UI.YLabel.Hide {
			chart += strings.Repeat(" ", YLabelSize) + " " + string(opt.UI.YBar)
		}

		for _, k := range keys {
			v := data[k]

			marginSep := " "

			if i == 1 {
				marginSep = string(opt.UI.XBar)
			}

			chart += strings.Repeat(marginSep, opt.Bars.Margin.Left)

			if v >= r {
				chart += strings.Repeat(string(opt.UI.FullValue), opt.Bars.Width)
			} else if v >= r-steps/2 {
				chart += strings.Repeat(string(opt.UI.HalfValue), opt.Bars.Width)
			} else {
				chart += strings.Repeat(marginSep, opt.Bars.Width)
			}

			chart += strings.Repeat(marginSep, opt.Bars.Margin.Right)
		}

		chart += strings.Repeat(" ", opt.Chart.Margin.Right)
		chart += "\n"
	}

	// draw X data
	chart += drawXRows(YLabelSize, &keys, opt)

	// draw chart margin
	for i := 0; i < opt.Chart.Margin.Bottom; i++ {
		chart += "\n"
	}

	return chart
}

func drawYLabel(label float64, maxSize int, precis int, sep string) string {
	return fmt.Sprintf("%"+strconv.Itoa(maxSize)+"v "+sep, strconv.FormatFloat(label, 'f', precis, 64))
}

func drawXRows(YLabelSize int, keys *[]string, opt Options) string {
	chart := ""
	barSize := opt.Bars.Width + opt.Bars.Margin.Left + opt.Bars.Margin.Right

	chart += strings.Repeat(" ", opt.Chart.Margin.Left)

	if !opt.UI.YLabel.Hide {
		chart += strings.Repeat(" ", YLabelSize)
		chart += strings.Repeat(" ", 2) // bar size => ( |)
	}

	chart += strings.Repeat(" ", opt.Bars.Margin.Left)

	for _, k := range *keys {
		subs := ""

		if len(k) > barSize-1 {
			r := []rune(k)
			subs += string(r[0:barSize-2]) + "… "
		} else {
			subs += k + strings.Repeat(" ", opt.Bars.Margin.Left-len(k)+opt.Bars.Margin.Right+opt.Bars.Width)
		}

		chart += subs
	}

	chart += strings.Repeat(" ", opt.Chart.Margin.Right)

	return chart
}
