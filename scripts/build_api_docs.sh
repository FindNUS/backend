#!/bin/bash
if og ../api/findnus.yaml -o ../api/ markdown; then
    mv ../api/openapi.md ../api/README.md
else
    echo "Error occured."
fi