CREATE TABLE IF NOT EXISTS names(
   id                   VARCHAR       PRIMARY KEY,
   arabic               VARCHAR UNIQUE NOT NULL,
   transliteration      VARCHAR UNIQUE NOT NULL,
   meaning_shaykh       VARCHAR NOT NULL,
   meaning_general      VARCHAR NOT NULL,
   explanation          VARCHAR NOT NULL
);