package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Problem: Concurrent Sensor Monitor

You need to simulate a system that monitors multiple sensors (e.g., temperature, pressure, humidity).

1. Launch 3 goroutines, each simulating a sensor that:
	Sleeps for a random time (≤ 3s).
	Then sends a reading (string or int) to a channel.
2. In the main goroutine, use a select loop to:
	Print whichever sensor reading comes in first.
	Continue listening until all 3 sensors have responded or until 4 seconds have passed.
3. If the timeout occurs first, print "Monitoring timed out" and exit.

Example Output:
Sensor 2: pressure=1020
Sensor 1: temperature=23
Sensor 3: humidity=54
All sensors responded!

or

Sensor 2: pressure=1021
Sensor 3: humidity=50
Monitoring timed out
*/

type Sensor struct {
	ID         int
	SensorType string
	Value      int
}

func main() {
	sensors := []Sensor{
		{1, "pressure", 1021},
		{2, "humidity", 50},
		{3, "temperature", 23},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	chSensorData := make(chan string, len(sensors))

	for _, v := range sensors {
		val := v
		wg.Go(func() {
			chSensorData <- getSensorData(ctx, val)
		})
	}

	go func() {
		wg.Wait()
		close(chSensorData)
	}()

	for sensor := range chSensorData {
		fmt.Println(sensor)
	}
}

func getSensorData(ctx context.Context, sensor Sensor) string {
	sleepDuration := time.Duration(rand.Intn(5)+1) * time.Second
	second := int(sleepDuration.Seconds())

	select {
	case <-time.After(sleepDuration):
		return fmt.Sprintf("Sensor %d retrieved data in %d seconds: %s=%d", sensor.ID, second, sensor.SensorType, sensor.Value)
	case <-ctx.Done():
		return fmt.Sprintf("Monitoring timed out for sensor %d after %d seconds", sensor.ID, second)
	}
}
