#!/bin/bash

# Démarrage de ScyllaDB
/usr/bin/scylla --developer-mode 1 --options-file /config.yml &

# Exécute le script d'initialisation
/docker-entrypoint-initdb.d/init-script.sh

tail -f /dev/null