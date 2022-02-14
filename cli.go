package MyDbms

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"golang.org/x/term"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"time"
)

var userName string

func main() {
	hostUrl := flag.String("url", "wss://localhost:6543", "Specify the Host Full Url")
	user := flag.String("user", "default", "Name of the User")
	flag.Parse()
	passwordBytes, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Printf("Error:\nInput Error\n%s\n", err.Error())
		return
	}

	ws := websocket.Dialer{
		HandshakeTimeout: time.Second * 20,
	}
	Uri, err := url.Parse(*hostUrl)
	if err != nil {
		fmt.Printf("Error:\nCommand Processing Error\n -host flag must be a valid url\n")
		return
	}

	conn, res, err := ws.Dial(Uri.String(), http.Header{})
	if err != nil {
		fmt.Printf("Error:\nConnection Error\n%s\n", err.Error())
		return
	}
	if res.StatusCode != 200 {
		fmt.Printf("Error:\nServer Did Not Respond 200 Status Code\n")
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error:\nConnection Closing Error\n%s\n", err.Error())
		}
	}(conn)

	err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("AUTH \"%s\" \"%s\"", *user,
		string(passwordBytes))))
	if err != nil {
		fmt.Printf("Error:\nAuthentication Error\n%s\n", err.Error())
		return
	}

	_, messageBytes, err := conn.NextReader()
	if err != nil {
		fmt.Printf("Error:\nAuthentication Error\n%s\n", err.Error())
		return
	}

	message, err := ioutil.ReadAll(messageBytes)
	if err != nil {
		fmt.Printf("Error:\nAuthentication Error Could Not Parse Authentication Result\n%s\n", err.Error())
		return
	}
	fmt.Println(string(message))

	userName = *user

	fmt.Printf("%s> ", userName)
	takeInputs := func() {
		in := input()
		if in == ".quit" {
			err = conn.Close()
			if err != nil {
				fmt.Printf("\nError:\nConnection Closing Error\n%s\n", err.Error())
			}

			os.Exit(0)
		}

		err = conn.WriteMessage(websocket.TextMessage, []byte(in))
		if err != nil {
			fmt.Printf("\nError:\nSocket Error\n%s\n", err.Error())
		}

		_, messageBytes, err = conn.NextReader()
		if err != nil {
			fmt.Printf("Error:\nMassage Receiving Error\n%s\n", err.Error())
			return
		}

		message, err := ioutil.ReadAll(messageBytes)
		if err != nil {
			fmt.Printf("Error:\nAuthentication Error Could Not Parse Authentication Result\n%s\n", err.Error())
			return
		}
		fmt.Println(string(message))
	}

	takeInputs()

}
func input() string {
	var userInput string

	_, err := fmt.Scanf("%s", &userInput)
	if err != nil {
		fmt.Printf("Error:\n User Input Error\n%s\n", err.Error())
		userInput = ""
	}

	return userInput
}
