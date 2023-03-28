#!/bin/bash 

echo "@@ Installing golangci-lint"

brew tap golangci/tap
brew install golangci/tap/golangci-lint

echo "@@ Installing pre-commit hook"

PRECOMMIT_FILE=.git/hooks/pre-commit

if [ -f "$PRECOMMIT_FILE" ]; then 
	echo "golangci-lint run" >> $PRECOMMIT_FILE
else
	cp .pre-commit $PRECOMMIT_FILE 
	chmod +x $PRECOMMIT_FILE
fi

echo "@@ Done"
