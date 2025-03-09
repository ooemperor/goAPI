package api

import (
	"goAPI/src/adapters"
)

type UserAdapter struct {
	*BaseApi
}

var baseApi = newBaseApi(adapters.UserAdapterObject)
var UserApiObject = UserAdapter{&baseApi}
