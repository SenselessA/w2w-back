package repository

import (
	"database/sql"
	"github.com/SenselessA/w2w_backend/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type RepoMovies struct {
	db *sqlx.DB
}

func initMovies(db *sqlx.DB) *RepoMovies {
	return &RepoMovies{db: db}
}

func (r *RepoMovies) GetByMovieId(input string) (*models.MovieDataByMovieIdOutput, error) {

	var movieData models.MovieDataByMovieIdOutput

	err := r.db.QueryRowx(`
		SELECT m.id, md.kinopoisk_rating, md.imdb_rating, md.kinopoisk_votes, md.imdb_votes,
		       md.poster_url, m.title, m.other_title, m.title_orig, m.type, m.episodes_count,
		       md.all_status, md.year, md.all_genres, md.minimal_age, md.duration, 
		       md.countries, md.actors, md.description, m.kodik_link
		FROM movies as m, material_data as md
		         WHERE $1 = m.id AND $1 = md.movie_id
	`, input).StructScan(&movieData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &movieData, nil
}

func (r *RepoMovies) GetLastUpdatedAt() (*string, error) {
	var lastUpdatedAtMovie string

	err := r.db.QueryRow(`
		SELECT updated_at FROM movies ORDER BY updated_at DESC
	`).Scan(&lastUpdatedAtMovie)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &lastUpdatedAtMovie, nil
}

func (r *RepoMovies) BatchInsertAndUpdate(movies *[]models.Movie) error {
	moviesList := *movies

	//var preparedMovies []map[string]interface{}
	//var preparedMaterials []map[string]interface{}
	//var materials []models.MaterialDataInput
	//var moviesOnly []models.MovieInput
	// мб зарефачить на pgx потом, вроде он может батчинг сделать https://stackoverflow.com/questions/70823061/bulk-insert-in-postgres-in-go-using-pgx
	//ctx := context.Background()
	//tx, err := pgx.Tx().Begin(ctx)

	for _, movie := range moviesList {

		material := models.MaterialDataInput{
			MovieId:         movie.Id,
			TitleEn:         movie.MaterialData.TitleEn,
			OtherTitles:     movie.MaterialData.OtherTitles,
			AnimeLicensedBy: movie.MaterialData.AnimeLicensedBy,
			AnimeKind:       movie.MaterialData.AnimeKind,
			AllStatus:       movie.MaterialData.AllStatus,
			Year:            movie.MaterialData.Year,
			Description:     movie.MaterialData.Description,
			PosterUrl:       movie.MaterialData.PosterUrl,
			Screenshots:     movie.MaterialData.Screenshots,
			Duration:        movie.MaterialData.Duration,
			Countries:       movie.MaterialData.Countries,
			AllGenres:       movie.MaterialData.AllGenres,
			KinopoiskRating: movie.MaterialData.KinopoiskRating,
			KinopoiskVotes:  movie.MaterialData.KinopoiskVotes,
			ImdbRating:      movie.MaterialData.ImdbRating,
			ImdbVotes:       movie.MaterialData.ImdbVotes,
			ShikimoriRating: movie.MaterialData.ShikimoriRating,
			ShikimoriVotes:  movie.MaterialData.ShikimoriVotes,
			PremiereRu:      movie.MaterialData.PremiereRu,
			PremiereWorld:   movie.MaterialData.PremiereWorld,
			NextEpisodeAt:   movie.MaterialData.NextEpisodeAt,
			MinimalAge:      movie.MaterialData.MinimalAge,
			EpisodesTotal:   movie.MaterialData.EpisodesTotal,
			EpisodesAired:   movie.MaterialData.EpisodesAired,
			Actors:          movie.MaterialData.Actors,
			Producers:       movie.MaterialData.Producers,
		}

		tx := r.db.MustBegin()

		_, err := tx.NamedExec(`
		INSERT INTO movies
		    (id, kinopoisk_id, imdb_id, title, title_orig, other_title, kodik_link, worldart_link, shikimori_id, type,
			quality, lgbt, created_at, updated_at, blocked_countries, seasons, blocked_seasons, last_season,
			last_episode, episodes_count, screenshots)
		VALUES
		    (:id, :kinopoisk_id, :imdb_id, :title, :title_orig, :other_title, :kodik_link, :worldart_link, :shikimori_id, :type,
			:quality, :lgbt, :created_at, :updated_at, :blocked_countries, :seasons, :blocked_seasons, :last_season,
			:last_episode, :episodes_count, :screenshots)
		ON CONFLICT (id) DO UPDATE SET
		   (id, kinopoisk_id, imdb_id, title, title_orig, other_title, kodik_link, worldart_link, shikimori_id, type,
			quality, lgbt, created_at, updated_at, blocked_countries, seasons, blocked_seasons, last_season,
			last_episode, episodes_count, screenshots)
		       =
		   (:id, :kinopoisk_id, :imdb_id, :title, :title_orig, :other_title, :kodik_link, :worldart_link, :shikimori_id, :type,
			:quality, :lgbt, :created_at, :updated_at, :blocked_countries, :seasons, :blocked_seasons, :last_season,
			:last_episode, :episodes_count, :screenshots);
	`, movie)
		if err != nil {
			log.Println("tx.NamedExec movie err: ", err)
			err := tx.Rollback()
			if err != nil {
				return err
			}
		}

		_, err = tx.NamedExec(`
		INSERT INTO material_data
		    (movie_id, title_en, other_titles, anime_licensed_by, anime_kind, all_status, year, description, poster_url,
		     screenshots, duration, countries, all_genres, kinopoisk_rating, kinopoisk_votes, imdb_rating, imdb_votes,
		     shikimori_rating, shikimori_votes, premiere_ru, premiere_world, next_episode_at, minimal_age, episodes_total,
		     episodes_aired, actors, producers)
		VALUES
		    (:movie_id, :title_en, :other_titles, :anime_licensed_by, :anime_kind, :all_status, :year, :description, :poster_url,
		     :screenshots, :duration, :countries, :all_genres, :kinopoisk_rating, :kinopoisk_votes, :imdb_rating, :imdb_votes,
		     :shikimori_rating, :shikimori_votes, :premiere_ru, :premiere_world, :next_episode_at, :minimal_age, :episodes_total,
		     :episodes_aired, :actors, :producers)
		ON CONFLICT (movie_id) DO UPDATE SET
		   (movie_id, title_en, other_titles, anime_licensed_by, anime_kind, all_status, year, description, poster_url,
		     screenshots, duration, countries, all_genres, kinopoisk_rating, kinopoisk_votes, imdb_rating, imdb_votes,
		     shikimori_rating, shikimori_votes, premiere_ru, premiere_world, next_episode_at, minimal_age, episodes_total,
		     episodes_aired, actors, producers)
		       =
		   (:movie_id, :title_en, :other_titles, :anime_licensed_by, :anime_kind, :all_status, :year, :description, :poster_url,
		     :screenshots, :duration, :countries, :all_genres, :kinopoisk_rating, :kinopoisk_votes, :imdb_rating, :imdb_votes,
		     :shikimori_rating, :shikimori_votes, :premiere_ru, :premiere_world, :next_episode_at, :minimal_age, :episodes_total,
		     :episodes_aired, :actors, :producers);
	`, material)
		if err != nil {
			log.Println("tx.NamedExec material_data err: ", err)
			err := tx.Rollback()
			if err != nil {
				return err
			}
		}

		err = tx.Commit()
		if err != nil {
			println("BatchInsertAndUpdate tx commit err: ", err)
			return err
		}

		//materials = append(materials, material)
		//
		//movieOnly := models.MovieInput{
		//	Id:               movie.Id,
		//	KinopoiskId:      movie.KinopoiskId,
		//	ImdbId:           movie.ImdbId,
		//	Title:            movie.Title,
		//	TitleOrig:        movie.TitleOrig,
		//	OtherTitle:       movie.OtherTitle,
		//	KodikLink:        movie.KodikLink,
		//	WorldartLink:     movie.WorldartLink,
		//	ShikimoriId:      movie.ShikimoriId,
		//	Type:             movie.Type,
		//	Quality:          movie.Quality,
		//	Lgbt:             movie.Lgbt,
		//	CreatedAt:        movie.CreatedAt,
		//	UpdatedAt:        movie.UpdatedAt,
		//	BlockedCountries: movie.BlockedCountries,
		//	Seasons:          movie.Seasons,
		//	BlockedSeasons:   movie.BlockedSeasons,
		//	LastSeason:       movie.LastSeason,
		//	LastEpisode:      movie.LastEpisode,
		//	EpisodesCount:    movie.EpisodesCount,
		//	Screenshots:      movie.Screenshots,
		//}
		//
		//moviesOnly = append(moviesOnly, movieOnly)
		//
		//preparedMovies = append(preparedMovies, map[string]interface{}{
		//	"id":                movie.Id,
		//	"kinopoisk_id":      movie.KinopoiskId,
		//	"imdb_id":           movie.ImdbId,
		//	"title":             movie.Title,
		//	"title_orig":        movie.TitleOrig,
		//	"other_title":       movie.OtherTitle,
		//	"kodik_link":        movie.KodikLink,
		//	"worldart_link":     movie.WorldartLink,
		//	"shikimori_id":      movie.ShikimoriId,
		//	"type":              movie.Type,
		//	"quality":           movie.Quality,
		//	"lgbt":              movie.Lgbt,
		//	"created_at":        movie.CreatedAt,
		//	"updated_at":        movie.UpdatedAt,
		//	"blocked_countries": movie.BlockedCountries,
		//	"seasons":           movie.Seasons,
		//	"blocked_seasons":   movie.BlockedSeasons,
		//	"last_season":       movie.LastSeason,
		//	"last_episode":      movie.LastEpisode,
		//	"episodes_count":    movie.EpisodesCount,
		//	"screenshots":       movie.Screenshots,
		//})
		//
		//preparedMaterials = append(preparedMaterials, map[string]interface{}{
		//	"movie_id":          movie.Id,
		//	"title_en":          movie.MaterialData.TitleEn,
		//	"other_titles":      movie.MaterialData.OtherTitles,
		//	"anime_licensed_by": movie.MaterialData.AnimeLicensedBy,
		//	"anime_kind":        movie.MaterialData.AnimeKind,
		//	"all_status":        movie.MaterialData.AllStatus,
		//	"year":              movie.MaterialData.Year,
		//	"description":       movie.MaterialData.Description,
		//	"poster_url":        movie.MaterialData.PosterUrl,
		//	"screenshots":       movie.MaterialData.Screenshots,
		//	"duration":          movie.MaterialData.Duration,
		//	"countries":         movie.MaterialData.Countries,
		//	"all_genres":        movie.MaterialData.AllGenres,
		//	"kinopoisk_rating":  movie.MaterialData.KinopoiskRating,
		//	"kinopoisk_votes":   movie.MaterialData.KinopoiskVotes,
		//	"imdb_rating":       movie.MaterialData.ImdbRating,
		//	"imdb_votes":        movie.MaterialData.ImdbVotes,
		//	"shikimori_rating":  movie.MaterialData.ShikimoriRating,
		//	"shikimori_votes":   movie.MaterialData.ShikimoriVotes,
		//	"premiere_ru":       movie.MaterialData.PremiereRu,
		//	"premiere_world":    movie.MaterialData.PremiereWorld,
		//	"next_episode_at":   movie.MaterialData.NextEpisodeAt,
		//	"minimal_age":       movie.MaterialData.MinimalAge,
		//	"episodes_total":    movie.MaterialData.EpisodesTotal,
		//	"episodes_aired":    movie.MaterialData.EpisodesAired,
		//	"actors":            movie.MaterialData.Actors,
		//	"producers":         movie.MaterialData.Producers,
		//})
	}
	return nil
}
