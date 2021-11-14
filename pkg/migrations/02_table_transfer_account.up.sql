CREATE TABLE IF NOT EXISTS transfer_account
(
    id CHARACTER varying(36) NOT NULL,
    account_origin_id CHARACTER varying(200) NOT NULL,
    account_destination_id CHARACTER varying(200) NOT NULL,
    amount NUMERIC DEFAULT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);
