CREATE TABLE users
(
    id serial not null UNIQUE,
    name varchar(255) not null,
    login varchar(255) not null UNIQUE,
    password varchar(255) not null
);
CREATE TABLE todo_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);
CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int REFERENCES users (id) on delete cascade not null,
    list_id int REFERENCES todo_lists (id) on delete CASCADE not null
);
CREATE TABLE todo_items
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    done boolean not null DEFAULT false
);
CREATE TABLE lists_items
(
    id serial not null unique,
    item_id int REFERENCES todo_items (id) on DELETE CASCADE not null,
    list_id int REFERENCES todo_lists (id) on DELETE CASCADE not null 
);