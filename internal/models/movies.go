package models

import "github.com/lib/pq"

type Movie struct {
	Id               string       `json:"id" db:"id"`
	KinopoiskId      string       `json:"kinopoisk_id" db:"kinopoisk_id"`
	ImdbId           string       `json:"imdb_id" db:"imdb_id"`
	Title            string       `json:"title" db:"title"`
	TitleOrig        string       `json:"title_orig" db:"title_orig"`
	OtherTitle       string       `json:"other_title" db:"other_title"`
	KodikLink        string       `json:"link" db:"kodik_link"`
	WorldartLink     string       `json:"worldart_link" db:"worldart_link"`
	ShikimoriId      string       `json:"shikimori_id" db:"shikimori_id"`
	Type             string       `json:"type" db:"type"`
	Quality          string       `json:"quality" db:"quality"`
	Lgbt             bool         `json:"lgbt" db:"lgbt"`
	CreatedAt        string       `json:"created_at" db:"created_at"`
	UpdatedAt        string       `json:"updated_at" db:"updated_at"`
	BlockedCountries []string     `json:"blocked_countries" db:"blocked_countries"`
	Seasons          interface{}  `json:"seasons" db:"seasons" dbx:"jsonb"`
	BlockedSeasons   interface{}  `json:"blocked_seasons" db:"blocked_seasons" dbx:"jsonb"`
	LastSeason       int          `json:"last_season" db:"last_season"`
	LastEpisode      int          `json:"last_episode" db:"last_episode"`
	EpisodesCount    int          `json:"episodes_count" db:"episodes_count"`
	Screenshots      []string     `json:"screenshots" db:"screenshots"`
	MaterialData     MaterialData `json:"material_data" db:"material_data"`
}

type MovieInput struct {
	Id               string      `json:"id" db:"id"`
	KinopoiskId      string      `json:"kinopoisk_id" db:"kinopoisk_id"`
	ImdbId           string      `json:"imdb_id" db:"imdb_id"`
	Title            string      `json:"title" db:"title"`
	TitleOrig        string      `json:"title_orig" db:"title_orig"`
	OtherTitle       string      `json:"other_title" db:"other_title"`
	KodikLink        string      `json:"link" db:"kodik_link"`
	WorldartLink     string      `json:"worldart_link" db:"worldart_link"`
	ShikimoriId      string      `json:"shikimori_id" db:"shikimori_id"`
	Type             string      `json:"type" db:"type"`
	Quality          string      `json:"quality" db:"quality"`
	Lgbt             bool        `json:"lgbt" db:"lgbt"`
	CreatedAt        string      `json:"created_at" db:"created_at"`
	UpdatedAt        string      `json:"updated_at" db:"updated_at"`
	BlockedCountries []string    `json:"blocked_countries" db:"blocked_countries"`
	Seasons          interface{} `json:"seasons" db:"seasons" dbx:"jsonb"`
	BlockedSeasons   interface{} `json:"blocked_seasons" db:"blocked_seasons" dbx:"jsonb"`
	LastSeason       int         `json:"last_season" db:"last_season"`
	LastEpisode      int         `json:"last_episode" db:"last_episode"`
	EpisodesCount    int         `json:"episodes_count" db:"episodes_count"`
	Screenshots      []string    `json:"screenshots" db:"screenshots"`
}

type MaterialData struct {
	TitleEn         string   `json:"title_en" db:"title_en"`
	OtherTitles     []string `json:"other_titles" db:"other_titles"`
	AnimeLicensedBy []string `json:"anime_licensed_by" db:"anime_licensed_by"`
	AnimeKind       string   `json:"anime_kind" db:"anime_kind"`
	AllStatus       string   `json:"all_status" db:"all_status"`
	Year            uint     `json:"year" db:"year"`
	Description     string   `json:"description" db:"description"`
	PosterUrl       string   `json:"poster_url" db:"poster_url"`
	Screenshots     []string `json:"screenshots" db:"screenshots"`
	Duration        uint     `json:"duration" db:"duration"`
	Countries       []string `json:"countries" db:"countries"`
	AllGenres       []string `json:"all_genres" db:"all_genres"`
	KinopoiskRating float32  `json:"kinopoisk_rating" db:"kinopoisk_rating"`
	KinopoiskVotes  uint     `json:"kinopoisk_votes" db:"kinopoisk_votes"`
	ImdbRating      float32  `json:"imdb_rating" db:"imdb_rating"`
	ImdbVotes       uint     `json:"imdb_votes" db:"imdb_votes"`
	ShikimoriRating float32  `json:"shikimori_rating" db:"shikimori_rating"`
	ShikimoriVotes  uint     `json:"shikimori_votes" db:"shikimori_votes"`
	PremiereRu      string   `json:"premiere_ru" db:"premiere_ru"`
	PremiereWorld   string   `json:"premiere_world" db:"premiere_world"`
	NextEpisodeAt   string   `json:"next_episode_at" db:"next_episode_at"`
	MinimalAge      uint     `json:"minimal_age" db:"minimal_age"`
	EpisodesTotal   uint     `json:"episodes_total" db:"episodes_total"`
	EpisodesAired   uint     `json:"episodes_aired" db:"episodes_aired"`
	Actors          []string `json:"actors" db:"actors"`
	Producers       []string `json:"producers" db:"producers"`
}

type MaterialDataInput struct {
	MovieId         string   `json:"movie_id" db:"movie_id"`
	TitleEn         string   `json:"title_en" db:"title_en"`
	OtherTitles     []string `json:"other_titles" db:"other_titles"`
	AnimeLicensedBy []string `json:"anime_licensed_by" db:"anime_licensed_by"`
	AnimeKind       string   `json:"anime_kind" db:"anime_kind"`
	AllStatus       string   `json:"all_status" db:"all_status"`
	Year            uint     `json:"year" db:"year"`
	Description     string   `json:"description" db:"description"`
	PosterUrl       string   `json:"poster_url" db:"poster_url"`
	Screenshots     []string `json:"screenshots" db:"screenshots"`
	Duration        uint     `json:"duration" db:"duration"`
	Countries       []string `json:"countries" db:"countries"`
	AllGenres       []string `json:"all_genres" db:"all_genres"`
	KinopoiskRating float32  `json:"kinopoisk_rating" db:"kinopoisk_rating"`
	KinopoiskVotes  uint     `json:"kinopoisk_votes" db:"kinopoisk_votes"`
	ImdbRating      float32  `json:"imdb_rating" db:"imdb_rating"`
	ImdbVotes       uint     `json:"imdb_votes" db:"imdb_votes"`
	ShikimoriRating float32  `json:"shikimori_rating" db:"shikimori_rating"`
	ShikimoriVotes  uint     `json:"shikimori_votes" db:"shikimori_votes"`
	PremiereRu      string   `json:"premiere_ru" db:"premiere_ru"`
	PremiereWorld   string   `json:"premiere_world" db:"premiere_world"`
	NextEpisodeAt   string   `json:"next_episode_at" db:"next_episode_at"`
	MinimalAge      uint     `json:"minimal_age" db:"minimal_age"`
	EpisodesTotal   uint     `json:"episodes_total" db:"episodes_total"`
	EpisodesAired   uint     `json:"episodes_aired" db:"episodes_aired"`
	Actors          []string `json:"actors" db:"actors"`
	Producers       []string `json:"producers" db:"producers"`
}

type MovieDataByMovieIdOutput struct {
	Id              string         `json:"id" db:"id"`
	Title           string         `json:"title" db:"title"`
	TitleOrig       string         `json:"titleOrig" db:"title_orig"`
	OtherTitle      string         `json:"otherTitle" db:"other_title"`
	KodikLink       string         `json:"link" db:"kodik_link"`
	Type            string         `json:"type" db:"type"`
	EpisodesCount   int            `json:"episodesCount" db:"episodes_count"`
	KinopoiskRating float32        `json:"kinopoiskRating" db:"kinopoisk_rating"`
	ImdbRating      float32        `json:"imdbRating" db:"imdb_rating"`
	KinopoiskVotes  uint           `json:"kinopoiskVotes" db:"kinopoisk_votes"`
	ImdbVotes       uint           `json:"imdbVotes" db:"imdb_votes"`
	PosterUrl       string         `json:"posterUrl" db:"poster_url"`
	AllStatus       string         `json:"allStatus" db:"all_status"`
	Year            uint           `json:"year" db:"year"`
	AllGenres       pq.StringArray `json:"genres" db:"all_genres"`
	MinimalAge      uint           `json:"minimalAge" db:"minimal_age"`
	Description     string         `json:"description" db:"description"`
	Countries       pq.StringArray `json:"countries,omitempty" db:"countries"`
	Actors          pq.StringArray `json:"actors,omitempty" db:"actors"`
	Duration        uint           `json:"duration" db:"duration"`
}
