protoc-gen:
	protoc --twirp_out=. --go_out=. protoc/niltasks.proto