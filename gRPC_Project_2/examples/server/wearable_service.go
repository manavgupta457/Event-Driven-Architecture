package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	wearablepb "practice.com/grpc_test/gen/go/wearable/v1"
)

type wearableServer struct {
	wearablepb.UnimplementedWearableServiceServer
}

func (w *wearableServer) BeatsPerMinute(
	req *wearablepb.BeatsPerMinuteRequest,
	stream wearablepb.WearableService_BeatsPerMinuteServer) error {

	for {
		select {

		case <-stream.Context().Done():
			return status.Error(codes.Canceled, "Stream has ended")

		default:
			time.Sleep(1 * time.Second)
			value := 30 + rand.Int31n(80)
			err := stream.SendMsg(&wearablepb.BeatsPerMinuteResponse{
				Value:  uint32(value),
				Minute: uint32(time.Now().Second()),
			})
			if err != nil {
				return status.Error(codes.Canceled, "Stream has ended")
			}
		}
	}

}

func (w *wearableService) ConsumeBeatsPerMinute(stream wearablepb.WearableService_ConsumeBeatsPerMinuteServer) error {
	var total uint32
	for {
		value, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&wearablepb.ConsumeBeatsPerMinuteResponse{
				Total: total,
			})
		}

		if err != nil {
			return err
		}

		fmt.Println(value.GetMinute(), value.GetUuid(), value.GetValue())
		total++
	}
}
