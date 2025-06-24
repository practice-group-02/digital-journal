CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE languages (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE universities (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    website TEXT
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE programs (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    deadline DATE,
    duration TEXT,
    degree TEXT,
    funding TEXT,
    url TEXT,
    country_id INT REFERENCES countries(id),
    language_id INT REFERENCES languages(id),
    university_id INT REFERENCES universities(id)
);

CREATE TABLE program_tags (
    program_id INT REFERENCES programs(id) ON DELETE CASCADE,
    tag_id INT REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (program_id, tag_id)
);

CREATE TABLE eligibility (
    id SERIAL PRIMARY KEY,
    program_id INT REFERENCES programs(id) ON DELETE CASCADE,
    citizenship TEXT,
    age_range TEXT,
    education_level TEXT,
    experience_required TEXT
);

CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    program_id INT REFERENCES programs(id) ON DELETE CASCADE,
    name TEXT NOT NULL
);
