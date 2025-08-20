package tppmessage

type CmdGetPlayerListEntry struct {
	EspionageLose int    `json:"espionage_lose"`
	EspionageWin  int    `json:"espionage_win"`
	FobGrade      int    `json:"fob_grade"`
	FobPoint      int    `json:"fob_point"`
	FobRank       int    `json:"fob_rank"`
	Index         int    `json:"index"`
	IsInsurance   int    `json:"is_insurance"`
	LeagueGrade   int    `json:"league_grade"`
	LeagueRank    int    `json:"league_rank"`
	Name          string `json:"name"`
	Playtime      int    `json:"playtime"`
	Point         int    `json:"point"`
}

type CmdGetPlayerListResponse struct {
	CryptoType string                  `json:"crypto_type"`
	Flowid     any                     `json:"flowid"`
	Msgid      string                  `json:"msgid"`
	PlayerList []CmdGetPlayerListEntry `json:"player_list"`
	PlayerNum  int                     `json:"player_num"`
	Result     string                  `json:"result"`
	Rqid       int                     `json:"rqid"`
	Xuid       any                     `json:"xuid"`
}

//type PlayerRepo struct {
//	db *sql.DB
//}
//
//func (p *PlayerRepo) CreateSchema(ctx context.Context) error {
//	schema := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
//	id INTEGER PRIMARY KEY AUTOINCREMENT,
//	steam_id INTEGER NOT NULL,
//	espionage_lose INTEGER,
//	espionage_win INTEGER,
//	fob_grade INTEGER,
//	fob_point INTEGER,
//	fob_rank INTEGER,
//	player_index INTEGER,
//	is_insurance INTEGER,
//	league_grade INTEGER,
//	league_rank INTEGER,
//	name TEXT,
//	playtime INTEGER,
//	point INTEGER
//);
//CREATE INDEX steam_id_idx ON %s (steam_id);
//`, p.TableName(), p.TableName())
//	_, err := p.db.ExecContext(ctx, schema)
//	if err != nil {
//		return fmt.Errorf("cannot create schema: %w", err)
//	}
//
//	return nil
//}
//
//func (p *PlayerRepo) TableName() string {
//	return "player"
//}
//
//func (p *PlayerRepo) WithDB(db *sql.DB) *PlayerRepo {
//	p.db = db
//	return p
//}
//
//func (p *PlayerRepo) Insert(ctx context.Context, steamID uint64) (uint64, error) {
//	query := fmt.Sprintf(`INSERT INTO %s (
//		steam_id,
//		espionage_lose,
//		espionage_win,
//		fob_grade,
//		fob_point,
//		fob_rank,
//		player_index,
//		is_insurance,
//		league_grade,
//		league_rank,
//		name,
//		playtime,
//		point) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`, p.TableName())
//
//	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{})
//	if err != nil {
//		return 0, err
//	}
//
//	_, err = tx.ExecContext(ctx, query,
//		steamID,
//		0, //e.EspionageLose,
//		0, //e.EspionageWin,
//		0, //e.FobGrade,
//		0, //e.FobPoint,
//		0, //e.FobRank,
//		1, //e.Index,
//		0, //e.IsInsurance,
//		0, //e.LeagueGrade,
//		0, //e.LeagueRank,
//		0, //e.Name,
//		0, //e.Playtime,
//		0, //e.Point
//	)
//
//	if err != nil {
//		return 0, err
//	}
//
//	q := fmt.Sprintf(`SELECT MAX(id) from %s;`, p.TableName())
//	row := tx.QueryRowContext(ctx, q)
//	pid := uint64(0)
//	if err = row.Scan(&pid); err != nil {
//		return 0, fmt.Errorf("get new player id: %w", err)
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		return 0, err
//	}
//
//	return pid, nil
//}
//
//func (p *PlayerRepo) GetAllByAccountID(ctx context.Context, steamID uint64) ([]CmdGetPlayerListEntry, error) {
//	res := make([]CmdGetPlayerListEntry, 0)
//	query := fmt.Sprintf(`select
//		id,
//		espionage_lose,
//		espionage_win,
//		fob_grade,
//		fob_point,
//		fob_rank,
//		player_index,
//		is_insurance,
//		league_grade,
//		league_rank,
//		name,
//		playtime,
//		point from %s where steam_id = ?`, p.TableName())
//
//	rows, err := p.db.QueryContext(ctx, query, steamID)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var entry CmdGetPlayerListEntry
//		err = rows.Scan(&entry.PlayerID, &entry.EspionageLose, &entry.EspionageWin, &entry.FobGrade,
//			&entry.FobPoint, &entry.FobRank, &entry.Index, &entry.IsInsurance, &entry.LeagueGrade,
//			&entry.LeagueRank, &entry.Name, &entry.Playtime, &entry.Point)
//
//		if err != nil {
//			return nil, err
//		}
//
//		res = append(res, entry)
//	}
//
//	return res, nil
//}
//
//func (p *PlayerRepo) Get(ctx context.Context, steamID uint64, index int) (CmdGetPlayerListEntry, error) {
//	query := fmt.Sprintf(`select
//		id,
//		espionage_lose,
//		espionage_win,
//		fob_grade,
//		fob_point,
//		fob_rank,
//		player_index,
//		is_insurance,
//		league_grade,
//		league_rank,
//		name,
//		playtime,
//		point from %s where steam_id = ? and player_index = ?`, p.TableName())
//
//	var entry CmdGetPlayerListEntry
//	err := p.db.QueryRowContext(ctx, query, steamID, index).Scan(&entry.PlayerID, &entry.EspionageLose, &entry.EspionageWin, &entry.FobGrade,
//		&entry.FobPoint, &entry.FobRank, &entry.Index, &entry.IsInsurance, &entry.LeagueGrade,
//		&entry.LeagueRank, &entry.Name, &entry.Playtime, &entry.Point)
//
//	if err != nil {
//		return CmdGetPlayerListEntry{}, err
//	}
//
//	return entry, nil
//}
