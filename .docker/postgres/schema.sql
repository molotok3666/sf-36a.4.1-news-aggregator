DROP TABLE IF EXISTS news;

CREATE TABLE news (
    guid text NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    pub_time BIGINT NOT NULL,
    link TEXT NOT NULL
);

CREATE UNIQUE INDEX news_guid_uindex
    on news (guid);

