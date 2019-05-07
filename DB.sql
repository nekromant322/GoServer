CREATE TABLE USERS(
    login           TEXT    NOT NULL   UNIQUE,
    password        TEXT    NOT NULL,
    rank            INTEGER NOT NULL,
    real_name       TEXT    NOT NULL,
    birthday        TEXT DEFAULT '_' NOT NULL,
    bonus_info      TEXT DEFAULT '_' NOT NULL
);
CREATE TABLE COURSES(
    courseID        INTEGER NOT NULL PRIMARY KEY,
    name            INTEGER NOT NULL,
    amount          INTEGER NOT NULL
);
CREATE TABLE MARKS(
    login           TEXT    NOT NULL,
    lesson_number   INTEGER NOT NULL,
    class_mark      INTEGER DEFAULT 0 NOT NULL,
    home_mark       INTEGER DEFAULT 0 NOT NULL,
    groupID         INTEGER NOT NULL
);
CREATE TABLE GROUPS(
    groupID         INTEGER NOT NULL PRIMARY KEY,
    group_name      TEXT NOT NULL,
    courseID        INTEGER NOT NULL,
    teacher         TEXT DEFAULT '_' NOT NULL,    
    info            TEXT DEFAULT '_' NOT NULL
); 
CREATE TABLE LESSONS(
    courseID        INTEGER NOT NULL,
    lesson_number   INTEGER NOT NULL,
    theme           TEXT DEFAULT '_' NOT NULL,
    homework        TEXT DEFAULT '_' NOT NULL
);
CREATE TABLE EVENTS(
    groupID         INTEGER NOT NULL,
    date            TEXT NOT NULL,
    event           TEXT NOT NULL
);
CREATE INDEX groupsAndLogins ON MARKS(groupID, login);