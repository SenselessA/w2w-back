CREATE TABLE user_movie_rating (
    user_id BIGSERIAL NOT NULL,
    movie_id VARCHAR(255) NOT NULL,
    rating INTEGER CHECK (rating >= 0 AND rating <= 10) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (user_id, movie_id)
);
