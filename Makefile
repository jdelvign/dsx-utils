
build : clean
	go build

test : clean build
	go test -v

install : clean
	go install

clean :
	if exist dsxutl.exe del dsxutl.exe

.DEFAULT_GOAL = build
