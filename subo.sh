#!/bin/bash

# Compilar el archivo Go
#go build main.go
go build -o bootstrap main.go

# Eliminar el archivo main.zip si existe
rm -f main.zip

# Crear un nuevo archivo main.zip
zip bootstrap.zip bootstrap

# Limpiar cualquier ejecutable generado
#rm -f main
