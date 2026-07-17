# Prepare the environment

```powershell
go mod init "github.com/Azrubael/260601-golang-tutor/llinked_list"
go mod tidy
```


## To run the tests for all included packages:
```pwsh
go test -v ./...
```

## To test the package `azll`
```pwsh
go test -v ./azll
```

## To run tests and generate coverage profile
```pwsh
go test -coverprofile=coverage.out -v ./azll
```

## To run an exact test function `TestIfEqualAny` in the package `azll`
```pwsh
go test -run ^TestIfEqualAny$ -v ./azll
```