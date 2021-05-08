package app

import (
	mirror2 "simpleDbMirror/controllers/mirror"
	ping2 "simpleDbMirror/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping2.Ping)

	router.GET("/config", mirror2.Config)
	router.GET("/activate", mirror2.Activate)
	router.GET("/deactivate", mirror2.Deactivate)
	router.GET("/logs", mirror2.GetLogs)
	router.GET("/execute", mirror2.Execute)
	router.GET("/status", mirror2.Status)

}


