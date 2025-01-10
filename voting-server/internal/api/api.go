package api

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func RunApi(mux *http.ServeMux) {
	port := os.Getenv("VOTING_SERVER_PORT")

	go func() {
		logrus.Infof("====> listening to port : %s\n", port)
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
		if err != nil {
			logrus.Fatalf("error trying to launch the server", err.Error())
		}

	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("Shutdown Server ...")

}
