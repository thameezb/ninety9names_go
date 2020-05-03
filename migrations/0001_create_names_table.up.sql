CREATE TABLE IF NOT EXISTS names(
   id                   INT      PRIMARY KEY,
   arabic               VARCHAR UNIQUE NOT NULL,
   translitiration      VARCHAR UNIQUE NOT NULL,
   meaningShaykh        VARCHAR NOT NULL,
   meaningGeneral       VARCHAR NOT NULL,
   explanation          VARCHAR NOT NULL
);