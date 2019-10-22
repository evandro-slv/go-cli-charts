# go-cli-charts

Chart library for command line go apps.

## Introduction

To use it, just download the module with:

    go get -u github.com/evandro-slv/go-cli-charts

Then you can use the library like so:

    data := make(map[string]float64)
    
    data["Copper"] = 8.94
    data["Silver"] = 10.49
    data["Gold"] = 19.30
    data["Platinum"] = 21.45
    data["Other"] = 1.5
    
    graph := bar.Draw(data, bar.Options{
        Chart: bar.Chart{
            Height: 14,
        },
        Bars: bar.Bars{
            Width: 6,
            Margin: bar.Margin{
                Left:  3,
                Right: 3,
            },
        },
    })
    
    ftm.Println(graph)

The result will be displayed in the terminal as:

     21.45 |                                       ██████               
     19.92 |               ▄▄▄▄▄▄                  ██████               
     18.39 |               ██████                  ██████               
     16.85 |               ██████                  ██████               
     15.32 |               ██████                  ██████               
     13.79 |               ██████                  ██████               
     12.26 |               ██████                  ██████               
     10.72 |               ██████                  ██████      ▄▄▄▄▄▄   
      9.19 |   ▄▄▄▄▄▄      ██████                  ██████      ██████   
      7.66 |   ██████      ██████                  ██████      ██████   
      6.13 |   ██████      ██████                  ██████      ██████   
      4.60 |   ██████      ██████                  ██████      ██████   
      3.06 |   ██████      ██████                  ██████      ██████   
      1.53 |___██████______██████______▄▄▄▄▄▄______██████______██████___
               Copper      Gold        Other       Platinum    Silver   
     
You can also run the sample project in the `examples` folder, you just need to download the [github.com/buger/goterm](goterm) package for this:

    cd examples
    go get github.com/buger/goterm
    go run main.go
          
## Documentation

### Bar charts ▄██ █▄

#### Options

    Chart Chart    // chart options
    Bars Bars      // bars options
    UI UI          // UI / theming options
    Precision int  // float64 values precision

#### Chart

    Margin Margin // The margin of the chart (in tiles)
    Height int    // The height of the chart, not counting margins (in tiles)

#### Bars

    Margin  Margin  // The margin of the bars (in tiles)
    Width   int     // The width of the bars (in tiles)

#### UI

    XBar      int32  // X Axis Bar icon (default: '_')
    YBar      int32  // Y Axis Bar icon (default: '|')
    FullValue int32  // bar with full value icon (default: '█')
    HalfValue int32  // bar with half value icon (default: '▄')
    YLabel YLabel    // YAxis label options

#### YLabel

    Hide bool    // hides the Y Axis labels
    Spacing int  // sets the spacing on the labels (in tiles)

#### Margin

    Left int
    Right int
    Top int
    Bottom int


## Roadmap

- Column chart
    - X Axis label alignment options
- Bar chart
- Scatter chart
- Line chart
- Area chart
- Timeline chart
- Candlestick chart
