package main

import (
	"fmt"
	"github.com/digitalocean/go-libvirt"
	"github.com/sgreben/sshtunnel"
	"golang.org/x/crypto/ssh"
	"log"
)

func main() {
	// Connect to "google.com:80" via a tunnel to "ubuntu@my-ssh-server-host:22"
	keyPath := "/home/matteo/.ssh/id_rsa"
	authConfig := sshtunnel.ConfigAuth{
		Keys:     []sshtunnel.KeySource{{Path: &keyPath}},
	}
	sshAuthMethods, _ := authConfig.Methods()
	clientConfig := ssh.ClientConfig{
		User: "madeo",
		Auth: sshAuthMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	tunnelConfig := sshtunnel.Config{
		SSHAddr: "homelab:22",
		SSHClient: &clientConfig,
	}
	c, _, err := sshtunnel.Dial("unix", "/var/run/libvirt/libvirt-sock", &tunnelConfig)
	if err != nil {
		panic(err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	fmt.Println("Connected!")
}
