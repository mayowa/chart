[![Build Status](https://api.travis-ci.org/tomarus/chart.svg)](https://travis-ci.org/tomarus/chart)
[![Coverage Status](https://coveralls.io/repos/github/tomarus/chart/badge.svg?branch=master)](https://coveralls.io/github/tomarus/chart?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mayowa/chart)](https://goreportcard.com/report/github.com/mayowa/chart)
[![GoDoc](https://godoc.org/github.com/mayowa/chart?status.svg)](https://godoc.org/github.com/mayowa/chart)

# Go Chart Lib

Dead simple rrd like bandwidth charts with focus on pixel perfect rendering of source data.

Written in Go, the output can either be an interactive SVG chart or a static PNG image.

It was written to be able to show tens or hundreds of charts in seconds without interactivity in mind.

The SVG image allows basic analytics to be performed on the chart, like measurements of time or volume, showing/hiding datasets and showing a weighted moving average on demand.

Source data can be upsampled using a simple stretch method (bar charts) or downsampled using the largest triangle three buckets algorithm.

The javascript embedded in the SVG image does not have any dependencies.

## Examples

Example screenshot:

[View as interactive SVG](http://s.chiparus.org/6/6b15c5349e894fe9.svg)

![Example Screenshot](http://s.chiparus.org/5/5caa4e08e4b2edb3.png)

Screenshots from the example website:

![Basics from example website](http://s.chiparus.org/3/3a23cff5f7c2c704.png)

## Example Usage

Example demo app:

```
go get github.com/mayowa/chart 
go run $GOPATH/src/github.com/mayowa/chart/examples/main.go
open http://localhost:3000
```

Example /proc monitoring app (linux only):

```
go get github.com/mayowa/chart
go run $GOPATH/src/github.com/mayowa/chart/examples/sysmon/main.go
open http://localhost:3001
```

Code example:

```go
import (
    "github.com/mayowa/chart"
    "github.com/mayowa/chart/svg"
)
opts := &chart.Options{
    Title:  "Traffic",
    Image:  svg.New(), // or png.New()
    Size:   "auto",    // big is 1440px, small is 720px, auto is size of dataset
    Height: 300,       // Defaults to -1, when size=auto height=width/4, otherwise set fixed height
    Scheme: "white",   // or black/random/pink/solarized or hsl:180,0.5,0.25
    Start:  start_epoch,
    End:    end_epoch,
    W:      w,
    SIBase: 1000, // or use 1024 to scale, only used when axes are not specified.
    // If you don't specify axes, they will be automatically calculated using some defaults.
    Axes: []*axis.Axis{
        axis.NewTime(axis.Bottom, "Mon 15:04").Duration(8 * time.Hour).Grid(4),
        axis.NewSI(axis.Left, 1000).Ticks(4).Grid(2),
    },
}
c, err := chart.NewChart(opts)
if err != nil {
    panic(err)
}
err = c.AddData(&data.Options{Title: "My Data Description"}, []yourData)
if err != nil {
    panic(err)
}
w.Header().Set("Content-Type", "image/svg+xml")
err = c.Render()
if err != nil {
    panic(err)
}
```

## Notes

This is an experimental work in progress for my own personal educational and research purposes.

This project has just started and a lot of stuf is still missing or incomplete. The API will not be stable until 1.0.0 is tagged in git.

This is a small list of ideas, todos and limitations:
* Custom lines and markers, like 95th percentile line, downtime markers, etc
* Add support negative values
* It supports only area charts atm
* Only 4 sources per chart supported currently
