package main

import (
	"log"
)

func main() {
    trace: very detailed information about the program run, usually for debugging purposes
    debug: used to output very detailed information, usually for debugging purpose
    info: informational logs
    warn (warning): a non-critical error or event that we should fix. The application can still operate.
    error: an error that we should fix. The application can still operate.
    fatal: a critical error, the application will be stopped
    panic: the highest level of criticality, the application will be stopped


	log.Trace("going to mars") // trace: very detailed information about the program run, usually for debugging purposes
	log.Debug("connected, received buffer from worker") //debug: used to output very detailed information, usually for debugging purpose
	log.Info("connection successful to db") //info: informational logs
	log.Warn("something went wrong with x") //warn (warning): a non-critical error or event that we should fix. The application can still operate.
	log.Error("an error occurred in worker x") //error: an error that we should fix. The application can still operate.
	log.Fatal("impossible to continue exec") //fatal: a critical error, the application will be stopped
	log.Panic("system emergency shutdown") panic: the highest level of criticality, the application will be stopped
}
