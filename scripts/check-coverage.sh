#!/bin/bash

set -e

# Configuration
GO_MIN_COVERAGE=95
VUE_MIN_COVERAGE=80

echo "=== Running Go Tests with Coverage ==="
go test -coverprofile=coverage.out ./...

echo ""
echo "=== Go Coverage Report ==="
go tool cover -func=coverage.out

# Calculate total coverage percentage
TOTAL_COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

echo ""
echo "Total Go Coverage: ${TOTAL_COVERAGE}%"
echo "Minimum Required: ${GO_MIN_COVERAGE}%"

# Check if coverage meets minimum threshold
if (( $(echo "$TOTAL_COVERAGE < $GO_MIN_COVERAGE" | bc -l) )); then
    echo "❌ FAILED: Coverage ${TOTAL_COVERAGE}% is below minimum ${GO_MIN_COVERAGE}%"
    exit 1
else
    echo "✅ PASSED: Coverage ${TOTAL_COVERAGE}% meets minimum ${GO_MIN_COVERAGE}%"
fi

# TODO: Add Vue/Vitest coverage check when frontend tests are set up
# echo ""
# echo "=== Running Vue Tests with Coverage ==="
# cd web && npm test -- --coverage

exit 0
