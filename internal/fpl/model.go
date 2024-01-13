package fpl

import (
	"fmt"
	"strings"
	"time"
)

type status struct {
	BonusAdded bool   `json:"bonus_added"`
	Date       string `json:"date"`
	Event      int    `json:"event"`
	Points     string `json:"points"`
}

type EventStatus struct {
	Status  []status `json:"status"`
	Leagues string   `json:"leagues"`
}

func (es EventStatus) DateInCurrentEvent(date time.Time) bool {
	for _, status := range es.Status {
		eventDate, err := time.Parse("2006-01-02", status.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

		if sameDate(eventDate, date) {
			return true
		}
	}
	return false
}

func (es EventStatus) BonusAdded(date time.Time) bool {
	if !es.DateInCurrentEvent(date) {
		return false
	}

	for _, status := range es.Status {
		eventDate, err := time.Parse("2006-01-02", status.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

		if sameDate(eventDate, date) && status.BonusAdded {
			return true
		}
	}
	return false
}

func sameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

type event struct {
	ID                     int            `json:"id"`
	Name                   string         `json:"name"`
	DeadlineTime           time.Time      `json:"deadline_time"`
	AverageEntryScore      int            `json:"average_entry_score"`
	Finished               bool           `json:"finished"`
	DataChecked            bool           `json:"data_checked"`
	HighestScoringEntry    int            `json:"highest_scoring_entry"`
	DeadlineTimeEpoch      int            `json:"deadline_time_epoch"`
	DeadlineTimeGameOffset int            `json:"deadline_time_game_offset"`
	HighestScore           int            `json:"highest_score"`
	IsPrevious             bool           `json:"is_previous"`
	IsCurrent              bool           `json:"is_current"`
	IsNext                 bool           `json:"is_next"`
	CupLeaguesCreated      bool           `json:"cup_leagues_created"`
	H2HKoMatchesCreated    bool           `json:"h2h_ko_matches_created"`
	ChipPlays              []chipPlay     `json:"chip_plays"`
	MostSelected           int            `json:"most_selected"`
	MostTransferredIn      int            `json:"most_transferred_in"`
	TopElement             int            `json:"top_element"`
	TopElementInfo         topElementInfo `json:"top_element_info"`
	TransfersMade          int            `json:"transfers_made"`
	MostCaptained          int            `json:"most_captained"`
	MostViceCaptained      int            `json:"most_vice_captained"`
}

type chipPlay struct {
	ChipName  string `json:"chip_name"`
	NumPlayed int    `json:"num_played"`
}

type topElementInfo struct {
	ID     int `json:"id"`
	Points int `json:"points"`
}

type gameSettings struct {
	LeagueJoinPrivateMax         int           `json:"league_join_private_max"`
	LeagueJoinPublicMax          int           `json:"league_join_public_max"`
	LeagueMaxSizePublicClassic   int           `json:"league_max_size_public_classic"`
	LeagueMaxSizePublicH2H       int           `json:"league_max_size_public_h2h"`
	LeagueMaxSizePrivateH2H      int           `json:"league_max_size_private_h2h"`
	LeagueMaxKoRoundsPrivateH2H  int           `json:"league_max_ko_rounds_private_h2h"`
	LeaguePrefixPublic           string        `json:"league_prefix_public"`
	LeaguePointsH2HWin           int           `json:"league_points_h2h_win"`
	LeaguePointsH2HLose          int           `json:"league_points_h2h_lose"`
	LeaguePointsH2HDraw          int           `json:"league_points_h2h_draw"`
	LeagueKoFirstInsteadOfRandom bool          `json:"league_ko_first_instead_of_random"`
	CupStartEventID              interface{}   `json:"cup_start_event_id"`
	CupStopEventID               interface{}   `json:"cup_stop_event_id"`
	CupQualifyingMethod          interface{}   `json:"cup_qualifying_method"`
	CupType                      interface{}   `json:"cup_type"`
	FeaturedEntries              []interface{} `json:"featured_entries"`
	PercentileRanks              []int         `json:"percentile_ranks"`
	SquadSquadplay               int           `json:"squad_squadplay"`
	SquadSquadsize               int           `json:"squad_squadsize"`
	SquadTeamLimit               int           `json:"squad_team_limit"`
	SquadTotalSpend              int           `json:"squad_total_spend"`
	UICurrencyMultiplier         int           `json:"ui_currency_multiplier"`
	UIUseSpecialShirts           bool          `json:"ui_use_special_shirts"`
	UISpecialShirtExclusions     []interface{} `json:"ui_special_shirt_exclusions"`
	StatsFormDays                int           `json:"stats_form_days"`
	SysViceCaptainEnabled        bool          `json:"sys_vice_captain_enabled"`
	TransfersCap                 int           `json:"transfers_cap"`
	TransfersSellOnFee           float64       `json:"transfers_sell_on_fee"`
	LeagueH2HTiebreakStats       []string      `json:"league_h2h_tiebreak_stats"`
	Timezone                     string        `json:"timezone"`
}

type phase struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StartEvent int    `json:"start_event"`
	StopEvent  int    `json:"stop_event"`
}

type team struct {
	Code                int         `json:"code"`
	Draw                int         `json:"draw"`
	Form                interface{} `json:"form"`
	ID                  int         `json:"id"`
	Loss                int         `json:"loss"`
	Name                string      `json:"name"`
	Played              int         `json:"played"`
	Points              int         `json:"points"`
	Position            int         `json:"position"`
	ShortName           string      `json:"short_name"`
	Strength            int         `json:"strength"`
	TeamDivision        interface{} `json:"team_division"`
	Unavailable         bool        `json:"unavailable"`
	Win                 int         `json:"win"`
	StrengthOverallHome int         `json:"strength_overall_home"`
	StrengthOverallAway int         `json:"strength_overall_away"`
	StrengthAttackHome  int         `json:"strength_attack_home"`
	StrengthAttackAway  int         `json:"strength_attack_away"`
	StrengthDefenceHome int         `json:"strength_defence_home"`
	StrengthDefenceAway int         `json:"strength_defence_away"`
	PulseID             int         `json:"pulse_id"`
}

type element struct {
	ChanceOfPlayingNextRound         int         `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound         int         `json:"chance_of_playing_this_round"`
	Code                             int         `json:"code"`
	CostChangeEvent                  int         `json:"cost_change_event"`
	CostChangeEventFall              int         `json:"cost_change_event_fall"`
	CostChangeStart                  int         `json:"cost_change_start"`
	CostChangeStartFall              int         `json:"cost_change_start_fall"`
	DreamteamCount                   int         `json:"dreamteam_count"`
	ElementType                      int         `json:"element_type"`
	EpNext                           string      `json:"ep_next"`
	EpThis                           string      `json:"ep_this"`
	EventPoints                      int         `json:"event_points"`
	FirstName                        string      `json:"first_name"`
	Form                             string      `json:"form"`
	ID                               int         `json:"id"`
	InDreamteam                      bool        `json:"in_dreamteam"`
	News                             string      `json:"news"`
	NewsAdded                        time.Time   `json:"news_added"`
	NowCost                          int         `json:"now_cost"`
	Photo                            string      `json:"photo"`
	PointsPerGame                    string      `json:"points_per_game"`
	SecondName                       string      `json:"second_name"`
	SelectedByPercent                string      `json:"selected_by_percent"`
	Special                          bool        `json:"special"`
	SquadNumber                      interface{} `json:"squad_number"`
	Status                           string      `json:"status"`
	Team                             int         `json:"team"`
	TeamCode                         int         `json:"team_code"`
	TotalPoints                      int         `json:"total_points"`
	TransfersIn                      int         `json:"transfers_in"`
	TransfersInEvent                 int         `json:"transfers_in_event"`
	TransfersOut                     int         `json:"transfers_out"`
	TransfersOutEvent                int         `json:"transfers_out_event"`
	ValueForm                        string      `json:"value_form"`
	ValueSeason                      string      `json:"value_season"`
	WebName                          string      `json:"web_name"`
	Minutes                          int         `json:"minutes"`
	GoalsScored                      int         `json:"goals_scored"`
	Assists                          int         `json:"assists"`
	CleanSheets                      int         `json:"clean_sheets"`
	GoalsConceded                    int         `json:"goals_conceded"`
	OwnGoals                         int         `json:"own_goals"`
	PenaltiesSaved                   int         `json:"penalties_saved"`
	PenaltiesMissed                  int         `json:"penalties_missed"`
	YellowCards                      int         `json:"yellow_cards"`
	RedCards                         int         `json:"red_cards"`
	Saves                            int         `json:"saves"`
	Bonus                            int         `json:"bonus"`
	Bps                              int         `json:"bps"`
	Influence                        string      `json:"influence"`
	Creativity                       string      `json:"creativity"`
	Threat                           string      `json:"threat"`
	IctIndex                         string      `json:"ict_index"`
	Starts                           int         `json:"starts"`
	ExpectedGoals                    string      `json:"expected_goals"`
	ExpectedAssists                  string      `json:"expected_assists"`
	ExpectedGoalInvolvements         string      `json:"expected_goal_involvements"`
	ExpectedGoalsConceded            string      `json:"expected_goals_conceded"`
	InfluenceRank                    int         `json:"influence_rank"`
	InfluenceRankType                int         `json:"influence_rank_type"`
	CreativityRank                   int         `json:"creativity_rank"`
	CreativityRankType               int         `json:"creativity_rank_type"`
	ThreatRank                       int         `json:"threat_rank"`
	ThreatRankType                   int         `json:"threat_rank_type"`
	IctIndexRank                     int         `json:"ict_index_rank"`
	IctIndexRankType                 int         `json:"ict_index_rank_type"`
	CornersAndIndirectFreekicksOrder interface{} `json:"corners_and_indirect_freekicks_order"`
	CornersAndIndirectFreekicksText  string      `json:"corners_and_indirect_freekicks_text"`
	DirectFreekicksOrder             interface{} `json:"direct_freekicks_order"`
	DirectFreekicksText              string      `json:"direct_freekicks_text"`
	PenaltiesOrder                   interface{} `json:"penalties_order"`
	PenaltiesText                    string      `json:"penalties_text"`
	ExpectedGoalsPer90               float32     `json:"expected_goals_per_90"`
	SavesPer90                       float32     `json:"saves_per_90"`
	ExpectedAssistsPer90             float32     `json:"expected_assists_per_90"`
	ExpectedGoalInvolvementsPer90    float32     `json:"expected_goal_involvements_per_90"`
	ExpectedGoalsConcededPer90       float32     `json:"expected_goals_conceded_per_90"`
	GoalsConcededPer90               float32     `json:"goals_conceded_per_90"`
	NowCostRank                      int         `json:"now_cost_rank"`
	NowCostRankType                  int         `json:"now_cost_rank_type"`
	FormRank                         int         `json:"form_rank"`
	FormRankType                     int         `json:"form_rank_type"`
	PointsPerGameRank                int         `json:"points_per_game_rank"`
	PointsPerGameRankType            int         `json:"points_per_game_rank_type"`
	SelectedRank                     int         `json:"selected_rank"`
	SelectedRankType                 int         `json:"selected_rank_type"`
	StartsPer90                      float32     `json:"starts_per_90"`
	CleanSheetsPer90                 float32     `json:"clean_sheets_per_90"`
}

type elementStat struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type elementType struct {
	ID                 int    `json:"id"`
	PluralName         string `json:"plural_name"`
	PluralNameShort    string `json:"plural_name_short"`
	SingularName       string `json:"singular_name"`
	SingularNameShort  string `json:"singular_name_short"`
	SquadSelect        int    `json:"squad_select"`
	SquadMinPlay       int    `json:"squad_min_play"`
	SquadMaxPlay       int    `json:"squad_max_play"`
	UIShirtSpecific    bool   `json:"ui_shirt_specific"`
	SubPositionsLocked []int  `json:"sub_positions_locked"`
	ElementCount       int    `json:"element_count"`
}

type LeagueBootstrap struct {
	Events       []event       `json:"events"`
	GameSettings gameSettings  `json:"game_settings"`
	Phases       []phase       `json:"phases"`
	Teams        []team        `json:"teams"`
	TotalPlayers int           `json:"total_players"`
	Elements     []element     `json:"elements"`
	ElementStats []elementStat `json:"element_stats"`
	ElementTypes []elementType `json:"element_types"`
}

type league struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

type playerStanding struct {
	ID         int    `json:"id"`
	EventTotal int    `json:"event_total"`
	PlayerName string `json:"player_name"`
	Rank       int    `json:"rank"`
	LastRank   int    `json:"last_rank"`
	RankSort   int    `json:"rank_sort"`
	Total      int    `json:"total"`
	Entry      int    `json:"entry"`
	EntryName  string `json:"entry_name"`
}

type standings struct {
	HasNext bool             `json:"has_next"`
	Page    int              `json:"page"`
	Results []playerStanding `json:"results"`
}

type LeagueData struct {
	League      league    `json:"league"`
	Standings   standings `json:"standings"`
	LastUpdated time.Time `json:"last_updated_data"`
}

func (ld LeagueData) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("🏆 *%s*\n\n", ld.League.Name))

	for _, s := range ld.Standings.Results {
		movementEmoji := getMovementEmoji(s.Rank, s.LastRank)
		medalEmoji := getMedalEmoji(s.Rank)

		line := fmt.Sprintf("%s _*%d*_. %s: *%d*", movementEmoji, s.Rank, s.EntryName, s.Total)
		if medalEmoji != "" {
			line = line + " " + medalEmoji
		}

		sb.WriteString(line + "\n")
	}

	sb.WriteString(fmt.Sprintf("\n🤖 _Automated message from FPL Go Bot_"))

	return sb.String()
}

func getMovementEmoji(currentPos, prevPos int) string {
	switch {
	case currentPos < prevPos:
		return "🔼"
	case currentPos > prevPos:
		return "🔽"
	default:
		return "⏺️"
	}
}

func getMedalEmoji(pos int) string {
	switch {
	case pos == 1:
		return "🥇"
	case pos == 2:
		return "🥈"
	case pos == 3:
		return "🥉"
	default:
		return ""
	}
}
