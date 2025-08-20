package main

import (
	"bytes"
	"encoding/json"
	fobeventreward "github.com/unknown321/fuse/fobevent/reward"
	"os"
	"text/template"
)

type events []int

var temp = `package fobeventreward 

var FOBRewards = []Reward{
{{- range .}}
	{
		ID:         {{ .ID }},
		EventID:    {{ .EventID }},
		Type:       {{ .Type }},
		Reward:     {{ .Reward }},
		TaskTypeID: {{ .TaskTypeID }}, 
		Threshold:  {{ .Threshold }},
	},
{{- end }}
}
`

type oetEvent struct {
	EventID    int                     `json:"event_id"`
	EventSneak []fobeventreward.Reward `json:"event_sneak"`
}

type fobEventTaskList struct {
	NormalDefense []fobeventreward.Reward `json:"normal_defense"`
	NormalSneak   []fobeventreward.Reward `json:"normal_sneak"`
	OneEventTask  []oetEvent              `json:"one_event_task"`
}

func main() {
	data, err := os.ReadFile("values.json")
	if err != nil {
		panic(err)
	}

	var txt fobEventTaskList
	err = json.Unmarshal(data, &txt)
	if err != nil {
		panic(err)
	}

	var ee []fobeventreward.Reward
	for _, v := range txt.NormalDefense {
		ee = append(ee, fobeventreward.Reward{
			EventID:    -1,
			Type:       fobeventreward.NormalDefense,
			Reward:     v.Reward,
			TaskTypeID: v.TaskTypeID,
			Threshold:  v.Threshold,
		})
	}

	for _, v := range txt.NormalSneak {
		ee = append(ee, fobeventreward.Reward{
			EventID:    -1,
			Type:       fobeventreward.NormalSneak,
			Reward:     v.Reward,
			TaskTypeID: v.TaskTypeID,
			Threshold:  v.Threshold,
		})
	}

	for _, v := range txt.OneEventTask {
		for _, r := range v.EventSneak {
			ee = append(ee, fobeventreward.Reward{
				EventID:    v.EventID,
				Type:       fobeventreward.OneEventTaskSneak,
				Reward:     r.Reward,
				TaskTypeID: r.TaskTypeID,
				Threshold:  r.Threshold,
			})
		}
	}

	for i := range ee {
		ee[i].ID = i
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

	if err = os.WriteFile("../rewards.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
