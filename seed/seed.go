package seed

import (
	"context"
	"encoding/json"
	"github.com/unknown321/fuse/emblem"
	"github.com/unknown321/fuse/fobrecord"
	"github.com/unknown321/fuse/fobstatus"
	"github.com/unknown321/fuse/motherbaseparam"
	"github.com/unknown321/fuse/pfranking"
	"github.com/unknown321/fuse/pfseason"
	"github.com/unknown321/fuse/player"
	"github.com/unknown321/fuse/playerresource"
	"github.com/unknown321/fuse/playerstatus"
	"github.com/unknown321/fuse/serveritem"
	serverproductparamplayer "github.com/unknown321/fuse/serverproductparam/player"
	"github.com/unknown321/fuse/sessionmanager"
	"github.com/unknown321/fuse/user"
	"log/slog"
)

type Entry struct {
	User                     user.User
	Player                   []player.Player
	ServerItem               []serveritem.ServerItem
	ServerProductParamPlayer []serverproductparamplayer.ServerProductParamPlayer
	FobRecord                fobrecord.FobRecord
	FobStatus                fobstatus.FobStatus
	PlayerStatus             playerstatus.PlayerStatus
	PlayerResource           []playerresource.PlayerResource
	PFSeason                 []pfseason.Season
	PFranking                []pfranking.Ranking
	Emblem                   []emblem.Emblem
	MotherBaseParam          []motherbaseparam.MotherBaseParam
}

type File struct {
	Entries []Entry `json:"entries"`
}

func Read(data []byte) (File, error) {
	f := File{}
	if err := json.Unmarshal(data, &f); err != nil {
		return f, err
	}

	return f, nil
}

func (f *File) Seed(manager *sessionmanager.SessionManager) error {
	ctx := context.Background()
	var err error

	total := 0
	if total, err = manager.UserRepo.Count(ctx); err != nil {
		return err
	}
	if total > 0 {
		slog.Info("already seeded")
		return nil
	}

	for _, e := range f.Entries {
		if err = manager.UserRepo.AddUser(ctx, &e.User); err != nil {
			return err
		}

		for _, p := range e.Player {
			if err = manager.PlayerRepo.AddPlayer(ctx, &p); err != nil {
				return err
			}
		}

		for _, s := range e.ServerItem {
			if err = manager.ServerItemRepo.Add(ctx, &s); err != nil {
				return err
			}
		}

		for _, s := range e.ServerProductParamPlayer {
			if err = manager.ServerProductParamPlayerRepo.Add(ctx, &s); err != nil {
				return err
			}
		}

		if err = manager.FobRecordRepo.Add(ctx, &e.FobRecord); err != nil {
			return err
		}

		if err = manager.FobStatusRepo.Add(ctx, &e.FobStatus); err != nil {
			return err
		}

		for _, v := range e.PlayerResource {
			if err = manager.PlayerResourceRepo.AddOrUpdate(ctx, &v); err != nil {
				return err
			}
		}

		if err = manager.PlayerStatusRepo.AddOrUpdate(ctx, &e.PlayerStatus); err != nil {
			return err
		}

		for _, v := range e.PFSeason {
			if err = manager.PFSeasonRepo.Add(ctx, &v); err != nil {
				return err
			}
		}

		for _, v := range e.PFranking {
			if err = manager.PFRankingRepo.Add(ctx, &v); err != nil {
				return err
			}
		}

		for _, v := range e.Emblem {
			if err = manager.EmblemRepo.AddOrUpdate(ctx, &v); err != nil {
				return err
			}
		}

		for _, v := range e.MotherBaseParam {
			if _, err = manager.MotherBaseParamRepo.AddOrUpdate(ctx, &v); err != nil {
				return err
			}
		}
	}

	return nil
}
