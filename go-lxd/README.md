# Go-LXD Project
## Overview
This project demonstrates how to interact with LXD (Incus) using Go. It includes functionality to create, start, and stop LXD containers programmatically. The project uses the `github.com/lxc/incus/client` package to communicate with the LXD server.
## Prerequisites
- Go: Version 1.18 or later.
- LXD (Incus): Installed and running. Ensure LXD is set up properly and that you have necessary permissions.
- LXD Image: An LXD image must be available to create containers. This project uses `ubuntu-22.04` as an example image.
## Installation
1. Clone the Repository:
```sh
git clone https://github.com/yourusername/go-lxd.git
cd go-lxd
```
2. Install Dependencies:
Ensure you have Go modules enabled and install the required dependencies:
```sh
go mod tidy
```
3. Set Up LXD:
Ensure LXD is installed and running:
```sh
sudo snap install lxd
sudo lxd init
```
Add an image if necessary:
```sh
sudo lxc image copy ubuntu:22.04 local: --alias ubuntu-22.04
```
## Usage
1. Create and Start a Container:
Ensure your `main.go` file is set up with the correct configuration. Run the application using:
```sh
sudo /usr/local/go/bin/go run main.go
```
This command will create, start, and then stop a container named `my-container` using the `ubuntu-22.04` image.
2. Verify Container Status:
After running the application, verify the container status using:
```sh
sudo lxc list
```
## Code Overview
- `main.go`: The main application file. It connects to the LXD server, creates a container, starts it, waits for some time, and then stops the container.