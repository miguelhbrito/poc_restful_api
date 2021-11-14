CREATE TABLE IF NOT EXISTS account
(
    id CHARACTER varying(36) NOT NULL,
    name CHARACTER varying(200) NOT NULL,
    cpf CHARACTER varying(200) NOT NULL,
    secret CHARACTER varying(100),
    balance NUMERIC DEFAULT 0,
    created_at TIMESTAMP NOT_NULL,
    PRIMARY KEY (id)
);
