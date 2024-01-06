package initialize

import "github.com/jaronnie/ragingcd/core/server/engine/db"

func Initialize() {
	if err := db.InitDB(); err != nil {
		panic(err)
	}

	if err := db.SyncTable(); err != nil {
		panic(err)
	}
}
