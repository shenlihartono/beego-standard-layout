create table if not exists struct
(
	id serial not null
		constraint struct_pk
			primary key,
	struct_id varchar not null,
	value integer default 0
);