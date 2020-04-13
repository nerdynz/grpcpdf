package main

import (
	"context"
	"net"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/nerdynz/trove"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/nerdynz/jeevesrpcpdf/makepdf"
)

func main() {
	settings := trove.Load()
	lis, err := net.Listen("tcp", ":"+settings.GetWithDefault("GRPC_PORT", "5534"))
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	s := grpc.NewServer()
	makepdf.RegisterMakePDFServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		logrus.Error(err)
		panic(err)
	}
}

type server struct {
}

func (s *server) GetPDFFromURL(ctx context.Context, params *makepdf.PDFParams) (*makepdf.PDFFile, error) {
	url := params.GetUrl()

	pdfg, err := pdf.NewPDFGenerator()
	if err != nil {
		logrus.Error("Failed on => ", url)
		return nil, err
	}

	pdfg.Dpi.Set(300)
	pdfg.NoCollate.Set(false)
	pdfg.PageSize.Set(pdf.PageSizeA4)

	if params.GetIsMarginless() {
		pdfg.MarginTop.Set(0)
		pdfg.MarginLeft.Set(0)
		pdfg.MarginBottom.Set(0)
		pdfg.MarginRight.Set(0)
	}

	if params.GetIsLandscape() {
		pdfg.Orientation.Set("landscape")
	}

	page := pdf.NewPage(url)
	if params.GetIsDebug() {
		page.DebugJavascript.Set(true)
	}

	readyFlag := params.GetJavascriptReadyFlag()
	if readyFlag == "" {
		delay := params.GetDelay()
		if delay == 0 {
			page.JavascriptDelay.Set(250)
		} else {
			page.JavascriptDelay.Set(uint(delay))
		}
	} else {
		page.WindowStatus.Set(readyFlag)
	}
	page.NoStopSlowScripts.Set(true)
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		logrus.Error("Failed on => ", url)
		return nil, err
	}

	bts := pdfg.Bytes()
	file := &makepdf.PDFFile{}
	file.Binary = bts
	return file, nil
}
