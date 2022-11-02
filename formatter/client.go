package main

import (
	"context"

	"strings"
	"time"

	langPb "github.com/Indosaram/gomunkong-backend/proto/lang_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	langClient langPb.LangClient
)

func Format(port string, code string, formatter string, args []string) string {
	port = strings.Replace(port, "tcp://", "", 1)

	dialCtx, dialCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dialCancel()

	conn, _ := grpc.DialContext(dialCtx, port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	langClient = langPb.NewLangClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := langClient.Formatter(ctx, &langPb.Input{Code: code, Formatter: formatter, Args: args})
	if err != nil {
		panic(err)
	}

	return resp.FormattedCode
}
