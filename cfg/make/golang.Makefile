CMD_DIR=./src/go/cmd
GO_RUN=cd ${CMD_DIR}/$1 && go run . ${CMD_ARGS}

cli:
	@$(call GO_RUN,cli)
grpc:
	@$(call GO_RUN,grpc)
gql:
	@$(call GO_RUN,gql)
http:
	@$(call GO_RUN,http)