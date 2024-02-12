CREATE KEYSPACE IF NOT EXISTS glb WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};

CREATE TABLE IF NOT EXISTS glb.people (
    id uuid PRIMARY KEY,
    firstname varchar(50),
    lastname varchar(50),
    email varchar(120),
    birthday date
);