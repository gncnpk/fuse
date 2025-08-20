package sessionmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/playerresource"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
)

type ProcessingRates struct {
	Raw       playerresource.Raw
	Processed playerresource.Processed
}

func GetCmdMiningResourceResponse(ctx context.Context, msg *message.Message, manager *SessionManager) tppmessage.CmdMiningResourceResponse {
	t := tppmessage.CmdMiningResourceResponse{}
	t.CryptoType = tppmessage.CRYPTO_TYPE_COMPOUND
	t.Msgid = tppmessage.CMD_MINING_RESOURCE.String()
	t.Result = tppmessage.RESULT_NOERR
	t.Rqid = 0

	// TODO to/from database
	t.BreakTime = 2100

	onlineRes, err := manager.PlayerResourceRepo.Get(ctx, msg.PlayerID, true)
	if err != nil {
		slog.Error("get online resources", "error", err.Error(), "playerID", msg.PlayerID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	if len(onlineRes) != 1 {
		slog.Error("get online resources", "not enough values, want 1, got", len(onlineRes), "playerID", msg.PlayerID)
		t.Result = tppmessage.RESULT_ERR
		return t
	}

	// before mining and after mining are 0

	r := onlineRes[0]

	// TODO what this parameter is for, how much resources are processed? must be in database or calculated based on fob development progress?
	// looks like amount of resources from online mining
	TEMPVALUE := 10000

	rates := ProcessingRates{
		Raw: playerresource.Raw{
			Fuel:          TEMPVALUE,
			Bio:           TEMPVALUE,
			CommonMetal:   TEMPVALUE,
			MinorMetal:    TEMPVALUE,
			PreciousMetal: TEMPVALUE,
		},
		Processed: playerresource.Processed{
			Fuel:          TEMPVALUE,
			Bio:           TEMPVALUE,
			CommonMetal:   TEMPVALUE,
			MinorMetal:    TEMPVALUE,
			PreciousMetal: TEMPVALUE,
		},
	}

	// raw
	{
		t.BeforeProcessResource.FuelResource = r.Raw.Fuel
		t.BeforeProcessResource.BioticResource = r.Raw.Bio
		t.BeforeProcessResource.CommonMetal = r.Raw.CommonMetal
		t.BeforeProcessResource.MinorMetal = r.Raw.MinorMetal
		t.BeforeProcessResource.PreciousMetal = r.Raw.PreciousMetal

		t.AfterProcessResource.FuelResource = t.BeforeProcessResource.FuelResource - rates.Raw.Fuel
		t.AfterProcessResource.BioticResource = t.BeforeProcessResource.BioticResource - rates.Raw.Bio
		t.AfterProcessResource.CommonMetal = t.BeforeProcessResource.CommonMetal - rates.Raw.CommonMetal
		t.AfterProcessResource.MinorMetal = t.BeforeProcessResource.MinorMetal - rates.Raw.MinorMetal
		t.AfterProcessResource.PreciousMetal = t.BeforeProcessResource.PreciousMetal - rates.Raw.PreciousMetal

		t.BeforeMiningResource.FixLimit(0, 1000000)
		t.AfterMiningResource.FixLimit(0, 1000000)
	}

	// processed
	{
		t.BeforeProcessResourceUsable.FuelResource = r.Processed.Fuel
		t.BeforeProcessResourceUsable.BioticResource = r.Processed.Bio
		t.BeforeProcessResourceUsable.CommonMetal = r.Processed.CommonMetal
		t.BeforeProcessResourceUsable.MinorMetal = r.Processed.MinorMetal
		t.BeforeProcessResourceUsable.PreciousMetal = r.Processed.PreciousMetal

		t.AfterProcessResourceUsable.FuelResource = t.BeforeProcessResourceProcessing.FuelResource + rates.Processed.Fuel
		t.AfterProcessResourceUsable.BioticResource = t.BeforeProcessResourceProcessing.BioticResource + rates.Processed.Bio
		t.AfterProcessResourceUsable.CommonMetal = t.BeforeProcessResourceProcessing.CommonMetal + rates.Processed.CommonMetal
		t.AfterProcessResourceUsable.MinorMetal = t.BeforeProcessResourceProcessing.MinorMetal + rates.Processed.MinorMetal
		t.AfterProcessResourceUsable.PreciousMetal = t.BeforeProcessResourceProcessing.PreciousMetal + rates.Processed.PreciousMetal

		t.BeforeProcessResourceProcessing.FixLimit(0, 1000000)
		t.AfterMiningResourceProcessing.FixLimit(0, 1000000)
	}

	return t
}

func HandleCmdMiningResourceResponse(message *message.Message, override bool) error {
	if !override {
		return nil
	}

	slog.Info("using overridden version")
	var err error
	//t := GetCmdMiningResourceResponse()
	t := tppmessage.CmdMiningResourceResponse{}

	message.MData, err = json.Marshal(t)
	if err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
