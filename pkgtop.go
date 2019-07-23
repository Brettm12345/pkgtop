package main

import (
	"log"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

var i int
var termGrid, dfGrid, pkgGrid *ui.Grid
var pkgText *widgets.Paragraph
var gau0, gau1, gau2, gau3 *widgets.Gauge
var dfgau []*widgets.Gauge
var pkgl0, pkgl1, pkgl2, pkgl3 *widgets.List

func initWidgets() {
	termGrid, dfGrid, pkgGrid = 
		ui.NewGrid(),
		ui.NewGrid(),
		ui.NewGrid()
	gau0, gau1, gau2, gau3 = 
		widgets.NewGauge(), 
		widgets.NewGauge(), 
		widgets.NewGauge(), 
		widgets.NewGauge()
	pkgl0, pkgl1, pkgl2, pkgl3 = 
		widgets.NewList(),
		widgets.NewList(),
		widgets.NewList(),
		widgets.NewList()
	pkgText = widgets.NewParagraph()
	dfgau = []*widgets.Gauge{
		gau0, gau1, gau2, gau3,
	}
}

func setDiskUsage(diskUsage map[string]int) bool {
	i := 0
	for name, perc := range diskUsage {
		dfgau[i].Title = name
		dfgau[i].Percent = perc
		i++
	}
	dfGrid.Set(
		ui.NewRow(1.0/4,
			ui.NewCol(1.0, gau0),
		),
		ui.NewRow(1.0/4,
			ui.NewCol(1.0, gau1),
		),
		ui.NewRow(1.0/4,
			ui.NewCol(1.0, gau2),
		),
		ui.NewRow(1.0/4,
			ui.NewCol(1.0, gau3),
		),
	)
	return true
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	initWidgets()
	defer ui.Close()

	diskUsage := map[string]int {
		"dev": 0,
		"run": 1,
		"/dev/sda1": 75,
		"tmpfs": 4,
	}

	setDiskUsage(diskUsage)

	pkgText.Text = "~"
	//pkgText.Border = false

	pkgs := []string{
		"apache~2.4.39-1~6.25MiB~'Fri 11 Jan 2019 03:34:39'",
		"autoconf~2.69-5~2.06MiB~'Fri 11 Jan 2019 03:34:39'",
		"automake~1.16.1-1~1598.00KiB~'Fri 11 Jan 2019 03:34:39'",
		"bind-tools~9.14.2-1~5.85MiB~'Fri 11 Jan 2019 03:34:39'",
		"bison~3.3.2-1~2013.00KiB~'Fri 11 Jan 2019 03:34:39'",
		"brook~20190401-1~13.98MiB~'Fri 11 Jan 2019 03:34:39'",
		"chafa~1.0.1-1~327.00KiB~'Fri 11 Jan 2019 03:34:39'",
		"cmatrix~2.0-1~95.00KiB~'Fri 11 Jan 2019 03:34:39'",
		"compton~6.2-2~306.00KiB~'Fri 11 Jan 2019 03:34:39'",
		"docker~1:18.09.6-1~170.98MiB~'Fri 11 Jan 2019 03:34:39'",
	}

	pkgl0.Title = "List"
	pkgl0.Rows = pkgs
	pkgl0.WrapText = false
	pkgl0.Border = false

	pkgGrid.Set(
		ui.NewRow(1.0,
			ui.NewCol(1.0/4, pkgl0),
			ui.NewCol(1.0/4, pkgl0),
			ui.NewCol(1.0/4, pkgl0),
			ui.NewCol(1.0/4, pkgl0),
		),
	)
	
	termWidth, termHeight := ui.TerminalDimensions()
	termGrid.SetRect(0, 0, termWidth, termHeight)
	termGrid.Set(
		ui.NewRow(1.0/4,
			ui.NewCol(1.0/2, dfGrid),
			ui.NewCol(1.0/2, pkgText),
		),
		ui.NewRow(1.0/1.6,
			ui.NewCol(1.0/1, pkgGrid),
		),
		ui.NewRow(1.0/8,
			ui.NewCol(1.0/1, pkgText),
		),
	)
	ui.Render(termGrid)
	uiEvents := ui.PollEvents()
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>", "<C-d>":
				return
			case "<Resize>":
				payload := e.Payload.(ui.Resize)
				termGrid.SetRect(0, 0, payload.Width, payload.Height)		
				ui.Clear()
				ui.Render(termGrid)
			}
		}
	}

}