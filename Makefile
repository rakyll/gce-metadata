binary = gce-metadata

release:
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(binary)
	gsutil cp bin/$(binary) gs://jbd-releases