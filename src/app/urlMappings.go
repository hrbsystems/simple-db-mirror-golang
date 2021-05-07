package app

import (
	mirror2 "simpleDbMirror/controllers/mirror"
	ping2 "simpleDbMirror/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping2.Ping)
	router.GET("/db-mirror-config", mirror2.GetDbMirrorConfig)
}
