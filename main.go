package main

import (
	"fmt"
	"log"
	"net/http"
)

var looks = map[string]string{
	"punching-fury":       "(ง •̀_•́)ง",
	"stoned-look":         "◙‿◙",
	"wink-look":           "◕‿↼",
	"bear-look":           "(`･ω･´)",
	"lenny-blank-stare":   "( ͡o ͜ʖ ͡o)",
	"flipping-table":      "(╯°□°)╯︵ ┻━┻",
	"epic-table-flip":     "(˚Õ˚)ر ~~~~╚╩╩╝",
	"flipping-table-look": "(┛ಠ_ಠ)┛彡┻━┻",
	"sad-look":            "ʘ︵ʘ",
	"so-sorry":            "（ﾉ´д｀）",
	"cute-look":           "•‿•",
	"puss-in-boots-look":  "◕ ◡ ◕",
	"sending-kisses":      "( ˘ ³˘)♥",
}

func looksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	look, ok := looks[r.URL.Path[1:]]
	if !ok {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", look)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprintf(w, "try another looks\n")
	}
}

func main() {
	http.HandleFunc("/", looksHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
