package main

import (
	"bytes"
	"encoding/json"
	"github.com/unknown321/fuse/clusterbuildcost"
	"os"
	"text/template"
)

type ClusterBuildCosts struct {
	ClusterBuildCosts []clusterbuildcost.ClusterBuildCost `json:"cluster_build_costs"`
}

type ClusterBuildCostsPerGrade struct {
	ClusterBuildCostsPerGrade []ClusterBuildCosts `json:"cluster_build_costs_per_grade"`
}

var temp = `package clusterbuildcost

var ClusterBuildCosts = []ClusterBuildCost{
{{- range .}}
{{- range .ClusterBuildCostsPerGrade }}
{{- range .ClusterBuildCosts }}
	{
		ID:             {{ .ID }},
		IDX:            {{ .IDX }},
		FOBNumber:      {{ .FOBNumber }},
		Grade:          {{ .Grade }},
		Gmp:            {{ .Gmp }},
		ResourceACount: {{ .ResourceACount }},
		ResourceAID:    {{ .ResourceAID }},
		ResourceBCount: {{ .ResourceBCount }},
		ResourceBID:    {{ .ResourceBID }},
		TimeMinute:     {{ .TimeMinute }},
	},
{{- end }}
{{- end }}
{{- end }}
}
`

func main() {
	data, err := os.ReadFile("values.json")
	if err != nil {
		panic(err)
	}

	var fobs []ClusterBuildCostsPerGrade
	err = json.Unmarshal(data, &fobs)
	if err != nil {
		panic(err)
	}

	id := 0
	for fobNumber, fob := range fobs {
		for gradeNum, costGroup := range fob.ClusterBuildCostsPerGrade {
			for k := range costGroup.ClusterBuildCosts {
				fobs[fobNumber].ClusterBuildCostsPerGrade[gradeNum].ClusterBuildCosts[k].ID = id
				id++
				fobs[fobNumber].ClusterBuildCostsPerGrade[gradeNum].ClusterBuildCosts[k].Grade = gradeNum
				fobs[fobNumber].ClusterBuildCostsPerGrade[gradeNum].ClusterBuildCosts[k].FOBNumber = fobNumber
				fobs[fobNumber].ClusterBuildCostsPerGrade[gradeNum].ClusterBuildCosts[k].IDX = k
			}
		}
	}

	t, err := template.New("a").Parse(temp)
	if err != nil {
		panic(err)
	}

	o := []byte{}
	res := bytes.NewBuffer(o)
	if err = t.Execute(res, fobs); err != nil {
		panic(err)
	}

	if err = os.WriteFile("../costs.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
