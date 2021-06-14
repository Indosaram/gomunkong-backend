package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	formatter "github.com/Indosaram/gomunkong-backend/formatter/file_io"
	"github.com/Indosaram/gomunkong-backend/language_servers/server_setting"
	langPb "github.com/Indosaram/gomunkong-backend/proto/lang_server"
	"google.golang.org/grpc"
)

type pythonServer struct {
	langPb.LangServer
}

func (s *pythonServer) Formatter(ctx context.Context, input *langPb.Input) (*langPb.FormattedCode, error) {
	ext := "py"
	formatter.WriteTempFile(ext, input.Code)
	tempFilePath := "./tmp/temp." + ext

	var command string
	var argstr []string
	switch formatterType := input.Formatter; formatterType {
	case "black":
		command = "python3"
		argstr = []string{"-m", "black"}
		argstr = append(argstr, input.Args...)
		argstr = append(argstr, tempFilePath)
	case "autopep8":
		command = "autopep8"
		argstr = []string{tempFilePath}
	case "yapf":
		command = "yapf"
		argstr = []string{tempFilePath}
	default:
		fmt.Println("Not supported yet")
	}
	formatter.FormatFile(tempFilePath, command, argstr)

	data, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		fmt.Println("Failed to read temp file")
		panic(err)
	}

	return &langPb.FormattedCode{
		FormattedCode: string(data),
	}, nil
}

func main() {
	ports := server_setting.NewServerPort()
	portNumber := ports.PythonPort
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	langPb.RegisterLangServer(grpcServer, &pythonServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
