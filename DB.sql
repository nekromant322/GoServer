CREATE TABLE USERS(
    login           TEXT    NOT NULL   UNIQUE,
    password        TEXT    NOT NULL,
    rank            INTEGER NOT NULL
);
CREATE TABLE COURSES(
    courseID        INTEGER NOT NULL PRIMARY KEY,
    name            INTEGER NOT NULL,
    amount          INTEGER
);
CREATE TABLE MARKS(
    login           TEXT    NOT NULL,
    lesson_number   INTEGER NOT NULL,
    class_mark      INTEGER,
    home_mark       INTEGER,
    groupID         INTEGER NOT NULL
);
CREATE TABLE GROUPS(
    groupID         INTEGER NOT NULL PRIMARY KEY,
    group_name      TEXT NOT NULL,
    courseID        INTEGER NOT NULL,
    teacher         TEXT    

); 
CREATE TABLE LESSONS(
    courseID        INTEGER NOT NULL,
    lesson_number   INTEGER NOT NULL,
    theme           TEXT,
    homework        TEXT
);
CREATE INDEX groupsAndLogins ON MARKS(groupID, login);