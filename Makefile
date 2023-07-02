gen_auth_serv:
	protoc --go_out=plugins=grpc:auth protos/auth/auth.proto

gen_item_serv:
	protoc --go_out=plugins=grpc:item protos/item/item.proto

