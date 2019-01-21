package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

//4 different structs to define: Config, Users, Messages, and  ChatServer.

type Config struct {
	Server struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Protocol string `json:"protocol"`
	} `json:"server"`
	ChatBot         string `json:"chatBot"`
	LogFileLocation string `json:"logFileLocation"`
}

type User struct {
	Name   string
	Output chan Message
}

type Message struct {
	Username string
	Text     string
}

type ChatServer struct {
	Users map[string]User
	Join  chan User
	Leave chan User
	Input chan Message
}

//The config file for this application is imported as JSON.  JSON gives us an easy to use config format that most engineers can work with.
//The following function can be used to turn the config file(json) into a config format (struct) that can be used throughout the chat application.

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		log.Fatal("Problem loading Configuration File.")
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

//all of the logic needed for chat server.
func (cs *ChatServer) Run() {
	config, _ := LoadConfiguration("config.json")
	//3 different cases possible: 1. a user joins the chatroom, 2. a user leaves the chatroom, 3. a message is sent to the chat server as user input to be distributed to the full range of users in the chatroom.
	for {
		select {
		//join a user to the chat server, display some text in the terminal letting all present existing users know a user has joined and also, update map of users.
		case user := <-cs.Join:
			cs.Users[user.Name] = user

			//included in a go routine so that the for loop can continue.
			go func() {
				cs.Input <- Message{
					Username: config.ChatBot,
					Text:     fmt.Sprintf("%s joined Torbit Chat", user.Name),
				}
			}()
		case user := <-cs.Leave:
			delete(cs.Users, user.Name)

			//included in a go routine so that the for loop can continue.
			go func() {
				cs.Input <- Message{
					Username: config.ChatBot,
					Text:     fmt.Sprintf("%s left Torbit Chat", user.Name),
				}
			}()

			//input from user (chat join, chat leave)goes to the output user channel and is distributed via for loop to all users in the chat room.  Input is in the form of a message struct and the for loop sends it as user.output in the desired one message to many users fashion.  Note, this is not for sending message text.
		case msg := <-cs.Input:
			for _, user := range cs.Users {
				select {
				case user.Output <- msg:
				default:
				}
			}
		}
	}
}

//the handle connection function takes the chatserver and telnet connection attempt as input and handles them appropriately.  The function uses a bufio scanner to check for username input and appropriately adds the user to

func handleConn(chatServer *ChatServer, conn net.Conn) {
	defer conn.Close()

	//grab the username from the user
	io.WriteString(conn, "Welcome to Torbit Chat!  Please, enter your username: ")

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	user := User{
		Name:   scanner.Text(),
		Output: make(chan Message, 10),
	}

	//join the user
	chatServer.Join <- user
	defer func() {
		chatServer.Leave <- user
	}()

	// Read message from connection and send the text to the chat server.
	go func() {
		for scanner.Scan() {
			ln := scanner.Text()
			chatServer.Input <- Message{user.Name, ln}
		}
	}()

	// everytime get a message from the chat server,  write it to the connection.
	for msg := range user.Output {
		// if msg.Username != user.Name {
		t := time.Now().Format("2006-01-02 15:04:05")
		log.Info(msg.Username + " said " + msg.Text)
		_, err := io.WriteString(conn, msg.Username+" "+t+": "+msg.Text+"\n")
		if err != nil {
			log.Error("Non fatal error writing chat message from server to terminal output.")
			break
		}
		// }
	}
}

//main function sets logging file parameters, sets server to listen on default telnet port 23 using TCP, creates and starts chat server,listens for attempted connections and passes connections to handleConn go routine upon success.

func main() {
	config, _ := LoadConfiguration("config.json")
	//standardizes log data into JSON format
	var filename string = "chatLog.log"

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	log.SetLevel(log.FatalLevel)
	log.SetFormatter(&log.JSONFormatter{})

	server, err := net.Listen(config.Server.Protocol, config.Server.Port)
	if err != nil {
		log.Fatal("Fatal error listening for telnet TCP traffic on " + config.Server.Port)
	} else {
		log.SetOutput(f)
	}

	defer server.Close()

	chatServer := &ChatServer{
		Users: make(map[string]User),
		Join:  make(chan User),
		Leave: make(chan User),
		Input: make(chan Message),
	}
	go chatServer.Run()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleConn(chatServer, conn)
	}
}
