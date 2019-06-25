package main

import (
	"github.com/dustin/go-humanize"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.SetFlags(log.Ltime | log.Lmicroseconds)

	log.Println(" ================= Start ================= ")

	res, err := http.DefaultClient.Get(`https://raw.githubusercontent.com/genez/dante/master/dante.txt`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(" ================= Got Response ================= ")

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(" ================= Content Downloaded ================= ")

	log.Printf(" ================= Size: %s ================= ", humanize.Bytes(uint64(len(content))))

	lines := strings.Split(string(content), "\n")



	for _, line := range lines {
		uc := strings.ToUpper(line)
		hash, _ := bcrypt.GenerateFromPassword([]byte(uc), bcrypt.DefaultCost)
		log.Printf("%s - %v", uc, hash)
	}

	log.Println(" ================= End ================= ")

	time.Sleep(1 * time.Hour)
}
