package main

import (
	"ServerGo/internal/handler"
	"ServerGo/internal/sqlite"
	"net/http"
)

func main() {
	_, err := sqlite.InitDataBase("./F:/tmp/F.db")

	http.HandleFunc("/meeting-room", handler.MeetingRoomHandler)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}
