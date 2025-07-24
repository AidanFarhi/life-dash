
DROP TABLE IF EXISTS user;

CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

DROP TABLE IF EXISTS session;

CREATE TABLE session (
    id TEXT NOT NULL UNIQUE,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

DROP TABLE IF EXISTS expense;

CREATE TABLE expense (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    date TEXT NOT NULL,
    category TEXT NOT NULL,
    amount REAL NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);