package adapters

import (
	"goAPI/src/db/handlers"
)

type UserAdapter struct {
	*BaseAdapter
}

var baseAdapter = NewBaseAdapter(handlers.UserHandlerObject)
var UserAdapterObject = UserAdapter{&baseAdapter}
