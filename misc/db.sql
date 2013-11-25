create database beatstorm;
create table beatvectors (uri varchar(25) not null, data text not null);
create unique index uri_index on beatvectors (uri);
