go build srk.go
GOOS=windows GOARCH=amd64 go build -o srk.exe srk.go
GOOS=darwin GOARCH=amd64 go build -o srk.mac srk.go 
# GOOS=darwin GOARCH=arm64 go build -o srk.m1 srk.go // fails  