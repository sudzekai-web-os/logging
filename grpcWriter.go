package logging

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/sudzekai-web-os/core"
	"github.com/sudzekai-web-os/logging/proto"
)

type GrpcWriter struct {
	conn   *grpc.ClientConn
	client proto.LoggerClient
}

func NewGrpcWriter(endpoint string) core.ILoggerWriter {
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

func (gw *GrpcWriter) Write(entry core.LogEntry) {
	if gw.client == nil {
		return
	}

	_, _ = gw.client.Write(
		context.Background(),
		logEntryToProto(entry),
	)
}

func (gw *GrpcWriter) WriteBatch(entries []core.LogEntry) {
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

func logEntryToProto(entry core.LogEntry) *proto.LogEntry {
	return &proto.LogEntry{
		Timestamp:   timestamppb.New(entry.TimeStamp),
		PreCategory: entry.PreCategory,
		Category:    entry.Category,
		SubCategory: entry.SubCategory,
		LogLevel:    logLevelToProto(entry.LogLevel),
		Message:     entry.Message,
	}
}

func logLevelToProto(level core.LogLevel) proto.LogLevel {
	switch level {
	case core.LogLevel_NONE:
		return proto.LogLevel_NONE
	case core.LogLevel_DEBUG:
		return proto.LogLevel_DEBUG
	case core.LogLevel_INFORMATION:
		return proto.LogLevel_INFORMATION
	case core.LogLevel_WARNING:
		return proto.LogLevel_WARNING
	case core.LogLevel_ERROR:
		return proto.LogLevel_ERROR
	case core.LogLevel_CRITICAL:
		return proto.LogLevel_CRITICAL
	default:
		return proto.LogLevel_NONE
	}
}
