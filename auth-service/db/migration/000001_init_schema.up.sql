CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  username varchar(50) UNIQUE NOT NULL,
  email varchar(100) UNIQUE NOT NULL,
  password_hash varchar(100) NOT NULL,
  role varchar(50) NOT NULL DEFAULT 'user',
  provider VARCHAR(50) DEFAULT 'local' NOT NULL,
  verified boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  password_changed_at timestamptz NOT NULL DEFAULT now(),
  last_login timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE sessions (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  username varchar(50) NOT NULL REFERENCES users(username),
  refresh_token varchar NOT NULL,
  user_agent varchar NOT NULL, 
  client_ip varchar NOT NULL,
  is_blocked boolean NOT NULL DEFAULT false,
  expires_at timestamptz NOT NULL ,
  created_at timestamptz NOT NULL DEFAULT now(),
  last_login timestamptz NOT NULL DEFAULT now()
);