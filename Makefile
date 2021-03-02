
build: 
	go build -ldflags "-X go-build-variable-sample/types.GitCommit=`git rev-parse --short HEAD`" -o app