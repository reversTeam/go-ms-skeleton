CREATE KEYSPACE IF NOT EXISTS global
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};

CREATE KEYSPACE IF NOT EXISTS auth
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};

USE global;

CREATE TABLE IF NOT EXISTS people (
    id uuid,
    created_at timestamp,
    updated_at timestamp,
    status varchar,
    expired_at timestamp,
    firstname varchar,
    lastname varchar,
    birthday date,
    email_id uuid,
    PRIMARY KEY (id)
);

CREATE INDEX ON people (status);

CREATE TABLE IF NOT EXISTS email (
    people_id uuid,
    created_at timestamp,
    updated_at timestamp,
    status varchar,
    validated_at timestamp,
    expired_at timestamp,
    email varchar,
    PRIMARY KEY (people_id, email)
);

CREATE INDEX ON email (status);

USE auth;

CREATE TABLE IF NOT EXISTS account (
    people_id uuid
    created_at timestamp,
    updated_at timestamp,
    status varchar,
    validated_at timestamp,
    expired_at timestamp,
    password varchar,
    email_id uuid,
    signin_id uuid,
    PRIMARY KEY (people_id)
);

CREATE INDEX ON account (email_id);
CREATE INDEX ON account (signin_id);
CREATE INDEX ON account (status);

CREATE TABLE IF NOT EXISTS connexion (
    account_id uuid,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    status VARCHAR,
    expired_at TIMESTAMP,
    session_token VARCHAR,
    refresh_token VARCHAR,
    PRIMARY KEY (account_id)
);
CREATE INDEX ON connexion (status);
CREATE INDEX ON connexion (session_token);
CREATE INDEX ON connexion (refresh_token);

CREATE TABLE IF NOT EXISTS password_reset_request (
    account_id uuid,
    created_at timestamp,
    updated_at timestamp,
    status varchar,
    expired_at timestamp,
    reset_token varchar,
    email_id uuid,
    PRIMARY KEY (account_id, email_id)
);
CREATE INDEX ON password_reset_request (status);

CREATE TABLE IF NOT EXISTS auth.signin (
    id UUID,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    status TEXT,
    expired_at TIMESTAMP,
    validated_at TIMESTAMP,
    email TEXT,
    firstname TEXT,
    lastname TEXT,
    birthday DATE,
    password TEXT,
    validation_token TEXT,
    account_id uuid,
    email_id uuid,
    PRIMARY KEY (id)
);

CREATE INDEX ON signin (status);
CREATE INDEX ON signin (email);
