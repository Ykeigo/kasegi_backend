CREATE TABLE IF NOT EXISTS users(
    id VARCHAR (50) PRIMARY KEY,
    user_name VARCHAR (50) NOT NULL
);
create index on users(id);

CREATE TABLE IF NOT EXISTS game_titles(
    id serial PRIMARY KEY,
    title VARCHAR (50) NOT NULL
);
create index on game_titles(id);

CREATE TABLE IF NOT EXISTS game_matches(
    id serial PRIMARY KEY,
    game_id int NOT NULL,
    user_id VARCHAR (50) NOT NULL,
    CONSTRAINT game_matches_game_titles_id_fk FOREIGN KEY (game_id) REFERENCES game_titles (id),
    CONSTRAINT game_matches_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
);
create index on game_matches(game_id, user_id);