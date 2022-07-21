#!/usr/bin/bash
go build -o /usr/local/tasks && \
  supervisorctl restart tasks