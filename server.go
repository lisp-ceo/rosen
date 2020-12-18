package main

import (
	"errors"

	"github.com/lisp-ceo/rosen/protocols/config"
	"github.com/lisp-ceo/rosen/protocols/https"
	"github.com/lisp-ceo/rosen/proxy"
)

func server(conf config.Configuration) (err error) {
	var server proxy.Server

	switch conf["protocol"] {
	case "":
		return errors.New("protocol must be specified in config file")
	case "https":
		server, err = https.NewServer(conf)
	default:
		return errors.New("unknown protocol: " + conf["protocol"])
	}
	if err != nil {
		return err
	}

	return server.Start()
}
