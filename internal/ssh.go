package internal

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func Connect(host string, port int, user string, password string, known_hosts string) {
	// ssh config
	//hostKeyCallback, err := knownhosts.New(known_hosts)
	//if err != nil {
	//	log.Fatal(err)
	//}
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// TODO retornar com o known_hosts
		// HostKeyCallback: hostKeyCallback,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// connect ot ssh server
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", host, port), config)
	if err != nil {
		log.Fatal("Error connecting to server. ", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatal("Error opening session. ", err)
	}
	defer session.Close()

	// configure terminal mode
	modes := ssh.TerminalModes{
		ssh.ECHO: 0, // supress echo

	}
	// run terminal session
	if err := session.RequestPty("xterm", 50, 80, modes); err != nil {
		log.Fatal("Error opening terminal. ", err)
	}
	// start remote shell
	if err := session.Shell(); err != nil {
		log.Fatal("Error starting remote shell. ", err)
	}

	/*
		var buff bytes.Buffer
		session.Stdout = &buff
		if err := session.Run("ls -la"); err != nil {
			log.Fatal("Error running command. ", err)
		}
		fmt.Println(buff.String())
	*/
}
