
# wkhtmltox

This package forms cgo bindings the wkhtmltopdf c library to provide html-to-pdf conversion

## Attribution
Most of the code in ./wkhtmltopdf was written by @jimmyw [https://github.com/jimmyw/wkhtmltopdf-go](https://github.com/jimmyw/wkhtmltopdf-go)

## Getting Started

Depends on wkhtmltopdf library. Get it at http://wkhtmltopdf.org/downloads.html
- Currently using a patched version to allow viewportSize parameter https://github.com/wkhtmltopdf/wkhtmltopdf/pull/3440

#### Basic Case
```golang

// create a converter
// - nil means use default page settings
conv := NewPdfConverter(nil)

// add content
// - nil means use default section settings
conv.AddHtml("<html><body><h1>Hello world</h1></body></html>", nil)

// do the conversion
pdfData, err := conv.Convert()
if err != nil {
    t.Fatal(err)
}
```

#### Page Settings
```golang

// create page settings with default settings
pageSettings := NewPdfConverterSettings()

// customize settings (see converter_settings.go for more props)
pageSettings.SetColorMode(ColorModeGrayScale)

// create converter using our custom settings
conv := NewPdfConverter(pageSettings)

// continue as before..
conv.AddHtml("<html><body><h1>Hello world</h1></body></html>", nil)

pdfData, err := conv.Convert()
if err != nil {
    t.Fatal(err)
}
```

#### Section Settings
```golang

// create converter with default page settings
conv := NewPdfConverter(nil)

// create section settings with default settings
sectionSettings := NewSectionSettings()

// customize settings (see section_settings.go for more props)
sectionSettings.SetEnableImages(false)

// add section with our custom section settings
conv.AddHtml("<html><body><h1>Hello world</h1></body></html>", sectionSettings)

pdfData, err := conv.Convert()
if err != nil {
    t.Fatal(err)
}
```