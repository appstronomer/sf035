package handler

import (
	"log"
	"net"
	"sf035/proverb"
	"time"
)

func Handle(conn net.Conn, proverb *proverb.Proverb) {
	defer conn.Close()
	for {
		_, err := conn.Write([]byte(proverb.Rnd()))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second * 3)
	}

}
