#!/bin/bash
# Attend que ScyllaDB soit prête
until cqlsh -e "describe keyspaces"; do
  echo "ScyllaDB is unavailable - sleeping"
  sleep 1
done

# Exécute les commandes CQL depuis le fichier
echo "Initializing ScyllaDB schema..."
cqlsh -f /docker-entrypoint-initdb.d/init-db.sql

echo "Schema initialized. Starting ScyllaDB..."
