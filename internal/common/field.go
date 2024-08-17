package common

type CronField interface {
	Expand() ([]int, error)
	Validate() error
}
