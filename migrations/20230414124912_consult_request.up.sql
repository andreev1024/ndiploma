CREATE TABLE consult_requests
(
    name varchar(100) not null,
    phone varchar(50) not null,
    role varchar(50) not null,
    available_time varchar(100),
    consult_date varchar(100),
    created_at timestamp default current_timestamp,
    constraint consult_requests_pk primary key (phone, created_at)
);
