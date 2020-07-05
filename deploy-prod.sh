#!/usr/bin/env bash
# Deploys the production build of Rankathon
# Usage: deploy-prod.sh

docker-compose -f docker-compose.prod.yml up --build