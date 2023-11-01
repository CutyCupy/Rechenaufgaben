package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pt := 25.0

	pdf.AddUTF8Font("MyFont", "B", "./font.ttf")
	pdf.SetFont("MyFont", "B", pt)
	tpl1 := gofpdi.ImportPage(pdf, "template.pdf", 1, "/MediaBox")

	// Draw imported template onto page
	gofpdi.UseImportedTemplate(pdf, tpl1, 0, 0, 210, 0)
	const (
		marginX = 15.0
		marginY = 30.0
	)

	pdf.Ln(marginY)
	pdf.SetLeftMargin(marginX)
	pdf.SetRightMargin(marginX)

	offset := (210.0 - 2.0*marginX) / 3.0

	lineHeight := pdf.PointConvert(pt)

	signs := []string{"+", "-", "x"}

	writer := func(a string) {
		w := pdf.GetStringWidth(a)
		pdf.SetX(pdf.GetX() + (lineHeight-w)/2)
		pdf.Write(lineHeight, a)
		pdf.SetX(pdf.GetX() + (lineHeight-w)/2)
	}

	for pdf.GetY() < 297-(marginX*2) {
		for x := 0; x < 3; x++ {
			pdf.SetX(offset*float64(x) + marginX)

			a := rand.Intn(10)
			b := rand.Intn(10)

			sign := signs[rand.Intn(len(signs))]

			if sign == "-" && a < b {
				a, b = b, a
			}

			writer(strconv.Itoa(a))
			writer(sign)
			writer(strconv.Itoa(b))
			writer("=")

			pdf.Rect(pdf.GetX(), pdf.GetY(), lineHeight, lineHeight, "D")

			pdf.MoveTo(pdf.GetX()+lineHeight+2, pdf.GetY())

			pdf.Rect(pdf.GetX(), pdf.GetY(), lineHeight, lineHeight, "D")
		}
		pdf.Ln(lineHeight * 1.5)
	}

	err := pdf.OutputFileAndClose(fmt.Sprintf("%d.pdf", time.Now().Unix()))
	if err != nil {
		panic(err)
	}

}
