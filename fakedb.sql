create table Client{
    id_client SERIAL PRIMARY KEY
    client_surename varchar(100) not null
    client_name varchar(100) not null
    client_secondname varchar(100) not null
    client_phone 
}

create table Application{
    id_application SERIAL PRIMARY KEY
    id_client int not null
    id_teacher int not null
    application_date date not null
    application_time time
    FOREIGN KEY (id_client) REFERENCES Application (id_client)
    FOREIGN KEY (id_teacher) REFERENCES Teacher (id_teacher)
}

create table Teacher{
    id_teacher SERIAL PRIMARY KEY
    teacher_surename varchar(100) not null
    teacher_name varchar(100) not null
    teacher_post varchar(100) not null
    id_schedule int not null
    FOREIGN KEY (id_schedule) REFERENCES Schedule (id_schedule)
}

create table Schedule{
    id_schedule SERIAL PRIMARY KEY
    start_time time
    end_time time
    workday
}

create table GroupEvent{
    id_group_event SERIAL PRIMARY KEY
    id_theme int not null
    id_teacher int not null
    class varchar(10) not null
    FOREIGN KEY (id_theme) REFERENCES Theme (id_theme)
    FOREIGN KEY (id_teacher) REFERENCES Teacher (id_teacher)
}

create table Meeting{
    id_meeting SERIAL PRIMARY KEY
    id_theme int not null
    id_application int not null
    meeting_date date
    meeting_time time
    FOREIGN KEY (id_theme) REFERENCES Theme (id_theme)
    FOREIGN KEY (id_application) REFERENCES Application (id_application)
}

create table Theme{
    id_theme SERIAL PRIMARY KEY
    Theme varchar(100) not null
    id_category int not null
    FOREIGN KEY (id_category) REFERENCES Category (id_category)
}
create table Category{
    id_category SERIAL PRIMARY KEY
    category varchar(100) not null
}