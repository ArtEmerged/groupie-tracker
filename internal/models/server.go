package models

import "html/template"

var (
	Tpl     *template.Template
	InitErr error
)

const (
	Address = "localhost"
	Port    = "8080"
)
