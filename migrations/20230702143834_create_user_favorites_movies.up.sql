CREATE TABLE user_favorite_movies (
      user_id BIGSERIAL NOT NULL,
      movie_id VARCHAR(255) NOT NULL,
      FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
      UNIQUE (user_id, movie_id)
);



