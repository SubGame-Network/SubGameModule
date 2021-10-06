#!/usr/bin/env bash

set -e

cp config/config.prod.yml config/config.yml

go run cmd/migrate/gorm/migration/main.go
go run cmd/migrate/gorm/seed/main.go

if [ ! -f ./main ]; then
    go build cmd/gin/main.go
fi

./main