#!/bin/bash
echo "Autogenerating API documentation..."
if openapi-markdown -i ../api/findnus.yaml -o ../api/README.md; then
    echo "Docs generated!"
else
    echo "Error! Did you npm install -g openapi-markdown?"
fi