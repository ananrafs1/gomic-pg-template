corePlugin := ${GOPATH}//src//github.com//ananrafs1//gomic
.DEFAULT_GOALS := create
create:
	@go build -o ${corePlugin}//plugins//scrapper//template.exe main.go
