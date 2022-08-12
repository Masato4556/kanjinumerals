go test -cover ./... -coverprofile=temp/cover.out
go tool cover -html=temp/cover.out -o temp/cover.html
open temp/cover.html
