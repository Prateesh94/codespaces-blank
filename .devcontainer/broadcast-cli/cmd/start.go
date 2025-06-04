/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var mu sync.RWMutex
var upgrader = websocket.Upgrader{

	CheckOrigin: func(r *http.Request) bool {

		return true
	},
}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)
var i int = 0
var count int = 0

func handleconnection(w http.ResponseWriter, r *http.Request) {
	ws, er := upgrader.Upgrade(w, r, nil)
	if er != nil {
		fmt.Fprintf(w, "%v", er)
		return
	}
	defer ws.Close()
	mu.Lock()
	clients[ws] = true
	count++
	i++
	mu.Unlock()
	fmt.Printf("A client has connected, Total Clients connected:- %d\n", i)
	for {

		_, msg, err := ws.ReadMessage()

		if err != nil {
			mu.Lock()
			delete(clients, ws)
			i--
			mu.Unlock()
			fmt.Printf("A client has disconnected, Total Clients left:-  %d\n", i)
			fmt.Printf("Total Clients served: %d\n", count)
			break
		}
		broadcast <- msg

	}
}
func handleMessages() {
	for {

		msg := <-broadcast
		mu.RLock()
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				client.Close()
				delete(clients, client)
			}
		}
		mu.RUnlock()
	}
}

// startCmd represents the start command
func NewServer(s string) {
	router := mux.NewRouter()
	router.HandleFunc("/ws", handleconnection)
	fmt.Println("Server started on :" + s)
	fmt.Println("Connect to ws://localhost:" + s + "/ws")
	go handleMessages()
	log.Fatal(http.ListenAndServe(":"+s, router))
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start comand is used to start a broadcast server at specified port, if no port is mentioned the default port 8080 will be used",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var i int
		var er error
		if len(args) == 0 {
			i = 8080
		} else {
			i, er = strconv.Atoi(args[0])
			fmt.Println(i)
			if er != nil || i < 1024 || i > 65535 {
				fmt.Println("Port number should be between 1024 and 65535")
				return
			}
		}
		if i == 0 {
			i = 8080
		}
		NewServer(strconv.Itoa(i))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
