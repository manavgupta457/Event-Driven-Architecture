package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	wearablepb "practice.com/grpc_test/gen/go/wearable/v1"
)

func main() {
	// opts := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// }

	// conn, err := grpc.Dial("localhost:9879", opts...)
	// if err != nil {
	// 	log.Fatalf("fail to dial: %v", err)
	// }

	// defer conn.Close()
	// client := userpb.NewUserServiceClient(conn)

	// res, err := client.GetUser(context.Background(), &userpb.GetUserRequest{
	// 	Uuid: "1-2-3",
	// })
	// if err != nil {
	// 	log.Fatalf("fail to GetUser: %v", err)
	// }

	// fmt.Printf("%+v\n", res)

	// wearable_main()
	wearable_clientstream_main()

}

func wearable_main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9879", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := wearablepb.NewWearableServiceClient(conn)

	if err := ui.Init(); err != nil {
		log.Fatalf("Couldn't init UI: %v", err)
	}
	defer ui.Close()

	lc := widgets.NewPlot()
	lc.Title = "Heartbeat Per Minute"
	lc.SetRect(0, 0, 70, 20)
	lc.Data = make([][]float64, 1)
	lc.DataLabels = []string{"Minute", "Value"}
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorGreen
	lc.Marker = widgets.MarkerDot

	data := make([]float64, 60)
	ui.Render(lc)

	//-

	res, err := client.BeatsPerMinute(context.Background(), &wearablepb.BeatsPerMinuteRequest{Uuid: "mario"})
	if err != nil {
		log.Fatalln("Couldn't request", err)
	}

	go func() {
		for {
			resp, err := res.Recv()
			if err == io.EOF {
				return
			}

			if err != nil {
				log.Fatalln("Receiving", err)

				return
			}

			data[int(resp.GetMinute())] = float64(resp.GetValue())
			lc.Data[0] = data

			ui.Render(lc)
		}
	}()

	//-

	uiEvents := ui.PollEvents()

	for {
		select {
		case e := <-uiEvents: // Listen for Keyboard
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-res.Context().Done():
			fmt.Println("All done, possible error", res.Context().Err())
		}
	}
}

func wearable_clientstream_main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:9878", opts...)

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := wearablepb.NewWearableServiceClient(conn)
	stream, err := client.ConsumeBeatsPerMinute(context.Background())

	if err != nil {
		log.Fatalln("Consuming stream", err)
	}

	for i := 0; i <= 10; i++ {
		err := stream.Send(&wearablepb.ConsumeBeatsPerMinuteRequest{
			Uuid:   "Manav",
			Value:  100,
			Minute: uint32(i),
		})
		if err != nil {
			log.Fatalln("Sending value", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Closing", err)
	}

	fmt.Println("Total Messages", resp.GetTotal())
}
