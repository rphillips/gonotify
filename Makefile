all: deps
	go build .

deps:
	go get -u github.com/stevenleeg/gowl
	go get -u github.com/vaughan0/go-ini
	go get -u bitbucket.org/kisom/gopush/pushover

.PHONY: all deps
