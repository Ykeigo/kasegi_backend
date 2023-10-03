CREATE TABLE IF NOT EXISTS users(
    id VARCHAR (64) PRIMARY KEY,
    email VARCHAR (256) NOT NULL
);

CREATE TABLE IF NOT EXISTS login_sessions(
    id VARCHAR (64) PRIMARY KEY,
    session_token VARCHAR (256) NOT NULL,
    user_id VARCHAR (64) NOT NULL,
    expires_at TIMESTAMP NOT NULL
);


CREATE TABLE IF NOT EXISTS game_titles(
    id VARCHAR (64) PRIMARY KEY,
    title VARCHAR (64) NOT NULL
);

CREATE TABLE IF NOT EXISTS game_matches(
    id VARCHAR (64) PRIMARY KEY,
    game_id VARCHAR (64) NOT NULL,
    user_id VARCHAR (64) NOT NULL,
    created_at TIMESTAMP NOT NULL
);
create index on game_matches(game_id, user_id);

CREATE TABLE IF NOT EXISTS game_match_checklist_items(
    id VARCHAR (64) PRIMARY KEY,
    match_id VARCHAR (64) NOT NULL,
    title VARCHAR (64) NOT NULL,
    is_checked BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS checklist_templates(
    id VARCHAR (64) PRIMARY KEY,
    template_name VARCHAR (64) NOT NULL,
    game_title_id VARCHAR (64) NOT NULL,
    created_by_user_id VARCHAR (64) NOT NULL
);

CREATE TABLE IF NOT EXISTS checklist_template_items(
    id VARCHAR (64) PRIMARY KEY,
    checklist_template_id VARCHAR (64) NOT NULL,
    title VARCHAR (64) NOT NULL,
    is_checked BOOLEAN NOT NULL
);