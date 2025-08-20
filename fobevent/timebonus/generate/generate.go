package main

import (
	"bytes"
	"encoding/json"
	fobeventtimebonus "github.com/unknown321/fuse/fobevent/timebonus"
	"os"
	"text/template"
)

type events []int

var temp = `package fobeventtimebonus

var FOBTimeBonuses = []TimeBonus{
{{- range .}}
	{
		ID:         {{ .ID }},
		EventID:    {{ .EventID }},
		Type:       {{ .Type }},
		SameTimeBonus:  [8]int{ 
			{{- range .SameTimeBonus }}
			{{ . }},
			{{- end }}
		},
		BonusMin:   {{ .BonusMin }},
		BonusMax:   {{ .BonusMax }},
	},
{{- end }}
}
`

type tb [8]int

type oetEvent struct {
	EventID                 int `json:"event_id"`
	EventSneakClearPointMin int `json:"event_sneak_clear_point_min"`
	EventSneakClearPointMax int `json:"event_sneak_clear_point_max"`
	EventSneak              tb  `json:"event_sneak_same_time_bonus"`
}

type fobEventTaskList struct {
	NormalDefense tb         `json:"normal_defense_same_time_bonus"`
	NormalSneak   tb         `json:"normal_sneak_same_time_bonus"`
	OneEventTask  []oetEvent `json:"one_event_param"`
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

	var ee []fobeventtimebonus.TimeBonus
	ee = append(ee, fobeventtimebonus.TimeBonus{
		ID:            0,
		EventID:       -1,
		Type:          fobeventtimebonus.NormalDefense,
		SameTimeBonus: txt.NormalDefense,
		BonusMin:      -1,
		BonusMax:      -1,
	})

	ee = append(ee, fobeventtimebonus.TimeBonus{
		ID:            0,
		EventID:       -1,
		Type:          fobeventtimebonus.NormalSneak,
		SameTimeBonus: txt.NormalSneak,
		BonusMin:      -1,
		BonusMax:      -1,
	})

	for _, v := range txt.OneEventTask {
		ee = append(ee, fobeventtimebonus.TimeBonus{
			EventID:       v.EventID,
			Type:          fobeventtimebonus.OneEventTaskSneak,
			SameTimeBonus: v.EventSneak,
			BonusMax:      v.EventSneakClearPointMax,
			BonusMin:      v.EventSneakClearPointMin,
		})
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

	if err = os.WriteFile("../timebonuses.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
