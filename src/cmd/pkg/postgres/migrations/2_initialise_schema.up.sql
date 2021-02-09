CREATE TABLE IF NOT EXISTS user
(
    creation_timestamp timestamp      NOT NULL,
    update_timestamp   timestamp,
    id                 uuid           NOT NULL PRIMARY KEY,
    name               VARCHAR        NOT NULL,
    isPremium            boolean        NOT NULL DEFAULT FALSE
);

CREATE TRIGGER user_creation_ts_trigger
    BEFORE INSERT
    ON user
    FOR EACH ROW
EXECUTE
    PROCEDURE set_creation_timestamp();

CREATE TRIGGER user_update_ts_trigger
    BEFORE UPDATE
    ON user
    FOR EACH ROW
EXECUTE
    PROCEDURE set_update_timestamp();
