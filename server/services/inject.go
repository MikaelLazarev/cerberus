/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package services

import (
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/services/accounts"
	"github.com/MikaelLazarev/cerberus/server/services/notifications"
	"github.com/MikaelLazarev/cerberus/server/services/transactions"
	"github.com/MikaelLazarev/cerberus/server/services/users"
	"github.com/MikaelLazarev/cerberus/server/store"
)

type Services struct {
	TransactionsService core.TransactionsService
	AccountsService     core.AccountsService
	NotificationService core.NotificationService
	UserService         core.UserService
}

func InjectServices(Store store.GlobalStore) *Services {

	notificationService := notifications.New(Store.NotificationStore)
	transactionsService := transactions.New(Store.TransactionsStore, notificationService)
	accountsService := accounts.New(Store.AccountStore, transactionsService)

	return &Services{
		AccountsService:     accountsService,
		TransactionsService: transactionsService,
		NotificationService: notificationService,
		UserService:         users.New(Store.UserStore, accountsService, transactionsService),
	}
}
