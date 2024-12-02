package app_config

import "os"

var PORT = ":8000"
var STATIC_ROUTE = "/public"
var STATIC_DIR = "./public"

func InitAppConfig() {
	portEnv := os.Getenv("APP_PORT")
	if portEnv != "" {
		PORT = portEnv
	}
}

func StaticRouteEnv() {
	staticRouteEnv := os.Getenv("STATIC_ROUTE")
	if staticRouteEnv != "" {
		STATIC_ROUTE = staticRouteEnv
	}
}

func StaticDirEnv() {
	staticDirEnv := os.Getenv("STATIC_DIR")
	if staticDirEnv != "" {
		STATIC_DIR = staticDirEnv
	}
}