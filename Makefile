.PHONY: pb

pb:
	protoc --go_out=. ./hello/*.proto
