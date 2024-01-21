package internal

type CronField interface {
	Expand() ([]int, error)
	Validate() error
}
