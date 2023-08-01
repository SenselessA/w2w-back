CREATE TABLE movies (
  id               VARCHAR(255) PRIMARY KEY CONSTRAINT unique_id UNIQUE,
  kinopoisk_id     VARCHAR(255),
  imdb_id          VARCHAR(255),
  title            VARCHAR(255),
  title_orig       VARCHAR(255),
  other_title      VARCHAR(255),
  kodik_link       VARCHAR(255),
  worldart_link    VARCHAR(255),
  shikimori_id     VARCHAR(255),
  type             VARCHAR(255),
  quality          VARCHAR(255),
  lgbt             BOOLEAN,
  created_at       VARCHAR(255),
  updated_at       VARCHAR(255),
  blocked_countries TEXT[],
  seasons          JSONB,
  blocked_seasons  JSONB,
  last_season      INTEGER,
  last_episode     INTEGER,
  episodes_count   INTEGER,
  screenshots      TEXT[]
);

CREATE TABLE material_data (
 movie_id         VARCHAR(255) CONSTRAINT unique_movie_id UNIQUE,
 title_en         VARCHAR(255),
 other_titles     TEXT[],
 anime_licensed_by TEXT[],
 anime_kind       VARCHAR(255),
 all_status       VARCHAR(255),
 year             INTEGER,
 description      TEXT,
 poster_url       VARCHAR(255),
 screenshots      TEXT[],
 duration         INTEGER,
 countries        TEXT[],
 all_genres       TEXT[],
 kinopoisk_rating REAL,
 kinopoisk_votes  INTEGER,
 imdb_rating      REAL,
 imdb_votes       INTEGER,
 shikimori_rating REAL,
 shikimori_votes  INTEGER,
 premiere_ru      VARCHAR(255),
 premiere_world   VARCHAR(255),
 next_episode_at  VARCHAR(255),
 minimal_age      INTEGER,
 episodes_total   INTEGER,
 episodes_aired   INTEGER,
 actors           TEXT[],
 producers        TEXT[],
 FOREIGN KEY (movie_id) REFERENCES movies (id) ON DELETE CASCADE
);

ALTER TABLE user_movie_rating ADD FOREIGN KEY (movie_id) REFERENCES movies (id);
ALTER TABLE user_favorite_movies ADD FOREIGN KEY (movie_id) REFERENCES movies (id);



