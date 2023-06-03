git add .
git commit -m "New checkpoint"
git push

set GOOS = linux
set GOARCH = amd64

go build main.go
tar.exe -a -cf main.zip main