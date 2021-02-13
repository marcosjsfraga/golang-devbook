
CREATE TABLE usuarios (
    id integer NOT NULL DEFAULT nextval('usuarios_id_seq'::regclass),
	nome varchar(50) NOT NULL,
    nick varchar(50) NOT NULL,
    email varchar(50),
    senha varchar(20) NOT NULL,
	criado_em timestamptz
);
