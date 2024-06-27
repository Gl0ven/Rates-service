-- +goose Up
CREATE TABLE IF NOT EXISTS rates (
    unix_timestamp INT,
    ask_price NUMERIC,
    bid_price NUMERIC
);

-- +goose Down
DROP TABLE IF EXISTS rates;