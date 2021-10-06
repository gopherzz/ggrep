buildWin:
	go build -o bin/ggrep.exe app/main.go

buildLinux:
	go build -o bin/ggrep app/main.go
