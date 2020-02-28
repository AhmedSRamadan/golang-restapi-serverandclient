#!/bin/bash
cd rest-server/ ; docker build -t server . ; kubectl apply -f deployment.yml ; kubectl apply -f  node-port-service.yml ; kubectl apply -f clusterip-service.yml ; cd ../rest-client/ ; docker build -t client . ; kubectl apply -f deployment.yml ; kubectl apply -f  node-port-service.yml 
