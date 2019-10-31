#!/bin/bash
cd ${GITHUB_WORKSPACE}
golangci-lint run ${INPUT_GOLANGCI_LINT_ARGS}
