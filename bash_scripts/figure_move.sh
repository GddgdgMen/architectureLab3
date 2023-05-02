#!/bin/bash

curl -X POST http://localhost:17000 -d "figure 0.5 0.5"
curl -X POST http://localhost:17000 -d "figure 0.2 0.9"
curl -X POST http://localhost:17000 -d "figure 0.3 0.7"
curl -X POST http://localhost:17000 -d "move 0.3 0.3"
curl -X POST http://localhost:17000 -d "update"