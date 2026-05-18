CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    cpf VARCHAR(20) NOT NULL,
    idade INT NOT NULL,
    email VARCHAR(100) NOT NULL,
    ocorrencia TEXT,
    tipo_sanguineo VARCHAR(3),
    sintomas TEXT,
    diagnostico TEXT
);