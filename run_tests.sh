gofmt -s -w .
go test -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
rm cover.out
