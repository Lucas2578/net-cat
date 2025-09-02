package server

import (
	"bufio"
	"net"
	"net-cat/utils"
)

func Log(conn net.Conn, userName string, log string, scanner *bufio.Scanner) string {
	canal := ""

	switch log {
	case "login":
		utils.Mutex.Lock()
		exist := false
		for _, user := range utils.UserNames {
			if user == userName {
				exist = true
			}
		}
		if !exist {
			utils.UserNames = append(utils.UserNames, userName)
		}
		utils.Mutex.Unlock()
		canal = ChooseCanal(conn, scanner)
		utils.AdjustTerminal(conn)
		for _, message := range utils.LogsCurrentMessages[canal] {
			conn.Write([]byte(message))
		}
		return canal
	case "logout":
		RemoveUserName(userName)
		for i, u := range utils.Users {
			if u == conn {
				utils.Users = append(utils.Users[:i], utils.Users[i+1:]...)
				break
			}
		}
		conn.Close()
	case "changepseudo":
		RemoveUserName(userName)
	}
	return ""
}

func RemoveUserName(userName string) {
	utils.Mutex.Lock()
	defer utils.Mutex.Unlock()

	for i, name := range utils.UserNames {
		if name == userName {
			utils.UserNames = append(utils.UserNames[:i], utils.UserNames[i+1:]...)
			break
		}
	}
}
