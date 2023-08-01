ALTER TABLE user_movie_rating DROP CONSTRAINT user_movie_rating_movie_id_fkey;
ALTER TABLE user_favorite_movies DROP CONSTRAINT user_favorite_movies_movie_id_key;

DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS material_data;