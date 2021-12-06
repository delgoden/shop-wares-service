CREATE TABLE categories
(
    id serial NOT NULL unique,
    title varchar(255) NOT NULL
);

CREATE TABLE categories_wares
(
    id serial NOT NULL unique,
    category_id int references categories (id) on delete cascade not null,
    ware_id int references wares (id) on delete cascade not null
);

CREATE TABLE wares
(
    id serial NOT NULL unique,
    title varchar(255) NOT NULL,
    description text,
    active BOOLEAN NOT NULL DEFAULT true,
	created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);