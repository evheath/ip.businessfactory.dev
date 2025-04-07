#!/bin/bash
git pull
go build -o main cmd/web/*.go
sudo supervisorctl stop ip.businessfactory.dev
sudo supervisorctl start ip.businessfactory.dev
