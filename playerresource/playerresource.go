package playerresource

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type Processed struct {
	Fuel          int
	Bio           int
	CommonMetal   int
	MinorMetal    int
	PreciousMetal int
}

type Raw struct {
	Fuel          int
	Bio           int
	CommonMetal   int
	MinorMetal    int
	PreciousMetal int
}

type Plants struct {
	Wormwood          int
	BlackCarrot       int
	GoldenCrescent    int
	Tarragon          int
	AfricanPeach      int
	DigitalisPurpurea int
	DigitalisLutea    int
	Haoma             int
}

type Placed struct {
	VolgaK12     int
	HMG3Wingate  int
	M2A304Mortar int
	Zhizdra45    int
	M276AAGGun   int
}

type Vehicles struct {
	ZaAZS84       int
	APET41LV      int
	ZiGRA6T       int
	Boar53CT      int
	ZhukBr3       int
	StoutIfvSc    int
	ZhukRsZo      int
	StoutIfvFs    int
	TT77Nosorog   int
	M84AMagloader int
}

type WalkerGears struct {
	WGPP        int
	CCCPWGTypeC int
	CCCPWGTypeA int
	CFAWGTypeC  int
	CFAWGTypeA  int
}

type Parasites struct {
	Mist       int
	Camouflage int
	Armor      int
}

type Nuclear struct {
	Weapon int
	Waste  int
}

type PlayerResource struct {
	PlayerID    int
	IsOnline    bool
	Raw         Raw         // 5, used in diff2 only
	Processed   Processed   // 5
	Plants      Plants      // 8
	Vehicles    Vehicles    // 10
	WalkerGears WalkerGears // 5
	Nuclear     Nuclear     // 2
	Parasites   Parasites   // 3
	Mystery01   int         // always 0?
	Placed      Placed      // 5
	Mystery02   [7]int
	Unused      [13]int
}

func (p *PlayerResource) FromArray(diff1 []int, diff2 []int) error {
	if len(diff1) != 59 {
		return fmt.Errorf("invalid input length: %d, want 59", len(diff1))
	}

	p.Raw.Fuel = diff2[0]
	p.Raw.Bio = diff2[1]
	p.Raw.CommonMetal = diff2[2]
	p.Raw.MinorMetal = diff2[3]
	p.Raw.PreciousMetal = diff2[4]

	for i, v := range diff2[5:] {
		if v != 0 {
			slog.Info("raw resource diff is not 0", "index", i, "value", v)
		}
	}

	p.Processed.Fuel = diff1[0]
	p.Processed.Bio = diff1[1]
	p.Processed.CommonMetal = diff1[2]
	p.Processed.MinorMetal = diff1[3]
	p.Processed.PreciousMetal = diff1[4]

	p.Plants.Wormwood = diff1[5]
	p.Plants.BlackCarrot = diff1[6]
	p.Plants.GoldenCrescent = diff1[7]
	p.Plants.Tarragon = diff1[8]
	p.Plants.AfricanPeach = diff1[9]
	p.Plants.DigitalisPurpurea = diff1[10]
	p.Plants.DigitalisLutea = diff1[11]
	p.Plants.Haoma = diff1[12]

	p.Vehicles.ZaAZS84 = diff1[13]
	p.Vehicles.APET41LV = diff1[14]
	p.Vehicles.ZiGRA6T = diff1[15]
	p.Vehicles.Boar53CT = diff1[16]
	p.Vehicles.ZhukBr3 = diff1[17]
	p.Vehicles.StoutIfvSc = diff1[18]
	p.Vehicles.ZhukRsZo = diff1[19]
	p.Vehicles.StoutIfvFs = diff1[20]
	p.Vehicles.TT77Nosorog = diff1[21]
	p.Vehicles.M84AMagloader = diff1[22]

	p.WalkerGears.WGPP = diff1[23]
	p.WalkerGears.CCCPWGTypeC = diff1[24]
	p.WalkerGears.CCCPWGTypeA = diff1[25]
	p.WalkerGears.CFAWGTypeC = diff1[26]
	p.WalkerGears.CFAWGTypeA = diff1[27]

	p.Nuclear.Weapon = diff1[28]
	p.Nuclear.Waste = diff1[29]

	p.Parasites.Mist = diff1[30]
	p.Parasites.Camouflage = diff1[31]
	p.Parasites.Armor = diff1[32]

	p.Mystery01 = diff1[33]
	if p.Mystery01 != 0 {
		slog.Info("mystery01 resource is not 0", "value", p.Mystery01)
	}

	p.Placed.VolgaK12 = diff1[34]
	p.Placed.HMG3Wingate = diff1[35]
	p.Placed.M2A304Mortar = diff1[36]
	p.Placed.Zhizdra45 = diff1[37]
	p.Placed.M276AAGGun = diff1[38]

	p.Mystery02[0] = diff1[39]
	p.Mystery02[1] = diff1[40]
	p.Mystery02[2] = diff1[41]
	p.Mystery02[3] = diff1[42]
	p.Mystery02[4] = diff1[43]
	p.Mystery02[5] = diff1[44]
	p.Mystery02[6] = diff1[45]

	for i, v := range p.Mystery02 {
		if v != 0 {
			slog.Info("mystery02 resource", "value", v, "index", i)
		}
	}

	for i, v := range diff1[46:] {
		p.Unused[i] = v
		if v != 0 {
			slog.Info("unused is not 0", "index", i, "value", v)
		}
	}

	return nil
}

func (p *PlayerResource) ToArray(isRaw bool) []int {
	out := make([]int, 59)

	if isRaw {
		out[0] = p.Raw.Fuel
		out[1] = p.Raw.Bio
		out[2] = p.Raw.CommonMetal
		out[3] = p.Raw.MinorMetal
		out[4] = p.Raw.PreciousMetal

		return out
	}

	out[0] = p.Raw.Fuel
	out[1] = p.Raw.Bio
	out[2] = p.Raw.CommonMetal
	out[3] = p.Raw.MinorMetal
	out[4] = p.Raw.PreciousMetal

	out[5] = p.Plants.Wormwood
	out[6] = p.Plants.BlackCarrot
	out[7] = p.Plants.GoldenCrescent
	out[8] = p.Plants.Tarragon
	out[9] = p.Plants.AfricanPeach
	out[10] = p.Plants.DigitalisPurpurea
	out[11] = p.Plants.DigitalisLutea
	out[12] = p.Plants.Haoma

	out[13] = p.Vehicles.ZaAZS84
	out[14] = p.Vehicles.APET41LV
	out[15] = p.Vehicles.ZiGRA6T
	out[16] = p.Vehicles.Boar53CT
	out[17] = p.Vehicles.ZhukBr3
	out[18] = p.Vehicles.StoutIfvSc
	out[19] = p.Vehicles.ZhukRsZo
	out[20] = p.Vehicles.StoutIfvFs
	out[21] = p.Vehicles.TT77Nosorog
	out[22] = p.Vehicles.M84AMagloader

	out[23] = p.WalkerGears.WGPP
	out[24] = p.WalkerGears.CCCPWGTypeC
	out[25] = p.WalkerGears.CCCPWGTypeA
	out[26] = p.WalkerGears.CFAWGTypeC
	out[27] = p.WalkerGears.CFAWGTypeA

	out[28] = p.Nuclear.Weapon
	out[29] = p.Nuclear.Waste

	out[30] = p.Parasites.Mist
	out[31] = p.Parasites.Camouflage
	out[32] = p.Parasites.Armor

	out[33] = p.Mystery01

	out[34] = p.Placed.VolgaK12
	out[35] = p.Placed.HMG3Wingate
	out[36] = p.Placed.M2A304Mortar
	out[37] = p.Placed.Zhizdra45
	out[38] = p.Placed.M276AAGGun

	out[39] = p.Mystery02[0]
	out[40] = p.Mystery02[1]
	out[41] = p.Mystery02[2]
	out[42] = p.Mystery02[3]
	out[43] = p.Mystery02[4]
	out[44] = p.Mystery02[5]
	out[45] = p.Mystery02[6]

	for i, v := range p.Unused {
		out[46+i] = v
	}

	return out
}

var TableName = "playerResource"

type Repo struct {
	db *sql.DB
}

func (r *Repo) WithDB(db *sql.DB) {
	r.db = db
}

func (r *Repo) Init(ctx context.Context) error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			player_id INTEGER,
			is_online INTEGER,

			raw_fuel INTEGER,
			raw_bio INTEGER,
			raw_common_metal INTEGER,
			raw_minor_metal INTEGER,
			raw_precious_metal INTEGER,

			processed_fuel INTEGER,
			processed_bio INTEGER,
			processed_common_metal INTEGER,
			processed_minor_metal INTEGER,
			processed_precious_metal INTEGER,

			plants_wormwood INTEGER,
			plants_black_carrot INTEGER,
			plants_golden_crescent INTEGER,
			plants_tarragon INTEGER,
			plants_african_peach INTEGER,
			plants_digitalis_purpurea INTEGER,
			plants_digitalis_lutea INTEGER,
			plants_haoma INTEGER,

			vehicles_zaazs84 INTEGER,
			vehicles_apet41lv INTEGER,
			vehicles_zigra6t INTEGER,
			vehicles_boar53ct INTEGER,
			vehicles_zhuk_br3 INTEGER,
			vehicles_stout_ifv_sc INTEGER,
			vehicles_zhuk_rs_zo INTEGER,
			vehicles_stout_ifv_fs INTEGER,
			vehicles_tt77_nosorog INTEGER,
			vehicles_m84a_magloader INTEGER,

			walker_gears_wgpp INTEGER,
			walker_gears_cccp_wg_type_c INTEGER,
			walker_gears_cccp_wg_type_a INTEGER,
			walker_gears_cfa_wg_type_c INTEGER,
			walker_gears_cfa_wg_type_a INTEGER,

			nuclear_weapon INTEGER,
			nuclear_waste INTEGER,

			parasites_mist INTEGER,
			parasites_camouflage INTEGER,
			parasites_armor INTEGER,

			mystery01 INTEGER,

			placed_volga_k12 INTEGER,
			placed_hmg3_wingate INTEGER,
			placed_m2a304_mortar INTEGER,
			placed_zhizdra45 INTEGER,
			placed_m276_aa_gun INTEGER,

			mystery02_0 INTEGER,
			mystery02_1 INTEGER,
			mystery02_2 INTEGER,
			mystery02_3 INTEGER,
			mystery02_4 INTEGER,
			mystery02_5 INTEGER,
			mystery02_6 INTEGER,

			unused_0 INTEGER,
			unused_1 INTEGER,
			unused_2 INTEGER,
			unused_3 INTEGER,
			unused_4 INTEGER,
			unused_5 INTEGER,
			unused_6 INTEGER,
			unused_7 INTEGER,
			unused_8 INTEGER,
			unused_9 INTEGER,
			unused_10 INTEGER,
			unused_11 INTEGER,
			unused_12 INTEGER,
			UNIQUE(player_id, is_online)
		);`,
		TableName,
	)

	_, err := r.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("cannot create schema: %w", err)
	}
	return nil
}

func (r *Repo) AddOrUpdate(ctx context.Context, c *PlayerResource) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	q := fmt.Sprintf(`INSERT OR REPLACE INTO %s(
			player_id,
			is_online,

			raw_fuel,
			raw_bio,
			raw_common_metal,
			raw_minor_metal,
			raw_precious_metal,

			processed_fuel,
			processed_bio,
			processed_common_metal,
			processed_minor_metal,
			processed_precious_metal,

			plants_wormwood,
			plants_black_carrot,
			plants_golden_crescent,
			plants_tarragon,
			plants_african_peach,
			plants_digitalis_purpurea,
			plants_digitalis_lutea,
			plants_haoma,

			vehicles_zaazs84,
			vehicles_apet41lv,
			vehicles_zigra6t,
			vehicles_boar53ct,
			vehicles_zhuk_br3,
			vehicles_stout_ifv_sc,
			vehicles_zhuk_rs_zo,
			vehicles_stout_ifv_fs,
			vehicles_tt77_nosorog,
			vehicles_m84a_magloader,

			walker_gears_wgpp,
			walker_gears_cccp_wg_type_c,
			walker_gears_cccp_wg_type_a,
			walker_gears_cfa_wg_type_c,
			walker_gears_cfa_wg_type_a,

			nuclear_weapon,
			nuclear_waste,

			parasites_mist,
			parasites_camouflage,
			parasites_armor,

			mystery01,

			placed_volga_k12,
			placed_hmg3_wingate,
			placed_m2a304_mortar,
			placed_zhizdra45,
			placed_m276_aa_gun,

			mystery02_0,
			mystery02_1,
			mystery02_2,
			mystery02_3,
			mystery02_4,
			mystery02_5,
			mystery02_6,

			unused_0,
			unused_1,
			unused_2,
			unused_3,
			unused_4,
			unused_5,
			unused_6,
			unused_7,
			unused_8,
			unused_9,
			unused_10,
			unused_11,
			unused_12
		) values (?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?);`,
		TableName)
	if _, err = tx.ExecContext(ctx, q,
		c.PlayerID,
		c.IsOnline,

		c.Raw.Fuel,
		c.Raw.Bio,
		c.Raw.CommonMetal,
		c.Raw.MinorMetal,
		c.Raw.PreciousMetal,

		c.Processed.Fuel,
		c.Processed.Bio,
		c.Processed.CommonMetal,
		c.Processed.MinorMetal,
		c.Processed.PreciousMetal,

		c.Plants.Wormwood,
		c.Plants.BlackCarrot,
		c.Plants.GoldenCrescent,
		c.Plants.Tarragon,
		c.Plants.AfricanPeach,
		c.Plants.DigitalisPurpurea,
		c.Plants.DigitalisLutea,
		c.Plants.Haoma,

		c.Vehicles.ZaAZS84,
		c.Vehicles.APET41LV,
		c.Vehicles.ZiGRA6T,
		c.Vehicles.Boar53CT,
		c.Vehicles.ZhukBr3,
		c.Vehicles.StoutIfvSc,
		c.Vehicles.ZhukRsZo,
		c.Vehicles.StoutIfvFs,
		c.Vehicles.TT77Nosorog,
		c.Vehicles.M84AMagloader,

		c.WalkerGears.WGPP,
		c.WalkerGears.CCCPWGTypeC,
		c.WalkerGears.CCCPWGTypeA,
		c.WalkerGears.CFAWGTypeC,
		c.WalkerGears.CFAWGTypeA,

		c.Nuclear.Weapon,
		c.Nuclear.Waste,

		c.Parasites.Mist,
		c.Parasites.Camouflage,
		c.Parasites.Armor,

		c.Mystery01,

		c.Placed.VolgaK12,
		c.Placed.HMG3Wingate,
		c.Placed.M2A304Mortar,
		c.Placed.Zhizdra45,
		c.Placed.M276AAGGun,

		c.Mystery02[0],
		c.Mystery02[1],
		c.Mystery02[2],
		c.Mystery02[3],
		c.Mystery02[4],
		c.Mystery02[5],
		c.Mystery02[6],

		c.Unused[0],
		c.Unused[1],
		c.Unused[2],
		c.Unused[3],
		c.Unused[4],
		c.Unused[5],
		c.Unused[6],
		c.Unused[7],
		c.Unused[8],
		c.Unused[9],
		c.Unused[10],
		c.Unused[11],
		c.Unused[12],
	); err != nil {
		slog.Error("add fail", "error", err.Error(), "table", TableName, "playerID", c.PlayerID)
		if err = tx.Rollback(); err != nil {
			return fmt.Errorf("insert rollback failed: %w", err)
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) Get(ctx context.Context, playerID int, isOnline bool) ([]PlayerResource, error) {
	q := fmt.Sprintf(`SELECT 
			player_id,
			is_online,

			raw_fuel,
			raw_bio,
			raw_common_metal,
			raw_minor_metal,
			raw_precious_metal,

			processed_fuel,
			processed_bio,
			processed_common_metal,
			processed_minor_metal,
			processed_precious_metal,

			plants_wormwood,
			plants_black_carrot,
			plants_golden_crescent,
			plants_tarragon,
			plants_african_peach,
			plants_digitalis_purpurea,
			plants_digitalis_lutea,
			plants_haoma,

			vehicles_zaazs84,
			vehicles_apet41lv,
			vehicles_zigra6t,
			vehicles_boar53ct,
			vehicles_zhuk_br3,
			vehicles_stout_ifv_sc,
			vehicles_zhuk_rs_zo,
			vehicles_stout_ifv_fs,
			vehicles_tt77_nosorog,
			vehicles_m84a_magloader,

			walker_gears_wgpp,
			walker_gears_cccp_wg_type_c,
			walker_gears_cccp_wg_type_a,
			walker_gears_cfa_wg_type_c,
			walker_gears_cfa_wg_type_a,

			nuclear_weapon,
			nuclear_waste,

			parasites_mist,
			parasites_camouflage,
			parasites_armor,

			mystery01,

			placed_volga_k12,
			placed_hmg3_wingate,
			placed_m2a304_mortar,
			placed_zhizdra45,
			placed_m276_aa_gun,

			mystery02_0,
			mystery02_1,
			mystery02_2,
			mystery02_3,
			mystery02_4,
			mystery02_5,
			mystery02_6,

			unused_0,
			unused_1,
			unused_2,
			unused_3,
			unused_4,
			unused_5,
			unused_6,
			unused_7,
			unused_8,
			unused_9,
			unused_10,
			unused_11,
			unused_12
		FROM %s
		WHERE player_id = ? and is_online = ?;`,
		TableName)
	rows, err := r.db.QueryContext(ctx, q, playerID, isOnline)
	if err != nil {
		return nil, err
	}

	var res []PlayerResource
	for rows.Next() {
		c := PlayerResource{}
		rows.Scan(
			&c.PlayerID,
			&c.IsOnline,

			&c.Raw.Fuel,
			&c.Raw.Bio,
			&c.Raw.CommonMetal,
			&c.Raw.MinorMetal,
			&c.Raw.PreciousMetal,

			&c.Processed.Fuel,
			&c.Processed.Bio,
			&c.Processed.CommonMetal,
			&c.Processed.MinorMetal,
			&c.Processed.PreciousMetal,

			&c.Plants.Wormwood,
			&c.Plants.BlackCarrot,
			&c.Plants.GoldenCrescent,
			&c.Plants.Tarragon,
			&c.Plants.AfricanPeach,
			&c.Plants.DigitalisPurpurea,
			&c.Plants.DigitalisLutea,
			&c.Plants.Haoma,

			&c.Vehicles.ZaAZS84,
			&c.Vehicles.APET41LV,
			&c.Vehicles.ZiGRA6T,
			&c.Vehicles.Boar53CT,
			&c.Vehicles.ZhukBr3,
			&c.Vehicles.StoutIfvSc,
			&c.Vehicles.ZhukRsZo,
			&c.Vehicles.StoutIfvFs,
			&c.Vehicles.TT77Nosorog,
			&c.Vehicles.M84AMagloader,

			&c.WalkerGears.WGPP,
			&c.WalkerGears.CCCPWGTypeC,
			&c.WalkerGears.CCCPWGTypeA,
			&c.WalkerGears.CFAWGTypeC,
			&c.WalkerGears.CFAWGTypeA,

			&c.Nuclear.Weapon,
			&c.Nuclear.Waste,

			&c.Parasites.Mist,
			&c.Parasites.Camouflage,
			&c.Parasites.Armor,

			&c.Mystery01,

			&c.Placed.VolgaK12,
			&c.Placed.HMG3Wingate,
			&c.Placed.M2A304Mortar,
			&c.Placed.Zhizdra45,
			&c.Placed.M276AAGGun,

			&c.Mystery02[0],
			&c.Mystery02[1],
			&c.Mystery02[2],
			&c.Mystery02[3],
			&c.Mystery02[4],
			&c.Mystery02[5],
			&c.Mystery02[6],

			&c.Unused[0],
			&c.Unused[1],
			&c.Unused[2],
			&c.Unused[3],
			&c.Unused[4],
			&c.Unused[5],
			&c.Unused[6],
			&c.Unused[7],
			&c.Unused[8],
			&c.Unused[9],
			&c.Unused[10],
			&c.Unused[11],
			&c.Unused[12],
		)

		res = append(res, c)
	}

	return res, nil
}
