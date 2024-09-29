#!/bin/bash

# This script runs when the MySQL container is started
if [ "$RESET_DB" = "true" ]; then
  echo "Resetting the database..."

  # Use environment variables for root user and password
  MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD:-rootpassword}

  # Drop and recreate the database
  # mysql -uroot -p"$MYSQL_ROOT_PASSWORD" -e "DROP DATABASE IF EXISTS userdb; CREATE DATABASE userdb;"
  
  # Add schema creation or data seeding here
  mysql -uroot -p"$MYSQL_ROOT_PASSWORD" user_db < /docker-entrypoint-initdb.d/schema.sql
  
  echo "Database reset and recreated."
else
  echo "Database reset not required."
fi
