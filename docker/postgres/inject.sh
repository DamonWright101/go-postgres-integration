#! /bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source ${DIR}/postgres.env
export  PGPASSWORD=${POSTGRES_PASSWORD}

psql ${POSTGRES_DB} ${POSTGRES_USER} -h localhost -p 5432 -a -f ${DIR}//schema.sql &>/dev/null
psql ${POSTGRES_DB} ${POSTGRES_USER} -h localhost -p 5432 -a -f ${DIR}//data.sql &>/dev/null