package main

import (
	"bytes"
	"encoding/json"
	"os"
	"text/template"
)

type T struct {
	CreateDate int `json:"create_date"`
	Develop    int `json:"develop"`
	Gmp        int `json:"gmp"`
	Id         int `json:"id"`
	LeftSecond int `json:"left_second"`
	MaxSecond  int `json:"max_second"`
	MbCoin     int `json:"mb_coin"`
	Open       int `json:"open"`
}

type params []T

var temp = `

{{- range .}}
	{
	ProductID: {{ .Id }},
	PlayerID   : pid,
	CreateDate : {{ .CreateDate }},
	Develop    : {{ .Develop    }},
	LeftSecond : {{ .LeftSecond }},
	MbCoin     : {{ .MbCoin     }},
	Open       : {{ .Open       }},
	},
{{- end }}
`

// generates personalized values for seeding
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
