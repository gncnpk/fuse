package main

import (
	"bytes"
	"encoding/json"
	"github.com/unknown321/fuse/servertext"
	"os"
	"text/template"
)

type texts []servertext.ServerText

var temp = `package servertext

var ServerTexts = []ServerText{
{{- range .}}
	{
		ID:         {{ .ID }},
		Identifier: "{{ .Identifier }}",
		Language:   "{{ .Language }}",
		Text:       ` + "`" + `{{ .Text }}` + "`" + `,
	},
{{- end }}
}
`

func main() {
	data, err := os.ReadFile("values.json")
	if err != nil {
		panic(err)
	}

	var txt texts
	err = json.Unmarshal(data, &txt)
	if err != nil {
		panic(err)
	}

	for i := range txt {
		txt[i].ID = i
	}

	t, err := template.New("a").Parse(temp)
	if err != nil {
		panic(err)
	}

	o := []byte{}
	res := bytes.NewBuffer(o)
	if err = t.Execute(res, txt); err != nil {
		panic(err)
	}

	if err = os.WriteFile("../texts.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
