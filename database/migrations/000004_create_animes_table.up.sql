CREATE TABLE animes (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre_id INT REFERENCES genres(id) ON DELETE CASCADE ON UPDATE CASCADE,
    sinopsis TEXT NOT NULL,
    episodes INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO animes (title, genre_id, sinopsis, episodes)
VALUES
    ('Attack on Titan', 1, 'Humanity fights for survival against man-eating giants.', 87),
    ('One Piece', 2, 'Follow the adventures of Luffy and his pirate crew in search of the ultimate treasure.', 1050),
    ('Naruto', 3, 'A young ninja strives to become the Hokage while protecting his village.', 720),
    ('Demon Slayer', 4, 'A young boy becomes a demon slayer to avenge his family and save his sister.', 26),
    ('My Hero Academia', 1, 'In a world where superpowers are the norm, a boy without powers aspires to be a hero.', 138),
    ('Your Lie in April', 12, 'A young pianist rediscovers his love for music through a spirited violinist.', 22);


-- CREATE TABLE animes (
--     id AUTO_INCREMENT PRIMARY KEY,
--     title VARCHAR(255) NOT NULL,
--     genre VARCHAR(255) NOT NULL,
--     review TEXT NOT NULL,
--     episodes INT NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );
