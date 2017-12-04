package main

import (
	"log"
	"sync"
	"time"
)

var (
	// error loggers
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	// internal session maps (id, username, timestamp, type of user)
	mu_user sync.RWMutex
	user_   map[string]string    = make(map[string]string)
	time_   map[string]time.Time = make(map[string]time.Time)
	// access to settings map
	mu_settings sync.RWMutex
	settings    map[string]string = make(map[string]string)
)
