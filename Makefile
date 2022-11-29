GQL_PATH=./services/gql

gql/gen:
	cd ${GQL_PATH} && gqlgen generate --config config.yaml

gen: gql/gen