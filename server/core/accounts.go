package core

import "context"

type (
	Account struct {
		ID      ID     `json:"id"`
		Address string `json:"address"`
	}

	AccountsStore interface {
		// Stores account obj and return account ID
		Insert(ctx context.Context, account *Account) (accounntID ID, err error)

		// List returns Accounts for particular user from database
		ListByUser(ctx context.Context, userID ID) (result []*Account, err error)

		FindByID(ctx context.Context, accountID ID) (*Account, error)
	}

	AccountsService interface {
		Retrieve(ctx context.Context, userID, accountID ID) (*Account, error)
		ListByUser(ctx context.Context, userID ID) (result []*Account, err error)
	}
)
