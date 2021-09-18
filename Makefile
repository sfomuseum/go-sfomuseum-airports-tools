cli: 	
	go build -mod vendor -o bin/build-sfomuseum-data cmd/build-sfomuseum-data/main.go

data: sfomuseum-data

sfomuseum-data:
	bin/build-sfomuseum-data > /usr/local/sfomuseum/go-sfomuseum-airports/sfomuseum/data.go
