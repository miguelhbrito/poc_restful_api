CREATE TABLE IF NOT EXISTS login_user
(
    cpf CHARACTER varying(11) NOT NULL,
    secret CHARACTER varying(200),
    PRIMARY KEY (cpf)
);
