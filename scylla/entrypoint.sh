#!/bin/bash
./wait_for_it.sh -t 0 scylla:9042 -- ./migrate.sh &
./docker-entrypoint.py