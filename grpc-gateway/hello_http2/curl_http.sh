#!/bin/bash

curl -X POST -k https://localhost:8000/example/echo -d '{"name": "gRPC-HTTP is working!"}'
