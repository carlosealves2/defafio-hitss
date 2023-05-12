CREATE TABLE IF NOT EXISTS users
(
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    contact TEXT NOT NULL,
    address TEXT NOT NULL,
    birth DATE NOT NULL,
    cpf TEXT NOT NULL
);
