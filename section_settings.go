package wkhtmltox

import (
	"fmt"
	"github.com/nbosscher/wkhtmltox/wkhtmltopdf"
	"strconv"
	"time"
)

// not all wkhtmltopdf arguements are implemented here.
// for full list see https://wkhtmltopdf.org/libwkhtmltox/pagesettings.html#pagePdfObject
type SectionSettings interface {

	// sets whether or not to enable javascript
	SetEnableJavascript(bool)

	// sets the amount of time to wait after the page is loaded before rendering
	SetJavascriptDelay(time.Duration)

	// sets whether or not to forward javascript warnings to Convert().error
	SetDebugJavascript(bool)

	// sets whether or not to load external or local images
	SetEnableImages(bool)

	// sets whether or not to intelligently shrink content to fit on a page
	SetEnableIntelligentShrinking(bool)

	// sets which media type to use when rendering
	SetCssMediaType(CssMediaType)

	// sets the encoding if it is not declared on the page
	SetDefaultEncoding(string)

	// sets whether or not to load local files referenced by the section
	SetLoadReferencedLocalFiles(bool)

	// sets what to do if objects fail to load
	SetLoadErrorHandling(LoadErrorHandleMethod)

	// sets the amount of space to put between the header and the content, e.g. "1.8".
	SetHeaderSpacing(float32)

	// sets the amount of space to put between the footer and the content, e.g. "1.8".
	SetFooterSpacing(float32)

	// sets whether or not external links in the HTML document are converted into external pdf links
	SetConvertExternalLinks(bool)

	// sets whether or not internal links in the HTML document are converted into internal pdf links
	SetConvertInternalLinks(bool)

	// sets whether or not to convert html forms to pdf forms
	SetConvertForms(bool)

	// sets the browser zoom factor (1.00 = 100%)
	SetZoomFactor(float32)
}

type sectionSettings struct {
	settings *wkhtmltopdf.ObjectSettings
}

func NewSectionSettings() SectionSettings {
	return &sectionSettings{
		settings: defaultObjectSettings(),
	}
}

func defaultObjectSettings() *wkhtmltopdf.ObjectSettings {
	set := wkhtmltopdf.NewObjectSettings()

	//set.Set("web.enableIntelligentShrinking", "false")
	set.Set("viewportSize", "1280x800")

	return set
}

// sets whether or not to enable javascript
func (s *sectionSettings) SetEnableJavascript(arg bool) {

	if arg {
		s.settings.Set("web.enableJavascript", "true")
	} else {
		s.settings.Set("web.enableJavascript", "false")
	}
}

// sets the amount of time to wait after the page is loaded before rendering
func (s *sectionSettings) SetJavascriptDelay(arg time.Duration) {

	ms := (arg.Nanoseconds() / 1e6)
	s.settings.Set("load.jsdelay", strconv.FormatInt(ms, 10))
}

// sets whether or not to forward javascript warnings to Convert().error
func (s *sectionSettings) SetDebugJavascript(arg bool) {

	if arg {
		s.settings.Set("load.debugJavascript", "true")
	} else {
		s.settings.Set("load.debugJavascript", "false")
	}
}

// sets whether or not to load external or local images
func (s *sectionSettings) SetEnableImages(arg bool) {

	if arg {
		s.settings.Set("web.loadImages", "true")
	} else {
		s.settings.Set("web.loadImages", "false")
	}
}

// sets whether or not to intelligently shrink content to fit on a page
func (s *sectionSettings) SetEnableIntelligentShrinking(arg bool) {

	if arg {
		s.settings.Set("web.enableIntelligentShrinking", "true")
	} else {
		s.settings.Set("web.enableIntelligentShrinking", "false")
	}
}

// sets which media type to use when rendering
func (s *sectionSettings) SetCssMediaType(arg CssMediaType) {

	switch arg {
	case CssMediaTypePrint:
		s.settings.Set("web.printMediaType", "true")
	case CssMediaTypeScreen:
		s.settings.Set("web.printMediaType", "false")
	}
}

// sets the encoding if it is not declared on the page
func (s *sectionSettings) SetDefaultEncoding(arg string) {

	s.settings.Set("web.defaultEncoding", arg)
}

// sets whether or not to load local files referenced by the section
func (s *sectionSettings) SetLoadReferencedLocalFiles(arg bool) {

	if arg {
		s.settings.Set("web.blockLocalFileAccess", "false")
	} else {
		s.settings.Set("web.blockLocalFileAccess", "true")
	}
}

// sets what to do if objects fail to load
func (s *sectionSettings) SetLoadErrorHandling(arg LoadErrorHandleMethod) {

	s.settings.Set("load.loadErrorHandling", string(arg))
}

// sets the amount of space to put between the header and the content, e.g. "1.8".
func (s *sectionSettings) SetHeaderSpacing(arg float32) {

	s.settings.Set("header.spacing", fmt.Sprintf("%.2f", arg))
}

// sets the amount of space to put between the footer and the content, e.g. "1.8".
func (s *sectionSettings) SetFooterSpacing(arg float32) {

	s.settings.Set("footer.spacing", fmt.Sprintf("%.2f", arg))
}

// sets whether or not external links in the HTML document are converted into external pdf links
func (s *sectionSettings) SetConvertExternalLinks(arg bool) {

	if arg {
		s.settings.Set("useExternalLinks", "true")
	} else {
		s.settings.Set("useExternalLinks", "false")
	}
}

// sets whether or not internal links in the HTML document are converted into internal pdf links
func (s *sectionSettings) SetConvertInternalLinks(arg bool) {

	if arg {
		s.settings.Set("useLocalLinks", "true")
	} else {
		s.settings.Set("useLocalLinks", "false")
	}
}

// sets whether or not to convert html forms to pdf forms
func (s *sectionSettings) SetConvertForms(arg bool) {

	if arg {
		s.settings.Set("produceForms", "true")
	} else {
		s.settings.Set("produceForms", "false")
	}
}

// sets the browser zoom factor (1.00 = 100%)
func (s *sectionSettings) SetZoomFactor(arg float32) {

	s.settings.Set("load.zoomFactor", fmt.Sprintf("%.2f", arg))
}
