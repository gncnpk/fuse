package main

import (
	"bytes"
	"encoding/json"
	"github.com/unknown321/fuse/onlinechallengetask"
	"os"
	"text/template"
)

type events []int

var temp = `package onlinechallengetask

var OnlineChallengeTasks = []OnlineChallengeTask{
{{- range .}}
	{
		ID:               {{ .ID }},
		MissionID:        {{ .MissionID }},
		RewardBottomType: {{ .RewardBottomType }},
		RewardRate:       {{ .RewardRate }},
		RewardSection:    {{ .RewardSection }},
		RewardType:       {{ .RewardType }},
		RewardValue:      {{ .RewardValue }},
		TaskTypeID:       {{ .TaskTypeID }},
		Threshold:        {{ .Threshold }},
		EndDate:          {{ .EndDate }},
		Version:          {{ .Version }},
	},
{{- end }}
}
`

type reward struct {
	BottomType int `json:"bottom_type"`
	Rate       int `json:"rate"`
	Section    int `json:"section"`
	Type       int `json:"type"`
	Value      int `json:"value"`
}

type task struct {
	MissionID  int    `json:"mission_id"`
	Reward     reward `json:"reward"`
	Status     int    `json:"status"`
	TaskTypeID int    `json:"task_type_id"`
	Threshold  int    `json:"threshold"`
}

type onlineTasks struct {
	EndDate  int    `json:"end_date"`
	Version  int    `json:"version"`
	TaskList []task `json:"task_list"`
}

func main() {
	data, err := os.ReadFile("values.json")
	if err != nil {
		panic(err)
	}

	var txt onlineTasks
	err = json.Unmarshal(data, &txt)
	if err != nil {
		panic(err)
	}

	var ee []onlinechallengetask.OnlineChallengeTask

	for i, v := range txt.TaskList {
		ot := onlinechallengetask.OnlineChallengeTask{
			ID:               i,
			MissionID:        v.MissionID,
			RewardBottomType: v.Reward.BottomType,
			RewardRate:       v.Reward.Rate,
			RewardSection:    v.Reward.Section,
			RewardType:       v.Reward.Type,
			RewardValue:      v.Reward.Value,
			TaskTypeID:       v.TaskTypeID,
			Threshold:        v.Threshold,
			EndDate:          txt.EndDate,
			Version:          txt.Version,
		}

		ee = append(ee, ot)
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

	if err = os.WriteFile("../tasks.go", res.Bytes(), 0644); err != nil {
		panic(err)
	}
}
