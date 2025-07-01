clear

go mod tidy  || exit $?
go run .     || exit $?
