CREATE TABLE IF NOT EXISTS users
(
    creation_timestamp timestamp      NOT NULL,
    update_timestamp   timestamp,
    id                 uuid           NOT NULL PRIMARY KEY,
    name               VARCHAR        NOT NULL,
    isPremium            boolean        NOT NULL DEFAULT FALSE
);