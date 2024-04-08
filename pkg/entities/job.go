package entities

import "time"

type Job struct {
	Id                                        uint64
	JobName, Description, Environment, Status string
	StartTime, EndTime                        time.Time
	Priority                                  uint16
}
