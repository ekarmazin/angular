.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/s3parser s3parser/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/scheduler scheduler/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/manual manual/main.go


clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
