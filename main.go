package main

import (
	"grid/route"

	config "grid/conf"

	"github.com/sirupsen/logrus"
)

func main() {
	router := route.Init()
	router.Start(":" + config.Config.Port)

}
func init() {

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
