package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/go-sql-driver/mysql"
	"github.com/rancher/cluster-manager/cluster"
)

func main() {
	config := mysql.Config{
		User:      "cattle",
		Passwd:    "cattle",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "cattle",
		Collation: "utf8_general_ci",
	}

	ip := os.Getenv("IP")
	if ip == "" {
		ip = "127.0.0.1"
	}
	cluster, err := cluster.New("mysql", config.FormatDSN(), ip, 3)
	if err != nil {
		logrus.WithField("err", err).Fatalf("Failed to create manager")
	}

	if err := cluster.Start(); err != nil {
		logrus.WithField("err", err).Fatalf("Failed to create manager")
	}
}
