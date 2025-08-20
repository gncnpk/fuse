package sessionmanager

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/sectionstat"
	"github.com/unknown321/fuse/tppmessage"
	"io"
	"log/slog"
)

type SollyDataClient struct {
	Padding1        uint8 // always 0
	DirectContract  uint8 // direct contract
	Padding2        uint8 // always 0
	MaybeGene       uint8 // values from 0 to 9, numbers look like amount of staff per platform (but there are only 7 platforms, brig+medi?)
	Padding3        uint8 // always 0
	MaybeBaseStats  uint8 // values 0, 3-9, probably something about stats, I have a lot of 9 (~2700) - base stats?
	Padding4        uint8 // always 0
	MaybeBaseStats2 uint8 // values 0, 3-9, numbers look like 5th byte, but slightly different - base stats+buffs?
}

type SollyDataServer struct {
	Header [4]byte
	Data   [12]byte
}

type Solly struct {
	DataClient SollyDataClient
	DataServer SollyDataServer
}

//seeds
//F0E1C000:Amber Fox
//FD4E2000:Code Talker
//FD866000:Code Talker
//F36E2000:Crimson Canine
//F6000000:Emerald Hound
//E35EB000:Eye
//E49B0000:Finger
//F417F000:Garnet Canine
//EFB4A000:Gold Fox
//6E545000:Hideo
//6E545800:Hideo
//FC154000:Huey
//FCA28000:Huey
//F2000000:Ivory Skull
//F9000000:Miller
//FA5D0000:Miller
//FA7C5000:Miller
//FA987000:Miller
//FAA1E800:Miller
//FAF62000:Miller
//F9000000:Ocelot
//F90A6000:Ocelot
//F9AE8800:Ocelot
//F9ED2000:Ocelot
//FB19E000:Quiet
//FB2B9800:Quiet
//FBB76000:Quiet
//4EBA2800:Rat
//11B23000:Silent Basilisk
//F1000000:Silver Skull
//F7000000:Snake
//F8000000:Snake
//F5000000:Viridian Hound
//6F3B7800:Ziang Tan

var HeaderMap = map[[4]byte]string{
	[4]byte{0x2A, 0xD0, 0x1F, 0x80}: "Amber Fox",
	[4]byte{0x00, 0x07, 0xDF, 0x80}: "Code Talker",
	[4]byte{0x2A, 0xB0, 0x1F, 0x80}: "Crimson Canine",
	[4]byte{0x2A, 0x80, 0x1F, 0x80}: "Emerald Hound",
	[4]byte{0x27, 0x30, 0x1F, 0x80}: "Eye",
	[4]byte{0x27, 0x40, 0x1F, 0x80}: "Finger",
	[4]byte{0x00, 0x06, 0xDF, 0x80}: "Huey",
	[4]byte{0x2A, 0xA0, 0x1F, 0x80}: "Garnet Canine",
	[4]byte{0x2A, 0xC0, 0x1F, 0x80}: "Gold Fox",
	[4]byte{0x26, 0xE8, 0x5F, 0x80}: "Hideo",
	[4]byte{0x2A, 0xF0, 0x1F, 0x80}: "Ivory Skull",
	[4]byte{0x00, 0x08, 0x9F, 0x80}: "Miller",
	[4]byte{0x2A, 0x40, 0x1F, 0x80}: "Miller",
	[4]byte{0x00, 0x08, 0x7F, 0x80}: "Ocelot",
	[4]byte{0x2A, 0x50, 0x1F, 0x80}: "Ocelot",
	[4]byte{0x00, 0x01, 0xBF, 0x80}: "Quiet",
	[4]byte{0x46, 0x98, 0x5F, 0x80}: "Rat",
	[4]byte{0x40, 0x08, 0x5F, 0x80}: "Silent Basilisk",
	[4]byte{0x2A, 0xE0, 0x1F, 0x80}: "Silver Skull",
	[4]byte{0x2A, 0x60, 0x1F, 0x80}: "Snake",
	[4]byte{0x2A, 0x70, 0x1F, 0x80}: "Snake",
	[4]byte{0x2A, 0x90, 0x1F, 0x80}: "Viridian Hound",
	[4]byte{0xA6, 0xD1, 0x5F, 0x80}: "Ziang Tan",
}

func ReadSoldierData(soldierParams string) ([]Solly, error) {
	clientParams, err := base64.StdEncoding.DecodeString(soldierParams)
	if err != nil {
		return nil, fmt.Errorf("cannot decode client params: %w", err)
	}

	r := bytes.NewReader(clientParams)

	out := []Solly{}
	for {
		s := Solly{}
		if err = binary.Read(r, binary.LittleEndian, &s); err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				return nil, fmt.Errorf("cannot read solly: %w", err)
			}
		}
		out = append(out, s)
		var name string
		var ok bool
		if name, ok = HeaderMap[s.DataServer.Header]; !ok {
			name = "?"
		}
		slog.Debug("sol",
			"contr", s.DataClient.DirectContract,
			"gene", s.DataClient.MaybeGene,
			"stat1", s.DataClient.MaybeBaseStats,
			"stat2", s.DataClient.MaybeBaseStats2,
			"name", name,
			"header", fmt.Sprintf("%02x", s.DataServer.Header),
			"data", fmt.Sprintf("% 02x", s.DataServer.Data),
		)
	}

	return out, nil
}

func WriteServerSoldierData(soldiers []Solly) string {
	soldierParam := []byte{}

	for _, s := range soldiers {
		soldierParam = append(soldierParam, s.DataServer.Header[:]...)
		soldierParam = append(soldierParam, s.DataServer.Data[:]...)
	}

	slog.Debug("wrote soldiers", "count", len(soldiers))

	out := base64.StdEncoding.EncodeToString(soldierParam)
	return out
}

func HandleCmdSyncSoldierBinRequest(ctx context.Context, message *message.Message, manager *SessionManager) error {
	var err error
	t := tppmessage.CmdSyncSoldierBinRequest{}
	err = json.Unmarshal(message.MData, &t)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}

	slog.Debug("CMD_SYNC_SOLDIER_BIN", "soldier_num request", t.SoldierNum)

	data := FromJSON(ctx, t.Msgid)
	if data != nil {
		message.MData = data
		message.Compress = true
		return nil
	}

	var soldiers []Solly
	soldiers, err = ReadSoldierData(t.SoldierParam)
	if err != nil {
		return err
	}
	slog.Debug("soldier param", "response count", len(soldiers))

	// TODO parse, save to database

	d := GetCmdSyncSoldierBinResponse()

	{
		stats := []sectionstat.SectionStat{
			{PlayerID: message.PlayerID, SectionID: sectionstat.Base, Level: t.Section.Base, SoldierNum: t.SectionSoldier.Base},
			{PlayerID: message.PlayerID, SectionID: sectionstat.Combat, Level: t.Section.Combat, SoldierNum: t.SectionSoldier.Combat},
			{PlayerID: message.PlayerID, SectionID: sectionstat.Develop, Level: t.Section.Develop, SoldierNum: t.SectionSoldier.Develop},
			{PlayerID: message.PlayerID, SectionID: sectionstat.Medical, Level: t.Section.Medical, SoldierNum: t.SectionSoldier.Medical},
			{PlayerID: message.PlayerID, SectionID: sectionstat.Security, Level: t.Section.Security, SoldierNum: t.SectionSoldier.Security},
			{PlayerID: message.PlayerID, SectionID: sectionstat.Spy, Level: t.Section.Spy, SoldierNum: t.SectionSoldier.Spy},
			{PlayerID: message.PlayerID, SectionID: sectionstat.Support, Level: t.Section.Suport, SoldierNum: t.SectionSoldier.Suport},
		}

		for _, v := range stats {
			if err = manager.SectionStatRepo.AddOrUpdate(ctx, &v); err != nil {
				slog.Error("section stat", "error", err.Error(), "msgid", t.Msgid)
				d.Result = tppmessage.RESULT_ERR
				goto marshalAndReturn
			}
		}
	}

	//if len(soldiers) != t.SoldierNum {
	//	return fmt.Errorf("invalid soldier count, want %d, got %d", t.SoldierNum, len(soldiers))
	//}

	d.SoldierParam = WriteServerSoldierData(soldiers)
	d.SoldierNum = t.SoldierNum
	//d.SoldierParam = string(bytes.ReplaceAll([]byte(d.SoldierParam), []byte("+"), []byte("%2B")))
	d.Version = t.Version + 1

marshalAndReturn:
	message.MData, err = json.Marshal(d)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	message.Compress = true

	return nil
}
