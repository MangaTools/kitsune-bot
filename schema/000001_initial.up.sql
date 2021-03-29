CREATE TABLE manga
(
    id     serial not null unique,
    name   varchar(512) unique,
    status int2,
    PRIMARY KEY (id)
);

CREATE TABLE chapter
(
    id       serial,
    manga_id int REFERENCES manga (id) ON DELETE CASCADE,
    number   float,
    pages    int2,
    status   int2,
    PRIMARY KEY (id)
);

CREATE TABLE "user"
(
    id               varchar(64) unique,
    username         varchar(255),
    score            int,
    translated_pages int,
    edited_pages     int,
    cleaned_pages    int,
    typed_chapters   int,
    PRIMARY KEY (id)
);

CREATE TABLE owner
(
    id         serial,
    user_id    varchar(64) REFERENCES "user" (id) ON DELETE CASCADE,
    chapter_id int references chapter (id) On DELETE CASCADE,
    page_start int2,
    page_end   int2,
    status     int2,
    work_type  int2,
    PRIMARY KEY (id)
);





