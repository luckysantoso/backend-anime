CREATE TABLE predictions (
    id SERIAL PRIMARY KEY,
    image_path VARCHAR(255),
    label VARCHAR(255),
    score FLOAT
);