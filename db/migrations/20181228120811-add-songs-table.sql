
-- +migrate Up
CREATE TABLE songs (
    id              INT     PRIMARY KEY NOT NULL,
    title           TEXT    NOT NULL,
    artist          TEXT    NOT NULL,
    year            INT     NOT NULL,
    video           TEXT    NOT NULL,
    duration_label  TEXT    NOT NULL,
    duration        INT     NOT NULL
);

-- +migrate Down
DROP TABLE songs;
