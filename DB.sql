CREATE TABLE USERS(
    login           TEXT    NOT NULL   UNIQUE,
    password        TEXT    NOT NULL,
    rank            INTEGER NOT NULL,
    real_name       TEXT    NOT NULL,
    birthday        TEXT,
    bonus_info      TEXT
);
CREATE TABLE COURSES(
    courseID        INTEGER NOT NULL PRIMARY KEY,
    name            TEXT    NOT NULL,
    amount          INTEGER,
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
    teacher         TEXT,    
    info            TEXT
); 
CREATE TABLE LESSONS(
    courseID        INTEGER NOT NULL,
    lesson_number   INTEGER NOT NULL,
    theme           TEXT,
    homework        TEXT
);
CREATE TABLE EVENTS(
    eventID         INTEGER NOT NULL PRIMARY KEY,
    groupID         INTEGER NOT NULL,
    date            TEXT NOT NULL,
    event           TEXT NOT NULL
);
CREATE INDEX groupsAndLogins ON MARKS(groupID, login);