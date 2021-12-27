\connect my_company;

CREATE SCHEMA hr;

CREATE TABLE hr.employees (
    pnumber uuid NOT NULL,
    name varchar(50) NOT NULL,
    address varchar(100) NOT NULL,
    email varchar(50),
    birth date,
    department varchar(25),
    job_title varchar(50) NOT NULL,
    PRIMARY KEY (pnumber)
);

CREATE USER technical_user with PASSWORD '!H&n#T75FogkxFB0';
GRANT USAGE ON SCHEMA hr TO technical_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA "hr" to technical_user;
