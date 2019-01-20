#!/bin/bash
psql -U root -d testdb << "EOSQL"
create table employee (emp_id int PRIMARY KEY, emp_number int);
insert into employee (emp_id, emp_number) values (1, 2222);
EOSQL