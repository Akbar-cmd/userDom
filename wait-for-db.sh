#!/bin/sh
set -e

host="$1"
user="$2"
password="$3"
dbname="$4"
shift 4
cmd="$@"

export PGPASSWORD="$password"

until psql -h "$host" -U "$user" -d "$dbname" -c '\q' > /dev/null 2>&1; do
  >&2 echo "PostgreSQL is unavailable - sleeping"
  sleep 2
done

>&2 echo "PostgreSQL is up - executing command"
exec "$@"