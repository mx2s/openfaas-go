#!/bin/bash
go build function/handler.go
mkdir -p function/dist
mv handler function/dist/handler
