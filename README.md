## How to install

```bash
# Copy and change the following values in .env file
cp .env.example to .env
cp .env.test.example to .env.test

# Install dependencies
go install github.com/cosmtrek/air@latest # for development
go install github.com/joho/godotenv/cmd/godotenv@latest
go mod tidy
go install

# pre-commit hook
git config core.hooksPath .githooks
```

## Local Development

Use [air](https://github.com/cosmtrek/air) to enable hot reload on development, run command below

```
air init
air
```

Migration
```bash
# Create migration
go run . migrate --create=name_of_migration

# Run migration
go run . migrate
```

## Testing

### Create test cases

All test case placed in tests folder, example in pkg/articles/tests, pkg/comments/tests, etc

```go
// add this line below at the top of file
//go:build unit || integration || articles || pkg || all
// Definition of that line above:
// Text after 'go:build' called as 'tags'
// unit is for unit test [make sure to chose only one unit or integration]
// integration is for integration test [make sure to chose only one unit or integration]
// articles is for articles pkg, it will make easier to run specific package test
// pkg is for pkg test, it will run all test case in pkg directory
// all tags is for all test case, it will run all test case in project

// NOTE
// Feel free to use any tags you want, e.g. utils, main, migration, etc.

package tests

// code for test case below 
```

### Run test cases

- Unit testing

```
 godotenv -f .env.test go test -p 1 ./.../tests -tags=unit -v 
```

- Integration testing

```
 godotenv -f .env.test go test -p 1 ./.../tests -tags=integration -v 
```

- Specific pkg test, example articles pkg

```
 godotenv -f .env.test go test -p 1 ./.../tests -tags=articles -v 
```

- All test case under pkg directory

```
 godotenv -f .env.test go test -p 1 ./.../tests -tags=pkg -v 
```

- All test case

```
 godotenv -f .env.test go test -p 1 ./.../tests -tags=all -v 
```
