package main

import (
	// "github.com/sjxiang/hole/api/resthandler"
	"github.com/sjxiang/hole/internal/util"
	"github.com/sjxiang/hole/pkg/db"
)

// 构造组件
func Initialize()  (*Server, error) {

	// init env

	// init log
	sugaredLogger := util.NewSugardLogger()
	
	// init db
	dbConfig, err := db.GetConfig()
	if err != nil {
		return nil, err
	}
	gormDB, err := db.NewDbConnection(dbConfig, sugaredLogger)
	if err != nil {
		return nil, err
	}
	
	_ = gormDB
	// init repo

	// init service

	// init handler
	// resthandler.NewUserRestHandlerImpl()

	// init router
	
	// return NewServer()

	return &Server{
		logger: sugaredLogger,
	}, nil 
}





// 	// init repo
// 	appRepositoryImpl := repository.NewAppRepositoryImpl(sugaredLogger, gormDB)
// 	resourceRepositoryImpl := repository.NewResourceRepositoryImpl(sugaredLogger, gormDB)
// 	actionRepositoryImpl := repository.NewActionRepositoryImpl(sugaredLogger, gormDB)

// 	// init service
// 	asi = app.NewAppServiceImpl(sugaredLogger, appRepositoryImpl, kvstateRepositoryImpl, treestateRepositoryImpl, setstateRepositoryImpl, actionRepositoryImpl)
// 	rsi = resource.NewResourceServiceImpl(sugaredLogger, resourceRepositoryImpl)

// 	return nil
// }
