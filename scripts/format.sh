#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "******************"
echo "Formatting $DIR/../swagger"
cd $DIR/../swagger
go fmt