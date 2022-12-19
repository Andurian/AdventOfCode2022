@echo off
mkdir %1 
cd %1
go mod init andurian/adventofcode/2022/%1
go mod edit -replace andurian/adventofcode/2022/util=../util
cd ..
copy /y template.txt %1\main.go >NUL
go work use %1
cd %1
copy /y NUL input.txt >NUL
copy /y NUL input_test.txt >NUL
go mod tidy