// wkhtmltox provides a wrapper for the C-API ./wkhtmltopdf.
//
// This package relies on wkhtmltopdf C library. See: https://github.com/wkhtmltopdf/wkhtmltopdf
package wkhtmltox

import (
	"errors"
	"github.com/nbosscher/wkhtmltox/wkhtmltopdf"
	"log"
	"strings"
	"time"
)

const (
	LoadErrorHandleMethodAbort  LoadErrorHandleMethod = "abort"
	LoadErrorHandleMethodSkip   LoadErrorHandleMethod = "skip"
	LoadErrorHandleMethodIgnore LoadErrorHandleMethod = "ignore"
)

const (
	CssMediaTypePrint  CssMediaType = iota
	CssMediaTypeScreen CssMediaType = iota
)

type LoadErrorHandleMethod string
type CssMediaType int

type Converter interface {
	AddHtml(string, SectionSettings)
	Convert() ([]byte, error)
}

type pdfConverter struct {
	converter *wkhtmltopdf.Converter
	converted bool
}

// NewPdfConverter accepts struct created with NewPdfConverterSettings or nil.
// Passing nil will use the default settings.
func NewPdfConverter(settings ConverterSettings) Converter {

	var set *pdfConverterSettings

	if settings == nil {

		set = NewPdfConverterSettings().(*pdfConverterSettings)

	} else {
		var ok bool

		set, ok = settings.(*pdfConverterSettings)
		if !ok {
			log.Panic("settings must be of type *pdfConverterSettings or nil")
		}
	}

	return &pdfConverter{
		converter: set.settings.NewConverter(),
	}
}

// AddHtml adds the contents of arg to the current document using the settings provided.
// Passing settings = nil will use the default section settings
func (p *pdfConverter) AddHtml(arg string, settings SectionSettings) {
	if p.converted {
		log.Panic("can't call .AddHtml after .Convert")
	}

	var set *sectionSettings

	if settings == nil {

		set = NewSectionSettings().(*sectionSettings)
	} else {

		var ok bool
		set, ok = settings.(*sectionSettings)
		if !ok {
			log.Panic("settings must be of type *sectionSettings or nil")
		}
	}

	p.converter.AddHtml(set.settings, arg)
}

func (p *pdfConverter) Convert() ([]byte, error) {

	p.converted = true

	defer p.converter.Destroy()

	errs := make(chan string)

	p.converter.Warning = func(c *wkhtmltopdf.Converter, arg string) {
		go func() {
			errs <- "warning: " + arg
		}()
	}

	p.converter.Error = func(c *wkhtmltopdf.Converter, arg string) {
		go func() {
			errs <- "error: " + arg
		}()
	}

	// must run in this control flow, can't run in separate go-routine because of C bindings
	status := p.converter.Convert()

	errList := []string{}
	done := false

	for !done {
		select {
		case errString := <-errs:

			errList = append(errList, errString)
		case <-time.After(1 * time.Millisecond):

			done = true
		}
	}

	if len(errList) != 0 {
		return nil, errors.New("wkhtmltopdf: " + strings.Join(errList, ",\n"))
	}

	if !status {
		return nil, errors.New("wkhtmltopdf: conversion failed")
	}

	return p.converter.OutputAsBuffer()
}
