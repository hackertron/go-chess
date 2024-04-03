-- Tables
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE user_info (
    user_id INTEGER PRIMARY KEY,
    user_rank TEXT CHECK (user_rank IN ('beginner', 'intermediate', 'advanced', 'expert')) DEFAULT 'beginner',
    user_points INTEGER DEFAULT 1000,
    KEY (user_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
