CREATE TABLE favorite_movies (
    movie_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tmdb_id TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL
);


