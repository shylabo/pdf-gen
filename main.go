package main

import (
	"fmt"
	"golang-pdf-maroto/data"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildFruitList(m)

	err := m.OutputFileAndClose(("pdfs/div_fruits.pdf"))
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully.")

}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/logo_div_shuya.jpg", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					fmt.Println("Image file was not loaded:", err)
				}

			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Prepared for you by the Div Shuya labo", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})

	})
}

func buildFruitList(m pdf.Maroto) {
	tableHeadings := []string{"Fruit", "Description", "Price"}
	contents := data.FruitList(20)
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealcolor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightPurpleColor,
	})
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getTealcolor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
