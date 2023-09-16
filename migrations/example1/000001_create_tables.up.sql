CREATE TABLE IF NOT EXISTS users(
    id VARCHAR (64) PRIMARY KEY,
    google_user_id VARCHAR (256) NOT NULL
);
create index on users(id);

CREATE TABLE IF NOT EXISTS login_sessions(
    id VARCHAR (64) PRIMARY KEY,
    session_token VARCHAR (256) NOT NULL,
    user_id VARCHAR (64) NOT NULL,
    expires_at TIMESTAMP NOT NULL
);
create index on login_sessions(id);


CREATE TABLE IF NOT EXISTS game_titles(
    id serial PRIMARY KEY,
    title VARCHAR (64) NOT NULL
);
create index on game_titles(id);

CREATE TABLE IF NOT EXISTS game_matches(
    id serial PRIMARY KEY,
    game_id int NOT NULL,
    user_id VARCHAR (64) NOT NULL
);
create index on game_matches(game_id, user_id);