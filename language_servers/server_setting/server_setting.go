package server_setting

type ServerPort struct {
	PythonPort     string
	JavascriptPort string
	JavaPort       string
	SwiftPort      string
	GolangPort     string
}

func NewServerPort() *ServerPort {
	s := new(ServerPort)
	s.PythonPort = "9001"
	s.JavascriptPort = "9002"
	s.JavaPort = "9003"
	s.SwiftPort = "9004"
	s.GolangPort = "9005"

	return s
}
