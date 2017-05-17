package wkhtmltox

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewPdfConverter(t *testing.T) {

	conv := NewPdfConverter(nil)
	conv.AddHtml("<html><body><h1>Hello world</h1></body></html>", nil)

	data, err := conv.Convert()
	if err != nil {
		t.Fatal(err)
	}

	os.Remove("test.pdf")

	err = ioutil.WriteFile("test.pdf", data, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPdfSettings_Orientation(t *testing.T) {

	value := "Landscape"

	set := defaultSettings()
	err := set.Set("orientation", value)
	if err != nil {
		t.Fatal(err)
	}

	sz, err := set.Get("orientation")
	if err != nil {
		t.Fatal(err)
	}

	if sz != value {
		t.Fatal("expecting", value, "got", sz)
	}
}

func TestPdfSettings_ViewPortSize(t *testing.T) {

	value := "1280x800"

	set := defaultSettings()
	err := set.Set("viewportSize", value)
	if err != nil {
		t.Fatal(err)
	}

	sz, err := set.Get("viewportSize")
	if err != nil {
		t.Fatal(err)
	}

	if sz != value {
		t.Fatal("expecting", value, "got", sz)
	}
}

func TestNewPdfConverter_SectionSettings_DisableImages(t *testing.T) {

	conv := NewPdfConverter(nil)

	sectionSettings := NewSectionSettings()
	sectionSettings.SetEnableImages(false)

	html := `<html>
	<body>
		<h1>Hello world</h1>
		<img style="border: 1px solid black; max-width: 200px; min-width: 100px; min-height: 100px;" src="http://cdn2-www.dogtime.com/assets/uploads/gallery/30-impossibly-cute-puppies/impossibly-cute-puppy-8.jpg" />
	</body>
	</html>`

	conv.AddHtml(html, sectionSettings)

	data, err := conv.Convert()
	if err != nil {
		t.Fatal(err)
	}

	os.Remove("test.pdf")

	err = ioutil.WriteFile("test.pdf", data, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewPdfConverter_SectionSettings_EnableImages(t *testing.T) {

	conv := NewPdfConverter(nil)

	sectionSettings := NewSectionSettings()
	sectionSettings.SetEnableImages(true)

	html := `<html>
	<body>
		<h1>Hello world</h1>
		<img style="border: 1px solid black; max-width: 200px; min-width: 100px; min-height: 100px;" src="http://cdn2-www.dogtime.com/assets/uploads/gallery/30-impossibly-cute-puppies/impossibly-cute-puppy-8.jpg" />
	</body>
	</html>`

	conv.AddHtml(html, sectionSettings)

	data, err := conv.Convert()
	if err != nil {
		t.Fatal(err)
	}

	os.Remove("test.pdf")

	err = ioutil.WriteFile("test.pdf", data, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewPdfConverter_PageSettings(t *testing.T) {

	settings := NewPdfConverterSettings()
	settings.SetColorMode(ColorModeGrayScale)

	conv := NewPdfConverter(settings)

	html := `<html>
	<body>
		<h1>Hello world</h1>
		<img style="border: 1px solid black; max-width: 200px; min-width: 100px; min-height: 100px;" src="http://cdn2-www.dogtime.com/assets/uploads/gallery/30-impossibly-cute-puppies/impossibly-cute-puppy-8.jpg" />
	</body>
	</html>`

	conv.AddHtml(html, nil)

	data, err := conv.Convert()
	if err != nil {
		t.Fatal(err)
	}

	os.Remove("test.pdf")

	err = ioutil.WriteFile("test.pdf", data, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
