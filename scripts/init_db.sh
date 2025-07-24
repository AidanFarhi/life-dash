#!/usr/bin/env bash

# delete the db file if it exists
[ -f "db/lifedash.db" ] && rm "db/lifedash.db"

# create the db and run the create table statements
echo "creating tables"
sqlite3 "db/lifedash.db" < "sql/create_tables.sql"

# seed the db
echo "seeding db"
sqlite3 "db/lifedash.db" < "sql/seed_db.sql"

echo "done"