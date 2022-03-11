#!/bin/bash
go run main.go
cd ./go-diagrams/
dot -Tpng workers.dot > ../diagram.png
rm -rf ./go-diagrams/
