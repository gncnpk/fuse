package main

import (
	"bytes"
	"encoding/json"
	"github.com/unknown321/fuse/staffrankrate"
	"github.com/unknown321/fuse/tppmessage"
	"os"
	"text/template"
)

// jaq ".staff_rank_bonus_rates" CMD_GET_LOGIN_PARAM_response

var temp = `package staffrankrate

var StaffRankBonusRates = []StaffRankBonusRate{
{{- range . }}
	{
		ID:       {{ .ID }},
		Grade:    {{ .Grade }},
		Negative: {{ .Negative }},
		Positive: {{ .Positive }},
	},
{{- end }}
}`

func main() {
	data, err := os.ReadFile("values.json")
	if err != nil {
		panic(err)
	}

	var rates []tppmessage.StaffRankBonusRate
	err = json.Unmarshal(data, &rates)
	if err != nil {
		panic(err)
	}

	var rr []staffrankrate.StaffRankBonusRate
	for grade, v := range rates {
		g := staffrankrate.StaffRankBonusRate{}
		g.ID = grade
		g.Grade = grade
		g.Negative = v.Rates[0]
		g.Positive = v.Rates[1]
		rr = append(rr, g)
	}

	t, err := template.New("a").Parse(temp)
	if err != nil {
		panic(err)
	}

	o := []byte{}
	res := bytes.NewBuffer(o)
	if err = t.Execute(res, rr); err != nil {
		panic(err)
	}

	if err = os.WriteFile("../rates.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
