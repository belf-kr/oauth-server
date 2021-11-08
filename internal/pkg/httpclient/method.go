package httpclient

type Method int

const (
	GET Method = iota
	POST
)

func (m Method) String() string {
	return []string{
		"GET",
		"POST",
	}[m]
}
