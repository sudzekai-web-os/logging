package logging

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"

	"github.com/sudzekai-web-os/logging/proto"
)

type GrpcWriter struct {
	conn   *grpc.ClientConn
	client proto.LoggerClient
}

func NewGrpcWriter(endpoint string) abstractions.ILoggerWriter {
	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return &GrpcWriter{}
	}

	return &GrpcWriter{
		conn:   conn,
		client: proto.NewLoggerClient(conn),
	}
}

func (gw *GrpcWriter) Write(entry types.LogEntry) {
	if gw.client == nil {
		return
	}

	_, _ = gw.client.Write(
		context.Background(),
		logEntryToProto(entry),
	)
}

func (gw *GrpcWriter) WriteBatch(entries []types.LogEntry) {
	if gw.client == nil {
		return
	}

	items := make([]*proto.LogEntry, 0, len(entries))

	for _, entry := range entries {
		items = append(items, logEntryToProto(entry))
	}

	_, _ = gw.client.WriteBatch(
		context.Background(),
		&proto.LogBatch{
			Entries: items,
		},
	)
}

func logEntryToProto(entry types.LogEntry) *proto.LogEntry {
	return &proto.LogEntry{
		Timestamp:   timestamppb.New(entry.TimeStamp),
		PreCategory: entry.PreCategory,
		Category:    entry.Category,
		SubCategory: entry.SubCategory,
		LogLevel:    logLevelToProto(entry.LogLevel),
		Message:     entry.Message,
	}
}

func logLevelToProto(level types.LogLevel) proto.LogLevel {
	switch level {
	case types.None:
		return proto.LogLevel_NONE
	case types.Debug:
		return proto.LogLevel_DEBUG
	case types.Information:
		return proto.LogLevel_INFORMATION
	case types.Warning:
		return proto.LogLevel_WARNING
	case types.Error:
		return proto.LogLevel_ERROR
	case types.Critical:
		return proto.LogLevel_CRITICAL
	default:
		return proto.LogLevel_NONE
	}
}
