create table category (
    id int primary key auto_increment,
    name varchar(200) not null
);

insert into category (name) 
values("ABC"), ("DEF"), ("GHI");