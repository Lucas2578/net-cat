package server

import (
	"bufio"
	"fmt"
	"net"
	"net-cat/utils"
	"os"
	"strings"
)

func Start(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(utils.ErrorUsage)
		os.Exit(1)
	} else {
		fmt.Printf(utils.SuccessStart, port)
	}

	defer listener.Close()

	utils.InitMaps()

	utils.Mutex.Lock()
	utils.Canal = append(utils.Canal, "general")
	utils.Mutex.Unlock()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(utils.ErrorListener)
			os.Exit(1)
		}

		utils.Mutex.Lock()
		utils.Users = append(utils.Users, conn)
		utils.Mutex.Unlock()

		if len(utils.Users) > 10 {
			conn.Write([]byte(utils.MaxUsersReached))

			utils.Mutex.Lock()
			for i, u := range utils.Users {
				if u == conn {
					utils.Users = append(utils.Users[:i], utils.Users[i+1:]...)
					break
				}
			}

			conn.Close()
			utils.Mutex.Unlock()
			continue
		}

		utils.WelcomeMessage(conn)

		go AnalyzeText(conn)
	}
}

func AnalyzeText(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	userName, oldUserName := ChooseUsername(conn, scanner, "")
	canal := ""
	utils.AdjustTerminal(conn)

	if userName != "" {
		canal = Log(conn, userName, "login", scanner)
		utils.NotifyLog(conn, userName, "login", canal)
	}

	defer func() {
		Log(conn, userName, "logout", scanner)
		if canal == "" {
			canal = "empty"
			utils.NotifyLog(conn, userName, "disconnect", canal)
		} else {
			utils.NotifyLog(conn, userName, "logout", canal)
		}
	}()

	for scanner.Scan() {
		hasChangePseudo := false
		line := scanner.Text()
		if cmd, ok := utils.Alias[strings.ToLower(line)]; ok {
			line = cmd
		}
		switch line {
		case "":
			continue
		case "/changepseudo":
			userName, oldUserName = ChooseUsername(conn, scanner, userName)
			hasChangePseudo = true
		case "/help":
			utils.HelpMessage(conn)
			continue
		case "/createcanal":
			canal = CreateCanal(conn, scanner, canal, userName)
			utils.NotifyLog(conn, userName, "login", canal)
			continue
		case "/logout":
			conn.Close()
			return
		case "/leavecanal":
			utils.NotifyLog(conn, userName, "logout", canal)
			message := fmt.Sprintf(utils.SuccessUserDisconnectCanal, canal)
			conn.Write([]byte(message))
			canal = ChooseCanal(conn, scanner)
			utils.NotifyLog(conn, userName, "login", canal)
			conn.Write([]byte("\n"))
			continue
		case "/canals":
			utils.DisplayCanal(conn)
			continue
		case "/users":
			utils.DisplayUsers(conn)
			continue
		}
		message := utils.CreateMessage(canal, userName, oldUserName, line, hasChangePseudo)
		fmt.Print(message)
		utils.BroadcastMessage(conn, message, canal)
	}
}
