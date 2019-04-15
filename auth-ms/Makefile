_MICROSERVICES = auth-ms

proto:
	for service in $(_MICROSERVICES) ; do \
   		protoc -I/usr/local/include \
   				-I. \
   				-I$(GOPATH)/src \
   				--go_out=plugins=grpc:. \
   				./pb/service.proto; \
   	done

