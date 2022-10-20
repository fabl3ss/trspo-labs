package plot

import (
	"awesomeProject1/src/config"
	"image/color"
	"log"
	"sync"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func PrintPlotValues(wg *sync.WaitGroup, g *plot.Plot, values []int) {
	wg.Add(1)
	coords := make(plotter.XYs, len(values))

	for i := range coords {
		coords[i].X = float64(i)
		coords[i].Y = float64(values[i])
	}
	addPlotLine(g, coords)

	wg.Done()
}

func CreatePlot() *plot.Plot {
	chart := plot.New()
	initDefaultPlot(chart)
	return chart
}

func SavePlot(p *plot.Plot) {
	err := p.Save(
		vg.Length(config.PlotWidth),
		vg.Length(config.PlotHeight),
		config.PlotFilename+".png",
	)
	if err != nil {
		log.Fatalln(err)
	}
}

func initDefaultPlot(plot *plot.Plot) {
	plot.Y.Color = color.White
	plot.X.Color = color.White
	plot.Title.Text = config.PlotTitle
	plot.BackgroundColor = config.PlotBackground
}

func addPlotLine(p *plot.Plot, coords plotter.XYs) {
	err := plotutil.AddLinePoints(p, "", coords)
	if err != nil {
		log.Fatalln(err)
	}
}
