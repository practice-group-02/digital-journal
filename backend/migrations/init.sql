CREATE TABLE program_types (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE programs (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    type INT,
    country TEXT,
    organization TEXT,
    deadline DATE,
    link TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    user_id INT,
    FOREIGN KEY (type) REFERENCES program_types(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE programs_tags (
    program_id INT,
    tag_id INT,
    PRIMARY KEY (program_id, tag_id),
    FOREIGN KEY (program_id) REFERENCES programs(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);