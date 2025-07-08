package constants

type Status string

const (
	PENDING Status = "pending"
	DONE    Status = "done"
)

func (s Status) String() string {
	return string(s)
}

func (s Status) Validate() bool {
	return s == PENDING || s == DONE
}
