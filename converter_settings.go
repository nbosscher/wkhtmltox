package wkhtmltox

import (
	"github.com/nbosscher/wkhtmltox/wkhtmltopdf"
	"strconv"
)

const (
	Portrait  Orientation = "Portrait"
	Landscape Orientation = "Landscape"
)

const (
	// based on https://github.com/wkhtmltopdf/wkhtmltopdf/blob/master/src/pdf/pdfdocparts.cc#L349
	// and http://doc.qt.io/qt-4.8/qprinter.html#PaperSize-enum
	PageSizeA0        PageSize = "A0"        // size = 841 x 1189 mm
	PageSizeA1        PageSize = "A1"        // size = 594 x 841 mm
	PageSizeA2        PageSize = "A2"        // size = 420 x 594 mm
	PageSizeA3        PageSize = "A3"        // size = 297 x 420 mm
	PageSizeA4        PageSize = "A4"        // size = 210 x 297 mm, 8.26 x 11.69 inches
	PageSizeA5        PageSize = "A5"        // size = 148 x 210 mm
	PageSizeA6        PageSize = "A6"        // size = 105 x 148 mm
	PageSizeA7        PageSize = "A7"        // size = 74 x 105 mm
	PageSizeA8        PageSize = "A8"        // size = 52 x 74 mm
	PageSizeA9        PageSize = "A9"        // size = 37 x 52 mm
	PageSizeB0        PageSize = "B0"        // size = 1000 x 1414 mm
	PageSizeB1        PageSize = "B1"        // size = 707 x 1000 mm
	PageSizeB2        PageSize = "B2"        // size = 500 x 707 mm
	PageSizeB3        PageSize = "B3"        // size = 353 x 500 mm
	PageSizeB4        PageSize = "B4"        // size = 250 x 353 mm
	PageSizeB5        PageSize = "B5"        // size = 176 x 250 mm, 6.93 x 9.84 inches
	PageSizeB6        PageSize = "B6"        // size = 125 x 176 mm
	PageSizeB7        PageSize = "B7"        // size = 88 x 125 mm
	PageSizeB8        PageSize = "B8"        // size = 62 x 88 mm
	PageSizeB9        PageSize = "B9"        // size = 33 x 62 mm
	PageSizeB10       PageSize = "B10"       // size = 31 x 44 mm
	PageSizeC5E       PageSize = "C5E"       // size = 163 x 229 mm
	PageSizeComm10E   PageSize = "Comm10E"   // size = 105 x 241 mm, U.S. Common 10 Envelope
	PageSizeDLE       PageSize = "DLE"       // size = 110 x 220 mm
	PageSizeExecutive PageSize = "Executive" // size = 7.5 x 10 inches, 190.5 x 254 mm
	PageSizeFolio     PageSize = "Folio"     // size = 210 x 330 mm
	PageSizeLedger    PageSize = "Ledger"    // size = 431.8 x 279.4 mm
	PageSizeLegal     PageSize = "Legal"     // size = 8.5 x 14 inches, 215.9 x 355.6 mm
	PageSizeLetter    PageSize = "Letter"    // size = 8.5 x 11 inches, 215.9 x 279.4 mm
	PageSizeTabloid   PageSize = "Tabloid"   // size = 279.4 x 431.8 mm
)

const (
	ColorModeColor     ColorMode = "Color"
	ColorModeGrayScale ColorMode = "Grayscale"
)

type Orientation string
type PageSize string
type ColorMode string

type MarginSetting struct {
	Top    string // e.g. 1cm, 4in...
	Bottom string
	Left   string
	Right  string
}

// Not all wkhtmltopdf settings are implemented.
// For full list see https://wkhtmltopdf.org/libwkhtmltox/pagesettings.html#pagePdfGlobal
type ConverterSettings interface {

	// sets the page orientation
	SetOrientation(Orientation)

	// sets the page size using standard sizes
	SetPageStandardSize(PageSize)

	// sets custom page dimensions using units
	// e.g. 4in, 2cm
	SetPageDimensions(w, h string)

	// sets the color mode (color or grayscale)
	SetColorMode(ColorMode)

	// sets the number that is added to all page numbers when printing headers,
	// footers and table of content.
	SetPageOffset(int)

	// sets the title of the PDF document.
	SetDocumentTitle(string)

	// sets whether or not to use loss less compression
	SetUseCompression(bool)

	// sets the margins
	SetMargins(*MarginSetting)

	// Sets the maximal DPI to use for images in the pdf document.
	SetImageDPI(int)

	// Sets the jpeg compression factor to use when producing the pdf document, e.g. "92".
	SetJpegCompression(int)

	// Sets the path of the file used to load and store cookies.
	SetCookieJar(string)
}

type pdfConverterSettings struct {
	settings *wkhtmltopdf.GlobalSettings
}

func NewPdfConverterSettings() ConverterSettings {
	return &pdfConverterSettings{
		settings: defaultSettings(),
	}
}

func defaultSettings() *wkhtmltopdf.GlobalSettings {

	set := wkhtmltopdf.NewGlobalSettings()

	set.Set("orientation", string(Landscape))
	set.Set("colorMode", "Color")
	set.Set("size.paperSize", string(PageSizeA4))

	return set
}

// sets the page orientation
func (p *pdfConverterSettings) SetOrientation(arg Orientation) {

	p.settings.Set("orientation", string(arg))
}

// sets the page size using standard sizes
func (p *pdfConverterSettings) SetPageStandardSize(arg PageSize) {

	p.settings.Set("size.paperSize", string(arg))
}

// sets custom page dimensions using units
// e.g. 4in, 2cm
func (p *pdfConverterSettings) SetPageDimensions(w, h string) {

	p.settings.Set("size.width", w)
	p.settings.Set("size.height", h)
}

// sets the color mode (color or grayscale)
func (p *pdfConverterSettings) SetColorMode(arg ColorMode) {

	p.settings.Set("colorMode", string(arg))
}

// sets the number that is added to all page numbers when printing headers,
// footers and table of content.
func (p *pdfConverterSettings) SetPageOffset(arg int) {

	p.settings.Set("pageOffset", strconv.Itoa(arg))
}

// sets the title of the PDF document.
func (p *pdfConverterSettings) SetDocumentTitle(arg string) {

	p.settings.Set("documentTitle", arg)
}

// sets whether or not to use loss less compression
func (p *pdfConverterSettings) SetUseCompression(arg bool) {

	if arg {
		p.settings.Set("useCompression", "true")
	} else {
		p.settings.Set("useCompression", "false")
	}
}

// sets the margins
func (p *pdfConverterSettings) SetMargins(arg *MarginSetting) {

	if arg.Top != "" {
		p.settings.Set("margin.top", arg.Top)
	}

	if arg.Bottom != "" {
		p.settings.Set("margin.bottom", arg.Bottom)
	}

	if arg.Left != "" {
		p.settings.Set("margin.left", arg.Left)
	}

	if arg.Right != "" {
		p.settings.Set("margin.right", arg.Right)
	}
}

// Sets the maximal DPI to use for images in the pdf document.
func (p *pdfConverterSettings) SetImageDPI(arg int) {

	p.settings.Set("imageDPI", strconv.Itoa(arg))
}

// Sets the jpeg compression factor to use when producing the pdf document, e.g. "92".
func (p *pdfConverterSettings) SetJpegCompression(arg int) {

	p.settings.Set("imageQuality", strconv.Itoa(arg))
}

// Sets the path of the file used to load and store cookies.
func (p *pdfConverterSettings) SetCookieJar(arg string) {

	p.settings.Set("load.cookieJar", arg)
}
