-- +goose Up
CREATE TABLE kanji (
  id BIGSERIAL PRIMARY KEY,
  kanji TEXT,
  kanji_meaning TEXT,
  kanji_level INTEGER
);

CREATE TABLE kanji_alternatives (
  id BIGSERIAL PRIMARY KEY,
  kanji_alternative TEXT,
  kanji_id BIGINT NOT NULL
           REFERENCES kanji(id) 
           ON DELETE CASCADE 
           ON UPDATE NO ACTION
);

CREATE TABLE onyomi (
  id BIGSERIAL PRIMARY KEY,
  kanji_onyomi TEXT,
  kanji_id BIGINT NOT NULL
           REFERENCES kanji(id) 
           ON DELETE CASCADE 
           ON UPDATE NO ACTION
);

CREATE TABLE kunyoumi (
  id BIGSERIAL PRIMARY KEY,
  kanji_kunyoumi TEXT,
  kanji_id BIGINT NOT NULL
           REFERENCES kanji(id) 
           ON DELETE CASCADE 
           ON UPDATE NO ACTION
);

CREATE TABLE words (
  id BIGSERIAL PRIMARY KEY,
  word TEXT,
  word_meaning TEXT,
  word_level INTEGER
);

CREATE TABLE word_alternatives (
  id BIGSERIAL PRIMARY KEY,
  word_alternative TEXT,
  word_id BIGINT NOT NULL
          REFERENCES word(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION
);

CREATE TABLE word_readings (
  id BIGSERIAL PRIMARY KEY,
  word_reading TEXT,
  word_id BIGINT NOT NULL
          REFERENCES word(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION
);

CREATE TABLE word_types (
  id BIGSERIAL PRIMARY KEY,
  word_type TEXT,
  word_id BIGINT NOT NULL
          REFERENCES word(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION
);

CREATE TABLE sentences (
  id BIGSERIAL PRIMARY KEY,
  japanese_sentence TEXT,
  word_id BIGINT NOT NULL
          REFERENCES word(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION
);

CREATE TABLE sentence_translations (
  id BIGSERIAL PRIMARY KEY,
  sentence_id BIGINT NOT NULL 
              REFERENCES sentence(id) 
              ON DELETE CASCADE 
              ON UPDATE NO ACTION,
  sentence_language INTEGER,
  sentence_translation TEXT

  
);

CREATE TABLE compositions (
  id BIGSERIAL PRIMARY KEY,
  kanji_id BIGINT NOT NULL 
           REFERENCES kanji(id) 
           ON DELETE CASCADE 
           ON UPDATE NO ACTION,

  word_id BIGINT NOT NULL 
          REFERENCES word(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION
  
);

CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username TEXT,
  user_password TEXT,
  user_email TEXT,
  first_name TEXT,
  second_name TEXT
);

CREATE TABLE user_progress (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL 
          REFERENCES users(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION,
  user_level INTEGER
);

CREATE TABLE kanji_progress (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL 
          REFERENCES users(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION,

  kanji_id BIGINT NOT NULL 
           REFERENCES kanji(id) 
           ON DELETE CASCADE 
           ON UPDATE NO ACTION,

  srs_level INTEGER,
  unlock_date date,
  next_date date,
  burn_date date
);

CREATE TABLE word_progress (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL 
          REFERENCES users(id) 
          ON DELETE CASCADE 
          ON UPDATE NO ACTION,

  kanji_id BIGINT NOT NULL 
           REFERENCES kanji(id) 
           ON DELETE CASCADE 
           ON UPDATE NO ACTION,
  srs_level INTEGER,
  unlock_date date,
  next_date date,
  burn_date date
);

-- +goose Down
DROP TABLE kanji;
DROP TABLE words;
DROP TABLE users;
