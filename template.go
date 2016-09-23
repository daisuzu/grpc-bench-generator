package main

import (
	"io"
	"text/template"
)

func Output(w io.Writer, bci BenchClientInfo) {
	tmpl := string(MustAsset("templates/main.go.tmpl"))
	t := template.Must(template.New("tmpl").Parse(tmpl))
	t.Execute(w, bci)
}
