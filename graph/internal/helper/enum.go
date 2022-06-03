package helper

type Status int

const (
	Published Status = iota
	Unpublished
)

func (s Status) String() string {
	return [...]string{"Published", "Unpublished"}[s]
}

func (s Status) EnumIndex() int {
	return int(s)
}
