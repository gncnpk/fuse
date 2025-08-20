package handlers

import (
	"context"
	"fmt"
	"github.com/unknown321/fuse/seed"
	"github.com/unknown321/fuse/tppmessage"
	"log/slog"
	"net/url"
	"os"
)

func (gh *GateHandler) InitDB(ctx context.Context, baseURL string, platform string) error {
	var err error
	if err = gh.initUrlList(ctx, baseURL, platform); err != nil {
		return fmt.Errorf("urllist: %w", err)
	}

	if err = gh.initUser(ctx); err != nil {
		return fmt.Errorf("user: %w", err)
	}

	if err = gh.initPlayer(ctx); err != nil {
		return fmt.Errorf("init player: %w", err)
	}

	if err = gh.initTaskRewards(ctx); err != nil {
		return fmt.Errorf("init task rewards: %w", err)
	}

	if err = gh.initAbolition(ctx); err != nil {
		return fmt.Errorf("init abolition: %w", err)
	}

	if err = gh.initSession(ctx); err != nil {
		return fmt.Errorf("init session: %w", err)
	}

	if err = gh.manager.Init(ctx); err != nil {
		return fmt.Errorf("manager: %w", err)
	}

	if err = gh.initClusterCosts(ctx); err != nil {
		return fmt.Errorf("init cluster costs: %w", err)
	}

	if err = gh.initStaffRankBonusRates(ctx); err != nil {
		return fmt.Errorf("init staff rank bonus rates: %w", err)
	}

	if err = gh.initServerTexts(ctx); err != nil {
		return fmt.Errorf("init server texts: %w", err)
	}

	if err = gh.initEspionageEvents(ctx); err != nil {
		return fmt.Errorf("init espionage events: %w", err)
	}

	if err = gh.initPFEvents(ctx); err != nil {
		return fmt.Errorf("init pf events: %w", err)
	}

	if err = gh.initFOBEventRewards(ctx); err != nil {
		return fmt.Errorf("init fob event rewards: %w", err)
	}

	if err = gh.initFOBEventTimeBonuses(ctx); err != nil {
		return fmt.Errorf("init fob event time bonuses: %w", err)
	}

	if err = gh.initOnlineChallengeTasks(ctx); err != nil {
		return fmt.Errorf("init online challenge tasks: %w", err)
	}

	if err = gh.initOnlineChallengeTaskPlayer(ctx); err != nil {
		return fmt.Errorf("init online challenge tasks player: %w", err)
	}

	if err = gh.initServerProductParams(ctx); err != nil {
		return fmt.Errorf("init server product params: %w", err)
	}

	if err = gh.initServerProductParamPlayer(ctx); err != nil {
		return fmt.Errorf("init server product params player: %w", err)
	}

	if err = gh.initInformationMessages(ctx); err != nil {
		return fmt.Errorf("init information messages: %w", err)
	}

	if err = gh.initServerItem(ctx); err != nil {
		return fmt.Errorf("init server items: %w", err)
	}

	if err = gh.initEquipFlagRepo(ctx); err != nil {
		return fmt.Errorf("init equip flag: %w", err)
	}

	if err = gh.initEquipGradeRepo(ctx); err != nil {
		return fmt.Errorf("init equip grade: %w", err)
	}

	if err = gh.initTapeFlagRepo(ctx); err != nil {
		return fmt.Errorf("init tape flag: %w", err)
	}

	if err = gh.initSecurityLevelRepo(ctx); err != nil {
		return fmt.Errorf("init security level: %w", err)
	}

	if err = gh.initLocalBaseRepo(ctx); err != nil {
		return fmt.Errorf("init local base: %w", err)
	}

	if err = gh.initPFSkillStaffRepo(ctx); err != nil {
		return fmt.Errorf("init pf skill staff: %w", err)
	}

	if err = gh.initClusterSecurityRepo(ctx); err != nil {
		return fmt.Errorf("init cluster security: %w", err)
	}

	if err = gh.initClusterParamRepo(ctx); err != nil {
		return fmt.Errorf("init cluster param: %w", err)
	}

	if err = gh.initMotherBaseParamRepo(ctx); err != nil {
		return fmt.Errorf("init mother base param: %w", err)
	}

	if err = gh.initSectionStatRepo(ctx); err != nil {
		return fmt.Errorf("init section stat: %w", err)
	}

	if err = gh.initFobRecordRepo(ctx); err != nil {
		return fmt.Errorf("init fob record: %w", err)
	}

	if err = gh.initFobStatusRepo(ctx); err != nil {
		return fmt.Errorf("init fob status: %w", err)
	}

	if err = gh.initPlayerResourceRepo(ctx); err != nil {
		return fmt.Errorf("init player resource: %w", err)
	}

	if err = gh.initPlayerStatusRepo(ctx); err != nil {
		return fmt.Errorf("init player status: %w", err)
	}

	if err = gh.initFobEventRepo(ctx); err != nil {
		return fmt.Errorf("init fob events: %w", err)
	}

	if err = gh.initServerStatusRepo(ctx); err != nil {
		return fmt.Errorf("init server status: %w", err)
	}

	if err = gh.initPFSeasonRepo(ctx); err != nil {
		return fmt.Errorf("init pf season: %w", err)
	}

	if err = gh.initPFRankingRepo(ctx); err != nil {
		return fmt.Errorf("init pf ranking: %w", err)
	}

	if err = gh.initEmblemRepo(ctx); err != nil {
		return fmt.Errorf("init emblem: %w", err)
	}

	if err = gh.initFOBPlacedRepo(ctx); err != nil {
		return fmt.Errorf("init fob placed: %w", err)
	}

	if err = gh.initFOBWeaponPlacementRepo(ctx); err != nil {
		return fmt.Errorf("init fob weapon placement: %w", err)
	}

	if err = gh.initIntruderRepo(ctx); err != nil {
		return fmt.Errorf("init intruder: %w", err)
	}

	return nil
}

func (gh *GateHandler) initUser(ctx context.Context) error {
	gh.UserRepo.WithDB(gh.DB)
	err := gh.UserRepo.Init(ctx)
	if err != nil {
		return fmt.Errorf("user repo: %w", err)
	}

	return nil
}

func (gh *GateHandler) initPlayer(ctx context.Context) error {
	gh.PlayerRepo.WithDB(gh.DB)
	err := gh.PlayerRepo.Init(ctx)
	if err != nil {
		return fmt.Errorf("init player schema: %w", err)
	}

	return nil
}

func (gh *GateHandler) initTaskRewards(ctx context.Context) error {
	gh.TaskRewardRepo.WithDB(gh.DB)
	gh.PlayerTaskRepo.WithDB(gh.DB)

	if err := gh.TaskRewardRepo.CreateSchema(ctx); err != nil {
		return fmt.Errorf("init task reward schema: %w", err)
	}

	res, err := gh.TaskRewardRepo.Count(ctx)
	if err != nil {
		return err
	}

	if res != 0 {
		slog.Info("task rewards already seeded")
		return nil
	}

	if err = gh.TaskRewardRepo.Seed(ctx); err != nil {
		return fmt.Errorf("init task reward, seed: %w", err)
	}

	if err = gh.PlayerTaskRepo.InitDB(ctx); err != nil {
		return fmt.Errorf("init player task repo: %w", err)
	}

	return nil
}

func (gh *GateHandler) initAbolition(ctx context.Context) error {
	gh.AbolitionRepo.WithDB(gh.DB)
	if err := gh.AbolitionRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initSession(ctx context.Context) error {
	gh.SessionRepo.WithDB(gh.DB)
	if err := gh.SessionRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initClusterCosts(ctx context.Context) error {
	gh.ClusterBuildCostRepo.WithDB(gh.DB)
	if err := gh.ClusterBuildCostRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.ClusterBuildCostRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initStaffRankBonusRates(ctx context.Context) error {
	gh.StaffRankBonusRateRepo.WithDB(gh.DB)
	if err := gh.StaffRankBonusRateRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.StaffRankBonusRateRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initServerTexts(ctx context.Context) error {
	gh.ServerTextRepo.WithDB(gh.DB)
	if err := gh.ServerTextRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.ServerTextRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initEspionageEvents(ctx context.Context) error {
	gh.EspionageEventRepo.WithDB(gh.DB)
	if err := gh.EspionageEventRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.EspionageEventRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initPFEvents(ctx context.Context) error {
	gh.PFEventRepo.WithDB(gh.DB)
	if err := gh.PFEventRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.PFEventRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initFOBEventRewards(ctx context.Context) error {
	gh.FOBEventRewardRepo.WithDB(gh.DB)
	if err := gh.FOBEventRewardRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.FOBEventRewardRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initFOBEventTimeBonuses(ctx context.Context) error {
	gh.FOBEventTimeBonusRepo.WithDB(gh.DB)
	if err := gh.FOBEventTimeBonusRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.FOBEventTimeBonusRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initOnlineChallengeTasks(ctx context.Context) error {
	gh.OnlineChallengeTaskRepo.WithDB(gh.DB)
	if err := gh.OnlineChallengeTaskRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.OnlineChallengeTaskRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initOnlineChallengeTaskPlayer(ctx context.Context) error {
	gh.OnlineChallengeTaskPlayerRepo.WithDB(gh.DB)
	if err := gh.OnlineChallengeTaskPlayerRepo.Init(ctx); err != nil {
		return err
	}
	return nil
}

func (gh *GateHandler) initServerProductParams(ctx context.Context) error {
	gh.ServerProductParamRepo.WithDB(gh.DB)
	if err := gh.ServerProductParamRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.ServerProductParamRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initServerProductParamPlayer(ctx context.Context) error {
	gh.ServerProductParamPlayerRepo.WithDB(gh.DB)
	if err := gh.ServerProductParamPlayerRepo.Init(ctx); err != nil {
		return err
	}
	return nil
}

func (gh *GateHandler) initInformationMessages(ctx context.Context) error {
	gh.InformationMessageRepo.WithDB(gh.DB)
	if err := gh.InformationMessageRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.InformationMessageRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initServerItem(ctx context.Context) error {
	gh.ServerItemRepo.WithDB(gh.DB)
	if err := gh.ServerItemRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initEquipFlagRepo(ctx context.Context) error {
	gh.EquipFlagRepo.WithDB(gh.DB)
	if err := gh.EquipFlagRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initEquipGradeRepo(ctx context.Context) error {
	gh.EquipGradeRepo.WithDB(gh.DB)
	if err := gh.EquipGradeRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initTapeFlagRepo(ctx context.Context) error {
	gh.TapeFlagRepo.WithDB(gh.DB)
	if err := gh.TapeFlagRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initSecurityLevelRepo(ctx context.Context) error {
	gh.SecurityLevelRepo.WithDB(gh.DB)
	if err := gh.SecurityLevelRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initLocalBaseRepo(ctx context.Context) error {
	gh.LocalBaseRepo.WithDB(gh.DB)
	if err := gh.LocalBaseRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initPFSkillStaffRepo(ctx context.Context) error {
	gh.PFSkillStaffRepo.WithDB(gh.DB)
	if err := gh.PFSkillStaffRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initClusterSecurityRepo(ctx context.Context) error {
	gh.ClusterSecurityRepo.WithDB(gh.DB)
	if err := gh.ClusterSecurityRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initClusterParamRepo(ctx context.Context) error {
	gh.ClusterParamRepo.WithDB(gh.DB)
	if err := gh.ClusterParamRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initMotherBaseParamRepo(ctx context.Context) error {
	gh.MotherBaseParamRepo.WithDB(gh.DB)
	if err := gh.MotherBaseParamRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initSectionStatRepo(ctx context.Context) error {
	gh.SectionStatRepo.WithDB(gh.DB)
	if err := gh.SectionStatRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initFobRecordRepo(ctx context.Context) error {
	gh.FobRecordRepo.WithDB(gh.DB)
	if err := gh.FobRecordRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initFobStatusRepo(ctx context.Context) error {
	gh.FobStatusRepo.WithDB(gh.DB)
	if err := gh.FobStatusRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initPlayerResourceRepo(ctx context.Context) error {
	gh.PlayerResourceRepo.WithDB(gh.DB)
	if err := gh.PlayerResourceRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initPlayerStatusRepo(ctx context.Context) error {
	gh.PlayerStatusRepo.WithDB(gh.DB)
	if err := gh.PlayerStatusRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initFobEventRepo(ctx context.Context) error {
	gh.FobEventRepo.WithDB(gh.DB)
	if err := gh.FobEventRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.FobEventRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initServerStatusRepo(ctx context.Context) error {
	gh.ServerStatusRepo.WithDB(gh.DB)
	if err := gh.ServerStatusRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.ServerStatusRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initPFRankingRepo(ctx context.Context) error {
	gh.PFRankingRepo.WithDB(gh.DB)
	if err := gh.PFRankingRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initPFSeasonRepo(ctx context.Context) error {
	gh.PFSeasonRepo.WithDB(gh.DB)
	if err := gh.PFSeasonRepo.Init(ctx); err != nil {
		return err
	}

	if err := gh.PFSeasonRepo.Seed(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initEmblemRepo(ctx context.Context) error {
	gh.EmblemRepo.WithDB(gh.DB)
	if err := gh.EmblemRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initFOBPlacedRepo(ctx context.Context) error {
	gh.FOBPlacedRepo.WithDB(gh.DB)
	if err := gh.FOBPlacedRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initFOBWeaponPlacementRepo(ctx context.Context) error {
	gh.FOBWeaponPlacementRepo.WithDB(gh.DB)
	if err := gh.FOBWeaponPlacementRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initIntruderRepo(ctx context.Context) error {
	gh.IntruderRepo.WithDB(gh.DB)
	if err := gh.IntruderRepo.Init(ctx); err != nil {
		return err
	}

	return nil
}

func (gh *GateHandler) initGDPRList(ctx context.Context) error {
	slog.Info("TODO implement me")
	return nil
}

func (gh *GateHandler) initUrlList(ctx context.Context, baseURL string, platform string) error {
	gh.URLListEntryRepo.WithDB(gh.DB)
	err := gh.URLListEntryRepo.CreateSchema()
	if err != nil {
		return fmt.Errorf("urllist init: %w", err)
	}
	err = gh.URLListEntryRepo.Clear(ctx)
	if err != nil {
		return fmt.Errorf("clear: %w", err)
	}

	b, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("invalid base url %s: %w", baseURL, err)
	}

	var platformWeb = platform + "web"

	// TODO to database or at least seed
	l := []tppmessage.URLListEntry{
		{
			Type:    "GATE",
			Url:     b.JoinPath(platform, "gate").String(), // tppstm/gate
			Version: 17,
		},
		{
			Type:    "WEB",
			Url:     b.JoinPath(platform, "main").String(), // tppstm/main
			Version: 17,
		},
		{
			Type:    "EULA",
			Url:     b.JoinPath(platformWeb, "eula", "eula.var").String(), // "/tppstmweb/eula/eula.var",
			Version: 6,
		},
		{
			Type:    "HEATMAP",
			Url:     "http://mgstpp-app.konamionline.com/tppstmweb/heatmap",
			Version: 0,
		},
		{
			Type:    "DEVICE",
			Url:     "http://mgstpp-app.konamionline.com/tppstm/main",
			Version: 0,
		},
		{
			Type:    "EULA_COIN",
			Url:     b.JoinPath(platformWeb, "coin", "coin.var").String(), // tppstmweb/coin/coin.var
			Version: 1,
		},
		{
			Type:    "POLICY_GDPR",
			Url:     b.JoinPath(platformWeb, "gdpr", "privacy.var").String(), // tppstmweb/gdpr/privacy.var"
			Version: 1,
		},
		{
			Type:    "POLICY_JP",
			Url:     b.JoinPath(platformWeb, "privacy_jp", "privacy.var").String(), // tppstmweb/privacy_jp/privacy.var",
			Version: 2,
		},
		{
			Type:    "POLICY_ELSE",
			Url:     b.JoinPath(platformWeb, "privacy", "privacy.var").String(), // tppstmweb/privacy/privacy.var",
			Version: 1,
		},
		{
			Type:    "LEGAL",
			Url:     "http://legal.konami.com/games/mgsvtpp/",
			Version: 1,
		},
		{
			Type:    "PERMISSION",
			Url:     "http://www.konami.com/",
			Version: 1,
		},
		{
			Type:    "POLICY_CCPA",
			Url:     b.JoinPath(platformWeb, "privacy_ccpa", "privacy.var").String(), // tppstmweb/privacy_ccpa/privacy.var",
			Version: 1,
		},
		{
			Type:    "EULA_TEXT",
			Url:     "http://legal.konami.com/games/mgsvtpp/terms/",
			Version: 1,
		},
		{
			Type:    "EULA_COIN_TEXT",
			Url:     "http://legal.konami.com/games/mgsvtpp/terms/currency/",
			Version: 1,
		},
		{
			Type:    "POLICY_GDPR_TEXT",
			Url:     "http://legal.konami.com/games/mgsvtpp/",
			Version: 1,
		},
		{
			Type:    "POLICY_JP_TEXT",
			Url:     "http://legal.konami.com/games/privacy/view/",
			Version: 2,
		},
		{
			Type:    "POLICY_ELSE_TEXT",
			Url:     "http://legal.konami.com/games/privacy/view/",
			Version: 1,
		},
		{
			Type:    "POLICY_CCPA_TEXT",
			Url:     "http://legal.konami.com/games/mgsvtpp/ppa4ca/",
			Version: 1,
		},
	}

	for _, e := range l {
		err = gh.URLListEntryRepo.Insert(ctx, e)
		if err != nil {
			return fmt.Errorf("cannot insert url list entry %s: %w", e.Type, err)
		}
	}

	return nil
}

func (gh *GateHandler) Seed(ctx context.Context) error {
	data, err := os.ReadFile("./seed.json")
	if err != nil {
		return err
	}

	s, err := seed.Read(data)
	if err != nil {
		return err
	}

	slog.Info("seeding")
	if err = s.Seed(gh.manager); err != nil {
		return fmt.Errorf("seed fail: %w", err)
	}
	slog.Info("seeding done")

	return nil
}
