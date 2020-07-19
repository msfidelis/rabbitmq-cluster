#!/bin/bash

set -e

/usr/local/bin/docker-entrypoint.sh rabbitmq-server -detached

sleep 5s

rabbitmqctl stop_app
rabbitmqctl join_cluster rabbit@rabbitmq01

rabbitmqctl stop

sleep 2s

rabbitmq-server