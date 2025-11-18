#!/bin/bash

set -e

# Configuration
GO_MIN_COVERAGE=90
VUE_MIN_COVERAGE=80

# Track overall pass/fail status
EXIT_CODE=0

echo "========================================"
echo "=== Running Go Tests with Coverage ==="
echo "========================================"
echo "Note: Excluding cmd package (main packages are typically not tested)"
go test -coverprofile=coverage.out ./internal/...

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
    echo "❌ FAILED: Go coverage ${TOTAL_COVERAGE}% is below minimum ${GO_MIN_COVERAGE}%"
    EXIT_CODE=1
else
    echo "✅ PASSED: Go coverage ${TOTAL_COVERAGE}% meets minimum ${GO_MIN_COVERAGE}%"
fi

echo ""
echo "========================================="
echo "=== Running Vue Tests with Coverage ==="
echo "========================================="

# Change to web directory and run tests with coverage
cd web

# Run Vitest with coverage and capture output
npm run test:coverage -- --run 2>&1 | tee /tmp/vue-coverage.txt

# Parse coverage from Vitest text output
# Look for the "All files" line in the coverage table
if grep -q "All files" /tmp/vue-coverage.txt; then
    # Extract coverage percentages from the "All files" line
    # Format: All files          |     100 |    96.66 |     100 |     100 |
    # Columns are: File | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s
    COVERAGE_LINE=$(grep "All files" /tmp/vue-coverage.txt | tr -s ' ')

    # Extract individual coverage metrics (columns: Stmts, Branch, Funcs, Lines)
    VUE_STATEMENTS=$(echo "$COVERAGE_LINE" | awk -F '|' '{print $2}' | tr -d ' ')
    VUE_BRANCHES=$(echo "$COVERAGE_LINE" | awk -F '|' '{print $3}' | tr -d ' ')
    VUE_FUNCTIONS=$(echo "$COVERAGE_LINE" | awk -F '|' '{print $4}' | tr -d ' ')
    VUE_LINES=$(echo "$COVERAGE_LINE" | awk -F '|' '{print $5}' | tr -d ' ')

    echo ""
    echo "=== Vue Coverage Summary ==="
    echo "Statements: ${VUE_STATEMENTS}%"
    echo "Branches:   ${VUE_BRANCHES}%"
    echo "Functions:  ${VUE_FUNCTIONS}%"
    echo "Lines:      ${VUE_LINES}%"
    echo "Minimum Required: ${VUE_MIN_COVERAGE}%"

    # Check all metrics against threshold
    FAILED_METRICS=""
    if (( $(echo "$VUE_STATEMENTS < $VUE_MIN_COVERAGE" | bc -l) )); then
        FAILED_METRICS="${FAILED_METRICS}Statements (${VUE_STATEMENTS}%) "
    fi
    if (( $(echo "$VUE_BRANCHES < $VUE_MIN_COVERAGE" | bc -l) )); then
        FAILED_METRICS="${FAILED_METRICS}Branches (${VUE_BRANCHES}%) "
    fi
    if (( $(echo "$VUE_FUNCTIONS < $VUE_MIN_COVERAGE" | bc -l) )); then
        FAILED_METRICS="${FAILED_METRICS}Functions (${VUE_FUNCTIONS}%) "
    fi
    if (( $(echo "$VUE_LINES < $VUE_MIN_COVERAGE" | bc -l) )); then
        FAILED_METRICS="${FAILED_METRICS}Lines (${VUE_LINES}%) "
    fi

    if [ -n "$FAILED_METRICS" ]; then
        echo "❌ FAILED: Vue coverage below minimum ${VUE_MIN_COVERAGE}%: ${FAILED_METRICS}"
        EXIT_CODE=1
    else
        echo "✅ PASSED: All Vue coverage metrics meet minimum ${VUE_MIN_COVERAGE}%"
    fi
else
    echo "❌ FAILED: Could not parse Vue coverage output"
    EXIT_CODE=1
fi

cd ..

echo ""
echo "========================================="
echo "=== Coverage Check Summary ==="
echo "========================================="
if [ $EXIT_CODE -eq 0 ]; then
    echo "✅ All coverage requirements met!"
else
    echo "❌ Coverage requirements not met. See details above."
fi

exit $EXIT_CODE
