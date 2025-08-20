package main

import (
	"bytes"
	"encoding/json"
	"github.com/unknown321/fuse/serverproductparam"
	"os"
	"text/template"
)

type params []serverproductparam.ServerProductParam

var temp = `package serverproductparam

var ServerProductParams = []ServerProductParam{
{{- range .}}
	{
		DevCoin:            {{.DevCoin }},
		DevGmp: 			{{ .DevGmp }},
		DevItem1:           {{ .DevItem1 }},
		DevItem2:           {{ .DevItem2 }},
		DevPlatlv01:        {{ .DevPlatlv01 }},
		DevPlatlv02:        {{ .DevPlatlv02 }},
		DevPlatlv03:        {{ .DevPlatlv03 }},
		DevPlatlv04:        {{ .DevPlatlv04 }},
		DevPlatlv05:        {{ .DevPlatlv05 }},
		DevPlatlv06:        {{ .DevPlatlv06 }},
		DevPlatlv07:        {{ .DevPlatlv07 }},
		DevRescount01Value: {{ .DevRescount01Value }},
		DevRescount02Value: {{ .DevRescount02Value }},
		DevResource01Id:    {{ .DevResource01Id }},
		DevResource02Id:    {{ .DevResource02Id }},
		DevSkil:            {{ .DevSkil }},
		DevSpecial:         {{ .DevSpecial }},
		DevTime:            {{ .DevTime }},
		ID:                 {{ .ID }},
		Type:               {{ .Type }},
		UseGmp:             {{ .UseGmp }},
		UseRescount01Value: {{ .UseRescount01Value }},
		UseRescount02Value: {{ .UseRescount02Value }},
		UseResource01Id:    {{ .UseResource01Id }},
		UseResource02Id:    {{ .UseResource02Id }},
	},
{{- end }}
}
`

func main() {
	data, err := os.ReadFile("values.json")
	if err != nil {
		panic(err)
	}

	var pp params
	err = json.Unmarshal(data, &pp)
	if err != nil {
		panic(err)
	}

	t, err := template.New("a").Parse(temp)
	if err != nil {
		panic(err)
	}

	o := []byte{}
	res := bytes.NewBuffer(o)
	if err = t.Execute(res, pp); err != nil {
		panic(err)
	}

	if err = os.WriteFile("../params.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
