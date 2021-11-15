CREATE TABLE IF NOT EXISTS account
(
    id CHARACTER varying(36) NOT NULL,
    name CHARACTER varying(200) NOT NULL,
    cpf CHARACTER varying(11) NOT NULL,
    secret_login CHARACTER varying(300),
    balance NUMERIC DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id, cpf)
);

