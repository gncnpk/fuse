package handlers

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/unknown321/fuse/abolition"
	"github.com/unknown321/fuse/challengetask"
	"github.com/unknown321/fuse/clusterbuildcost"
	"github.com/unknown321/fuse/clusterparam"
	"github.com/unknown321/fuse/clustersecurity"
	"github.com/unknown321/fuse/coder"
	"github.com/unknown321/fuse/emblem"
	"github.com/unknown321/fuse/equipflag"
	"github.com/unknown321/fuse/equipgrade"
	fobevent "github.com/unknown321/fuse/fobevent/event"
	fobeventreward "github.com/unknown321/fuse/fobevent/reward"
	fobeventtimebonus "github.com/unknown321/fuse/fobevent/timebonus"
	"github.com/unknown321/fuse/fobplaced"
	"github.com/unknown321/fuse/fobrecord"
	"github.com/unknown321/fuse/fobstatus"
	"github.com/unknown321/fuse/fobweaponplacement"
	"github.com/unknown321/fuse/informationmessage"
	"github.com/unknown321/fuse/intruder"
	"github.com/unknown321/fuse/localbase"
	"github.com/unknown321/fuse/message"
	"github.com/unknown321/fuse/motherbaseparam"
	"github.com/unknown321/fuse/onlinechallengetask"
	onlinechallengetaskplayer "github.com/unknown321/fuse/onlinechallengetask/player"
	"github.com/unknown321/fuse/pfranking"
	"github.com/unknown321/fuse/pfseason"
	"github.com/unknown321/fuse/pfskillstaff"
	"github.com/unknown321/fuse/player"
	"github.com/unknown321/fuse/playerresource"
	"github.com/unknown321/fuse/playerstatus"
	"github.com/unknown321/fuse/playertask"
	"github.com/unknown321/fuse/ranking/espionageevent"
	"github.com/unknown321/fuse/ranking/pfevent"
	"github.com/unknown321/fuse/sectionstat"
	"github.com/unknown321/fuse/securitylevel"
	"github.com/unknown321/fuse/serveritem"
	"github.com/unknown321/fuse/serverproductparam"
	serverproductparamplayer "github.com/unknown321/fuse/serverproductparam/player"
	"github.com/unknown321/fuse/serverstatus"
	"github.com/unknown321/fuse/servertext"
	"github.com/unknown321/fuse/session"
	"github.com/unknown321/fuse/sessionmanager"
	"github.com/unknown321/fuse/staffrankrate"
	"github.com/unknown321/fuse/tapeflag"
	"github.com/unknown321/fuse/tppmessage"
	"github.com/unknown321/fuse/user"
	"github.com/unknown321/fuse/util"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"

	_ "modernc.org/sqlite"
)

type GateHandler struct {
	manager                       *sessionmanager.SessionManager
	coder                         *coder.Coder
	PassThrough                   bool
	FromJSON                      bool
	DB                            *sql.DB
	URLListEntryRepo              tppmessage.URLListEntryRepo
	UserRepo                      user.Repo
	PlayerRepo                    player.Repo
	AbolitionRepo                 abolition.Repo
	TaskRewardRepo                challengetask.TaskRewardRepo
	PlayerTaskRepo                playertask.Repo
	SessionRepo                   session.Repo
	ClusterBuildCostRepo          clusterbuildcost.Repo
	StaffRankBonusRateRepo        staffrankrate.Repo
	ServerTextRepo                servertext.Repo
	EspionageEventRepo            espionageevent.Repo
	PFEventRepo                   pfevent.Repo
	FOBEventRewardRepo            fobeventreward.Repo
	FOBEventTimeBonusRepo         fobeventtimebonus.Repo
	OnlineChallengeTaskRepo       onlinechallengetask.Repo
	OnlineChallengeTaskPlayerRepo onlinechallengetaskplayer.Repo
	ServerProductParamRepo        serverproductparam.Repo
	ServerProductParamPlayerRepo  serverproductparamplayer.Repo
	InformationMessageRepo        informationmessage.Repo
	ServerItemRepo                serveritem.Repo
	EquipFlagRepo                 equipflag.Repo
	EquipGradeRepo                equipgrade.Repo
	TapeFlagRepo                  tapeflag.Repo
	SecurityLevelRepo             securitylevel.Repo
	LocalBaseRepo                 localbase.Repo
	PFSkillStaffRepo              pfskillstaff.Repo
	ClusterSecurityRepo           clustersecurity.Repo
	ClusterParamRepo              clusterparam.Repo
	MotherBaseParamRepo           motherbaseparam.Repo
	SectionStatRepo               sectionstat.Repo
	FobRecordRepo                 fobrecord.Repo
	FobStatusRepo                 fobstatus.Repo
	PlayerResourceRepo            playerresource.Repo
	PlayerStatusRepo              playerstatus.Repo
	FobEventRepo                  fobevent.Repo
	ServerStatusRepo              serverstatus.Repo
	PFRankingRepo                 pfranking.Repo
	PFSeasonRepo                  pfseason.Repo
	EmblemRepo                    emblem.Repo
	FOBPlacedRepo                 fobplaced.Repo
	FOBWeaponPlacementRepo        fobweaponplacement.Repo
	IntruderRepo                  intruder.Repo
}

func (gh *GateHandler) WithManager(m *sessionmanager.SessionManager) {
	gh.manager = m
	m.URLListEntryRepo = &gh.URLListEntryRepo
	m.UserRepo = &gh.UserRepo
	m.PlayerRepo = &gh.PlayerRepo
	m.TaskRewardRepo = &gh.TaskRewardRepo
	m.PlayerTaskRepo = &gh.PlayerTaskRepo
	m.AbolitionRepo = &gh.AbolitionRepo
	m.SessionRepo = &gh.SessionRepo
	m.ClusterBuildCostRepo = &gh.ClusterBuildCostRepo
	m.StaffRankBonusRateRepo = &gh.StaffRankBonusRateRepo
	m.ServerTextRepo = &gh.ServerTextRepo
	m.EspionageEventRepo = &gh.EspionageEventRepo
	m.PFEventRepo = &gh.PFEventRepo
	m.FOBEventRewardRepo = &gh.FOBEventRewardRepo
	m.FOBEventTimeBonusRepo = &gh.FOBEventTimeBonusRepo
	m.OnlineChallengeTaskRepo = &gh.OnlineChallengeTaskRepo
	m.OnlineChallengeTaskPlayerRepo = &gh.OnlineChallengeTaskPlayerRepo
	m.ServerProductParamRepo = &gh.ServerProductParamRepo
	m.ServerProductParamPlayerRepo = &gh.ServerProductParamPlayerRepo
	m.InformationMessageRepo = &gh.InformationMessageRepo
	m.ServerItemRepo = &gh.ServerItemRepo
	m.EquipFlagRepo = &gh.EquipFlagRepo
	m.EquipGradeRepo = &gh.EquipGradeRepo
	m.TapeFlagRepo = &gh.TapeFlagRepo
	m.SecurityLevelRepo = &gh.SecurityLevelRepo
	m.LocalBaseRepo = &gh.LocalBaseRepo
	m.PFSkillStaffRepo = &gh.PFSkillStaffRepo
	m.ClusterSecurityRepo = &gh.ClusterSecurityRepo
	m.ClusterParamRepo = &gh.ClusterParamRepo
	m.MotherBaseParamRepo = &gh.MotherBaseParamRepo
	m.SectionStatRepo = &gh.SectionStatRepo
	m.FobRecordRepo = &gh.FobRecordRepo
	m.FobStatusRepo = &gh.FobStatusRepo
	m.PlayerResourceRepo = &gh.PlayerResourceRepo
	m.PlayerStatusRepo = &gh.PlayerStatusRepo
	m.FobEventRepo = &gh.FobEventRepo
	m.ServerStatusRepo = &gh.ServerStatusRepo
	m.PFRankingRepo = &gh.PFRankingRepo
	m.PFSeasonRepo = &gh.PFSeasonRepo
	m.EmblemRepo = &gh.EmblemRepo
	m.FOBPlacedRepo = &gh.FOBPlacedRepo
	m.FOBWeaponPlacementRepo = &gh.FOBWeaponPlacementRepo
	m.IntruderRepo = &gh.IntruderRepo
}

func (gh *GateHandler) WithCoder(c *coder.Coder) {
	gh.coder = c
}

func (gh *GateHandler) DBConnect(dsnURI string) error {
	var err error
	gh.DB, err = sql.Open("sqlite", dsnURI)
	if err != nil {
		return fmt.Errorf("cannot open database %s: %w", dsnURI, err)
	}

	err = gh.DB.Ping()
	if err != nil {
		return fmt.Errorf("cannot ping: %w", err)
	}

	_, err = gh.DB.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		return fmt.Errorf("failed to enable WAL mode: %w", err)
	}

	_, err = gh.DB.Exec(`
		PRAGMA synchronous=NORMAL;
		PRAGMA wal_autocheckpoint=1000;
	`)
	if err != nil {
		return fmt.Errorf("failed to configure WAL settings: %w", err)
	}

	return nil
}

func (gh *GateHandler) DecodeIn(request *http.Request, m *message.Message, username *string) ([]byte, error) {
	b, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot read request body: %w", err)
	}

	q := bytes.TrimPrefix(b, []byte("httpMsg="))
	data, err := url.QueryUnescape(string(q))
	if err != nil {
		return nil, fmt.Errorf("cannot query unescape: %w", err)
	}

	m.IsRequest = true
	err = m.Decode([]byte(data))
	if err != nil {
		return nil, fmt.Errorf("cannot decode request %w", err)
	}

	return b, nil
}

func (gh *GateHandler) DecodeOut(res []byte, m *message.Message, username *string) error {
	m.IsRequest = false
	err := m.Decode(res)
	if err != nil {
		return fmt.Errorf("gate decode out err: %w", err)
	}

	return nil
}

func (gh *GateHandler) Handle(writer http.ResponseWriter, request *http.Request) {
	m := message.Message{}
	m.WithCoder(gh.coder)

	var userName string
	b, err := gh.DecodeIn(request, &m, &userName)
	if err != nil {
		slog.Error("cannot decode gate message", "error", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	slog.Info("client request", "passthrough", gh.PassThrough, "ip", request.RemoteAddr)

	ctx := request.Context()
	if gh.FromJSON {
		ctx = context.WithValue(ctx, "fromjson", 1)
	}

	err = gh.manager.Handle(ctx, &m)
	if err != nil {
		slog.Error("cannot handle incoming message", "error", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var res []byte
	code := http.StatusOK

	if gh.PassThrough {
		resp, err := ToKojiPro(request.URL.Path, bytes.NewReader(b), request.ContentLength)
		if err != nil {
			slog.Error("kojiPro fail", "error", err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		code = resp.StatusCode

		res, err = io.ReadAll(resp.Body)
		if err != nil {
			slog.Error("cannot read kojiPro resp body", "error", err.Error())
			return
		}

		err = gh.DecodeOut(res, &m, &userName)
		if err != nil {
			slog.Error("cannot decode response", "error", err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		slog.Info("kojipro response")
		err = gh.manager.Handle(request.Context(), &m)
		if err != nil {
			slog.Error("cannot handle outgoing message", "error", err.Error())
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	encode, err := m.Encode()
	if err != nil {
		slog.Error("cannot encode outgoing message", "error", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	lines := util.SplitByteString(encode, 76)
	slog.Debug("adding newlines", "lines", len(lines), "total", len(lines)*2)
	encode = bytes.Join(lines, []byte("\r\n"))
	encode = append(encode, []byte("\r\n")...)

	writer.Header().Set("Content-Length", strconv.Itoa(len(encode)))
	writer.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	slog.Debug("responding with", "content-length", strconv.Itoa(len(encode)))
	writer.WriteHeader(code)
	_, _ = writer.Write(encode)
	slog.Info("client request handled", "type", m.MsgID.String(), "ip", request.RemoteAddr)
}
