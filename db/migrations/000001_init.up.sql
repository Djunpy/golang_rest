CREATE TYPE user_role AS ENUM ('is_active', 'is_staff', 'is_superuser');
CREATE TYPE post_category AS ENUM ('Python', 'Golang', 'JavaScript');

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    photo VARCHAR NOT NULL DEFAULT '',
    verified BOOL DEFAULT false,
    password VARCHAR NOT NULL,
    role user_role NOT NULL DEFAULT 'is_active',
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL UNIQUE,
    category post_category DEFAULT 'Python',
    content TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);