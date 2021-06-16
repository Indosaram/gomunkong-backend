package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"

	formatter "github.com/Indosaram/gomunkong-backend/formatter/file_io"
	"github.com/Indosaram/gomunkong-backend/language_servers/server_setting"
	langPb "github.com/Indosaram/gomunkong-backend/proto/lang_server"
	"google.golang.org/grpc"
)

type javaServer struct {
	langPb.LangServer
}

func (s *javaServer) Formatter(ctx context.Context, input *langPb.Input) (*langPb.FormattedCode, error) {
	ext := "java"
	formatter.WriteTempFile(ext, input.Code)
	tempFilePath := "./tmp/temp." + ext

	var command string
	var argstr []string
	switch formatterType := input.Formatter; formatterType {
	case "google-java-format":
		command = "java"
		argstr = []string{"-jar", "google-java-format.jar"}
		argstr = append(argstr, input.Args...)
		argstr = append(argstr, tempFilePath)
	default:
		fmt.Println("Not supported yet")
	}
	formatter.FormatFile(tempFilePath, command, argstr)

	data, err := ioutil.ReadFile(tempFilePath)
	if err != nil {
		fmt.Printf("Failed to read temp file")
	}

	return &langPb.FormattedCode{
		FormattedCode: string(data),
	}, err
}

func main() {
	ports := server_setting.NewServerPort()
	portNumber := ports.JavaPort
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	langPb.RegisterLangServer(grpcServer, &javaServer{})

	fmt.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %s", err)
	}
}
