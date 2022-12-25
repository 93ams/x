package pattern

import (
	"github.com/tilau2328/x/src/go/service/repo/go/package/domain/model"
)

type AdaptorProps struct {
	Provider, Models, Grpc, Proto model.SearchReq
	Mappers, Handler, Requester   string
}

func Adaptor() {

}
