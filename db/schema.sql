CREATE DATABASE bilDev; 

CREATE TABLE "users" (
    "id" INTEGER    PRIMARY KEY,
    "username" VARCHAR (100),
    "email"  VARCHAR (100),
    "created_at" TIMESTAMP
);


CREATE TABLE "posts" (
    "id" SERIAL    PRIMARY KEY ,
    "title" VARCHAR (100) NOT NULL,
    "body"  VARCHAR (556)  NOT NULL,
    "created_at" TIMESTAMP
);

CREATE UNIQUE INDEX post_id ON posts (id);