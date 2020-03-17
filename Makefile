
build :
	go build ./cmd/dsxutl

test :
	go test -v ./integration

clean :
	del dsxutl.exe

.DEFAULT_GOAL = build
