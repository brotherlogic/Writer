package main

import (
	"fmt"
	"strconv"
)

func (s *Server) handle(name, lat, lon, ip, port string) string {
	latval, err := strconv.ParseFloat(lat, 32)
	if err != nil {
		return fmt.Sprintf("Error handling location update %v", err)
	}
	lonval, err := strconv.ParseFloat(lon, 32)
	if err != nil {
		return fmt.Sprintf("Error handling location update %v", err)
	}

	err = s.locwriter.writeLocation(name, float32(latval), float32(lonval), ip, port)
	if err != nil {
		return fmt.Sprintf("Error writing Location: %v", err)
	}
	return fmt.Sprintf("Location Written")
}
