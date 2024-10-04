package pdf

import (
	"fmt"
	"os"
	"path"

	"github.com/jung-kurt/gofpdf"
	"training.go/pdf_certif_project/cert"
)

type PdfSaver struct {
	OutputDir string
}

func New(outputDir string) (*PdfSaver, error) {
	err := os.MkdirAll(outputDir, os.ModePerm)

	if err != nil {
		return nil, err
	}

	p := &PdfSaver{
		OutputDir: outputDir,
	}
	return p, nil
}

func (p *PdfSaver) Save(c *cert.Cert) error {
    pdf := gofpdf.New(gofpdf.OrientationLandscape, "mm", "A4","")
	pdf.SetTitle(c.LabelTitle, false)
	pdf.AddPage()

	// header
	header(pdf, c)

	// body
	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelPresented, "C")

	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.Name, "C")

	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelParticipation, "C")

	pdf.Ln(30)
	pdf.SetFont("Helvetica", "I", 20)
	pdf.WriteAligned(0, 50, c.LabelDate, "C")

	// save file

	filename := fmt.Sprintf("%v.pdf", c.LabelTitle)
	path := path.Join(p.OutputDir, filename)
	err := pdf.OutputFileAndClose(path)

	if err != nil {
		return err
	}


	fmt.Printf("My file is saved on path %v\n", path)
	return nil
}

func header(pdf *gofpdf.Fpdf, c *cert.Cert) {
	opts := gofpdf.ImageOptions {
		ImageType: "png",
	}

	margin := 10.0
	x := 0.0
	imageWidth := 30.0
	filename := "images/gopher.png"

	pdf.ImageOptions(filename, 
	     x + margin,20,
		 imageWidth, 0,
		 false, opts, 0, "")
	
	pageWidth, _ := pdf.GetPageSize()

	margin = pageWidth - imageWidth
	pdf.ImageOptions(filename, 
		x + margin,20,
		imageWidth, 0,
		false, opts, 0, "")

    pdf.SetFont("Helvetica","", 40)
	pdf.WriteAligned(0, 50, c.LabelCompletion, "C")

}