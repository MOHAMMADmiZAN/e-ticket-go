#!/bin/sh

# Check if the ENV environment variable is set to 'dev'
if [ "$ENV" = "dev" ]; then
  # Run air for hot reloading in development
  exec air
else
  # Run the compiled application in production
  exec ./auth-service
fi
