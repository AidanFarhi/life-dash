#!/usr/bin/env bash

# delete the db file if it exists
[ -f "db/lifedash.db" ] && rm "db/lifedash.db"

# create the db and run the create table statements
sqlite3 "db/lifedash.db" < "sql/create_tables.sql"

echo "done"