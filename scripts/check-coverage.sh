#!/bin/bash

# Coverage enforcement script for Mahler platform
# Checks Go coverage (95% minimum) and Vue coverage (80% minimum)

set -e

# Get the script directory and project root
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# Change to project root
cd "$PROJECT_ROOT"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Thresholds
GO_THRESHOLD=95.0
VUE_THRESHOLD=80.0

echo "=========================================="
echo "Checking Test Coverage"
echo "=========================================="
echo ""

# Track overall success
overall_success=true

# ========================================
# Go Coverage Check
# ========================================
echo "📊 Checking Go coverage..."
echo ""

# Run Go tests with coverage (excluding cmd and testutil)
go test ./internal/app ./internal/api/... ./internal/project ./internal/service ./internal/resource/... ./internal -cover -coverprofile=coverage.out > /dev/null 2>&1

# Extract coverage percentage
GO_COVERAGE=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

echo "Go Coverage: ${GO_COVERAGE}%"
echo "Threshold:   ${GO_THRESHOLD}%"

# Compare coverage
if (( $(echo "$GO_COVERAGE >= $GO_THRESHOLD" | bc -l) )); then
    echo -e "${GREEN}✓ Go coverage meets requirements${NC}"
else
    echo -e "${RED}✗ Go coverage below threshold${NC}"
    echo -e "${RED}  Required: ${GO_THRESHOLD}%, Actual: ${GO_COVERAGE}%${NC}"
    overall_success=false
fi

echo ""

# ========================================
# Vue Coverage Check
# ========================================
echo "📊 Checking Vue coverage..."
echo ""

# Change to web directory and run tests
cd web

# Run Vue tests with coverage
npm test -- --run --coverage > /tmp/vue-coverage.txt 2>&1 || true

# Extract coverage percentages
VUE_STMTS=$(grep "All files" /tmp/vue-coverage.txt | awk -F'|' '{print $2}' | tr -d ' ' | head -1)
VUE_BRANCH=$(grep "All files" /tmp/vue-coverage.txt | awk -F'|' '{print $3}' | tr -d ' ' | head -1)
VUE_FUNCS=$(grep "All files" /tmp/vue-coverage.txt | awk -F'|' '{print $4}' | tr -d ' ' | head -1)
VUE_LINES=$(grep "All files" /tmp/vue-coverage.txt | awk -F'|' '{print $5}' | tr -d ' ' | head -1)

# Use statements coverage as the main metric
VUE_COVERAGE=$VUE_STMTS

echo "Vue Coverage (Statements): ${VUE_COVERAGE}%"
echo "Vue Coverage (Branches):   ${VUE_BRANCH}%"
echo "Vue Coverage (Functions):  ${VUE_FUNCS}%"
echo "Vue Coverage (Lines):      ${VUE_LINES}%"
echo "Threshold:                 ${VUE_THRESHOLD}%"

# Compare coverage (remove % symbol if present)
VUE_COVERAGE_NUM=$(echo $VUE_COVERAGE | sed 's/%//')

if (( $(echo "$VUE_COVERAGE_NUM >= $VUE_THRESHOLD" | bc -l) )); then
    echo -e "${GREEN}✓ Vue coverage meets requirements${NC}"
else
    echo -e "${RED}✗ Vue coverage below threshold${NC}"
    echo -e "${RED}  Required: ${VUE_THRESHOLD}%, Actual: ${VUE_COVERAGE}%${NC}"
    overall_success=false
fi

# Return to root
cd ..

echo ""
echo "=========================================="

# Final result
if [ "$overall_success" = true ]; then
    echo -e "${GREEN}✓ All coverage checks passed!${NC}"
    echo "=========================================="
    exit 0
else
    echo -e "${RED}✗ Coverage checks failed${NC}"
    echo "=========================================="
    exit 1
fi
