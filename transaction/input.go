package transaction

import "bluelight/user"

type GetTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
