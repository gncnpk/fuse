package main

import (
	"bytes"
	"encoding/json"
	"github.com/unknown321/fuse/informationmessage"
	"os"
	"text/template"
)

type texts []im

var temp = `package informationmessage

var InformationMessages = []InformationMessage{
{{- range .}}
	{
		ID:        {{ .ID }},
		InfoID:    {{ .InfoID }},
		Date:      {{ .Date }},
		Important: {{ .Important }},
		MesBody:   ` + "`" + `{{ .MesBody }}` + "`" + `,
		MesSubject:` + "`" + `{{ .MesSubject}}` + "`" + `,
		Language:  "{{ .Language }}",
		Region:    "{{ .Region}}",
	},
{{- end }}
}
`

type im struct {
	InfoID     int    `json:"info_id"`
	Date       int    `json:"date"`
	Important  string `json:"important"`
	MesBody    string `json:"mes_body"`
	MesSubject string `json:"mes_subject"`
}

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

	var msgs []informationmessage.InformationMessage
	for i, v := range txt {
		imp := false
		if v.Important != "FALSE" {
			imp = true
		}
		mm := informationmessage.InformationMessage{
			ID:         i,
			InfoID:     v.InfoID,
			Date:       v.Date,
			Important:  imp,
			MesBody:    v.MesBody,
			MesSubject: v.MesSubject,
			Language:   "EN",
			Region:     "REGION_NA",
		}
		msgs = append(msgs, mm)
	}

	t, err := template.New("a").Parse(temp)
	if err != nil {
		panic(err)
	}

	o := []byte{}
	res := bytes.NewBuffer(o)
	if err = t.Execute(res, msgs); err != nil {
		panic(err)
	}

	if err = os.WriteFile("../messages.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
