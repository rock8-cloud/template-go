CREATE TABLE space_target (
    id                   BIGSERIAL PRIMARY KEY,
    name                 VARCHAR(255)     NOT NULL,
    type                 VARCHAR(100)     NOT NULL,
    distance_light_years DOUBLE PRECISION NOT NULL
);

INSERT INTO space_target (name, type, distance_light_years) VALUES
    ('Betelgeuse',       'Red Supergiant',  700),
    ('NGC 253',          'Spiral Galaxy',   11400000),
    ('Eagle Nebula',     'Emission Nebula', 7000);
