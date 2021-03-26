CREATE TABLE manga
(
    id     serial not null unique,
    name   varchar(512) unique,
    status varchar(128),
    PRIMARY KEY (id)
);

CREATE TABLE chapter
(
    id       serial,
    manga_id int REFERENCES manga (id) ON DELETE CASCADE,
    number   float,
    pages    int,
    PRIMARY KEY (id)
);

CREATE TABLE "user"
(
    id               varchar(64) unique,
    score            int,
    translated_pages int,
    edited_pages     int,
    checked_pages    int,
    cleaned_pages    int,
    typed_chapters   int,
    PRIMARY KEY (id)
);

CREATE TABLE owner
(
    id         serial,
    user_id    varchar(64) REFERENCES "user" (id) ON DELETE CASCADE,
    chapter_id int references chapter (id) On DELETE CASCADE,
    page_start int,
    page_end   int,
    status     varchar(64),
    work_type  varchar(64),
    PRIMARY KEY (id)
);

