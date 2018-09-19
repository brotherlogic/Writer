package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"

	pb "github.com/brotherlogic/location/proto"
)

type locwriter interface {
	writeLocation(name string, lat, lon float32, ip, port string) error
}

// Server main server type
type Server struct {
	ip        string
	port      string
	locwriter locwriter
}

type mainWriter struct{}

func (m *mainWriter) writeLocation(name string, lat, lon float32, ip, port string) error {
	conn, err := grpc.Dial(ip+":"+port, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	proxy := pb.NewLocationServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = proxy.AddLocation(ctx, &pb.AddLocationRequest{Location: &pb.Location{Name: name, Lat: lat, Lon: lon, Time: time.Now().Unix()}})
	return err
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	vals := strings.Split("/", r.URL.Path[1:])
	if len(vals) != 3 {
		fmt.Fprintf(w, "Can't handle request: %v -> %v from %v", len(vals), vals, r.URL.Path[1:])
	} else {
		fmt.Fprintf(w, s.handle(vals[0], vals[1], vals[2], s.ip, s.port))
	}
}

func main() {
	s := &Server{os.Args[1], os.Args[2], &mainWriter{}}
	http.HandleFunc("/", s.handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
