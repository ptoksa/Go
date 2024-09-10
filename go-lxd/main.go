package main

import (
	"fmt"
	"log"
	"time"

	incus "github.com/lxc/incus/client"    // Use the alias 'incus' for this import
	"github.com/lxc/incus/shared/api"      // Updated module path
)

func main() {
	// Connect to the local LXD (Incus) instance using the correct Unix socket path
	server, err := incus.ConnectIncusUnix("/var/snap/lxd/common/lxd/unix.socket", nil)
	if err != nil {
		log.Fatalf("Failed to connect to LXD: %v", err)
	}

	// Define the container creation request
	req := api.InstancesPost{
		Name: "my-container",
		Source: api.InstanceSource{
			Type:  "image",
			Alias: "ubuntu-22.04", // Ensure this alias is correct and the image is available
		},
	}

	// Create the container
	op, err := server.CreateInstance(req)
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	// Wait for the creation operation to finish
	err = op.Wait()
	if err != nil {
		log.Fatalf("Error waiting for operation: %v", err)
	}
	fmt.Println("Container created successfully")

	// Start the container
	op, err = server.UpdateInstanceState("my-container", api.InstanceStatePut{
		Action:  "start",
		Timeout: -1,
	}, "")
	if err != nil {
		log.Fatalf("Failed to start container: %v", err)
	}

	// Wait for the container to start
	err = op.Wait()
	if err != nil {
		log.Fatalf("Error waiting for container start: %v", err)
	}
	fmt.Println("Container started successfully")

	// Wait for some time before stopping the container
	time.Sleep(10 * time.Second)

	// Stop the container
	op, err = server.UpdateInstanceState("my-container", api.InstanceStatePut{
		Action:  "stop",
		Timeout: -1,
		Force:   true,
	}, "")
	if err != nil {
		log.Fatalf("Failed to stop container: %v", err)
	}

	// Wait for the container to stop
	err = op.Wait()
	if err != nil {
		log.Fatalf("Error waiting for container stop: %v", err)
	}
	fmt.Println("Container stopped successfully")
}
