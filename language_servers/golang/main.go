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

type golangServer struct {
	langPb.LangServer
}

func (s *golangServer) Formatter(ctx context.Context, input *langPb.Input) (*langPb.FormattedCode, error) {
	ext := "go"
	formatter.WriteTempFile(ext, input.Code)
	tempFilePath := "./tmp/temp." + ext

	var command string
	var argstr []string
	switch formatterType := input.Formatter; formatterType {
	case "gofmt":
		command = "gofmt"
		argstr = []string{"-w"}
		argstr = append(argstr, input.Args...)
		argstr = append(argstr, tempFilePath)
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
	portNumber := ports.GolangPort
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	langPb.RegisterLangServer(grpcServer, &golangServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
