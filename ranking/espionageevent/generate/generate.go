package main

import (
	"bytes"
	"encoding/json"
	"github.com/unknown321/fuse/ranking/espionageevent"
	"os"
	"text/template"
)

type events []int

var temp = `package espionageevent

var EspionageEvents = []EspionageEvent{
{{- range .}}
	{
		ID:      {{ .ID }},
		EventID: {{ .EventID }},
	},
{{- end }}
}
`

func main() {
	data, err := os.ReadFile("values.json")
	if err != nil {
		panic(err)
	}

	var txt events
	err = json.Unmarshal(data, &txt)
	if err != nil {
		panic(err)
	}

	var ee []espionageevent.EspionageEvent
	for i, v := range txt {
		ee = append(ee, espionageevent.EspionageEvent{
			ID:      i,
			EventID: v,
		})
	}

	t, err := template.New("a").Parse(temp)
	if err != nil {
		panic(err)
	}

	o := []byte{}
	res := bytes.NewBuffer(o)
	if err = t.Execute(res, ee); err != nil {
		panic(err)
	}

	if err = os.WriteFile("../events.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
