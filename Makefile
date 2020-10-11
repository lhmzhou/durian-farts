test:
	go test

cover:
	go test -cover

cover-report:
	go test -covermode=count -coverprofile=.cover
	go tool cover -html=.cover