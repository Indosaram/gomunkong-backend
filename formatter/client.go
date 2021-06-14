package main

import (
	"context"
	"strings"
	"time"

	langPb "github.com/Indosaram/gomunkong-backend/proto/lang_server"
	"google.golang.org/grpc"
)

var (
	langClient langPb.LangClient
)

func Format(port string, code string, formatter string, args []string) string {
	port = strings.Replace(port, "tcp://", "", 1)
	conn, _ := grpc.Dial(port,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	print(1)
	langClient = langPb.NewLangClient(conn)

	print(2)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	print(3)
	resp, err := langClient.Formatter(ctx, &langPb.Input{Code: code, Formatter: formatter, Args: args})
	if err != nil {
		panic(err)
	}

	print(4)
	return resp.FormattedCode
}
