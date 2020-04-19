package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func requestIDGenerator() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// NewServer create new HTTP server
func NewServer(db *gorm.DB, port int) *http.Server {
	logrusWriter := logrus.New().Writer()
	defer logrusWriter.Close()

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      NewRouter(db),
		ErrorLog:     log.New(logrusWriter, "", 0),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return server
}
