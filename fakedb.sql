create table Client{
    id SERIAL PRIMARY KEY,
    client_surename varchar(100) not null,
    client_name varchar(100) not null,
    client_secondname varchar(100) not null,
    client_phone varchar(100) not null
}

create table Application{
    id SERIAL PRIMARY KEY,
    id_client int not null,
    id_teacher int not null,
    application_date date not null,
    application_time time not null,
    FOREIGN KEY (id_client) REFERENCES Client (id),
    FOREIGN KEY (id_teacher) REFERENCES Teacher (id)
}

create table Teacher{
    id SERIAL PRIMARY KEY
    teacher_surename varchar(100) not null,
    teacher_name varchar(100) not null,
    teacher_post varchar(100) not null,
    id_schedule int not null,
    FOREIGN KEY (id_schedule) REFERENCES Schedule (id)
}

create table Schedule{
    id SERIAL PRIMARY KEY,
    start_time time not null,
    end_time time not null,
    workday varchar(100) not null
}

create table GroupEvent{
    id SERIAL PRIMARY KEY,
    id_theme int not null,
    id_teacher int not null,
    class varchar(10) not null,
    FOREIGN KEY (id_theme) REFERENCES Theme (id),
    FOREIGN KEY (id_teacher) REFERENCES Teacher (id)
}

create table Meeting{
    id SERIAL PRIMARY KEY,
    id_theme int not null,
    id_application int not null,
    meeting_date date not null,
    meeting_time time not null,
    FOREIGN KEY (id_theme) REFERENCES Theme (id),
    FOREIGN KEY (id_application) REFERENCES Application (id)
}

create table Theme{
    id SERIAL PRIMARY KEY,
    Theme varchar(100) not null,
    id_category int not null,
    FOREIGN KEY (id_category) REFERENCES Category (id)
}
create table Category{
    id SERIAL PRIMARY KEY,
    category varchar(100) not null
}