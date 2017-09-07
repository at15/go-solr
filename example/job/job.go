package main

import "time"

type Job struct {
	Id        string
	Namespace string
	Pool      string
	Created   time.Time
	Updated   time.Time
	Type      int
}
