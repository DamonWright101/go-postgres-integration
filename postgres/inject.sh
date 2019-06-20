#! /bin/bash
source postgres.env
export  PGPASSWORD=${POSTGRES_PASSWORD}

psql ${POSTGRES_DB} ${POSTGRES_USER} -h localhost -p 5432 -a -f ./schema.sql 2>/dev/null
psql ${POSTGRES_DB} ${POSTGRES_USER} -h localhost -p 5432 -a -f ./data.sql 2>/dev/null

echo TABLES:
psql ${POSTGRES_DB} ${POSTGRES_USER} -h localhost -p 5432 -c '\dt;'
echo; echo;
psql ${POSTGRES_DB} ${POSTGRES_USER} -h localhost -p 5432 -c 'SELECT * FROM persons;'