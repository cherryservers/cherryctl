#!/usr/bin/env bash
go test ./internal/plans -run '^$' -fuzz '^FuzzPlanListResponse$' -fuzztime=1m

echo "Fuzzing complete"
