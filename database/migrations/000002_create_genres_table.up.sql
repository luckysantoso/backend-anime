CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO genres (name) 
VALUES
    ('Action'),
    ('Adventure'),
    ('Comedy'),
    ('Drama'),
    ('Fantasy'),
    ('Magic'),
    ('Mecha'),
    ('Music'),
    ('Romance'),
    ('Sci-Fi'),
    ('Shounen'),
    ('Slice of Life'),
    ('Sports'),
    ('Supernatural');
