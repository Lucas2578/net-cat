package utils

import (
	"fmt"
	"net"
	"time"
)

func AdjustTerminal(conn net.Conn) {
	fmt.Fprint(conn, "\033[H\033[2J")
}

func WelcomeMessage(conn net.Conn) {
	conn.Write([]byte(Welcome))
}

func CreateMessage(canal string, userName string, oldUserName string, line string, hasChangePseudo bool) string {
	date := time.Now()
	formattedDate := date.Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(MessageFormatted, formattedDate, canal, userName, line)
	if hasChangePseudo {
		message = fmt.Sprintf(NotifyChangePseudo, formattedDate, oldUserName, userName)
	}
	return message
}

func NotifyLog(conn net.Conn, userName string, log string, canal string) {
	Mutex.Lock()
	defer Mutex.Unlock()
	message := ""
	messageServer := ""
	if userName == "" {
		userName = "Stranger"
	}

	if canal == "" {
		return
	}

	switch log {
	case "login":
		messageServer = fmt.Sprintf(NotifyServerEnter, userName, canal)
		message = fmt.Sprintf(NotifyUserEnter, userName)
	case "logout":
		messageServer = fmt.Sprintf(NotifyServerLeave, userName, canal)
		message = fmt.Sprintf(NotifyUserLeave, userName)
	case "disconnect":
		messageServer = fmt.Sprintf(NotifyUserLogout, userName)
		message = fmt.Sprintf(NotifyUserLogout, userName)
	}
	WriteFile(fileNameLogs, messageServer)
	fmt.Print(messageServer)

	for _, user := range Users {
		userCanal, ok := UsersCanal[user]
		if ok && userCanal == canal {
			if user != conn {
				user.Write([]byte(message))
			}
		}
	}
}

func BroadcastMessage(conn net.Conn, message string, canal string) {
	Mutex.Lock()
	defer Mutex.Unlock()

	WriteFile(fileNameLogs, message)

	LogsCurrentMessages[canal] = append(LogsCurrentMessages[canal], message)

	for _, user := range Users {
		if user != conn {
			userCanal, ok := UsersCanal[user]
			if ok && userCanal == canal {
				user.Write([]byte(message))
			}
		}
	}
}

func HelpMessage(conn net.Conn) {
	for _, message := range Help {
		conn.Write([]byte(message))
	}
}

func DisplayCanal(conn net.Conn) {
	conn.Write([]byte(DisplayAllCanal))
	for _, canal := range Canal {
		message := fmt.Sprintf(DisplayOneCanal, canal)
		conn.Write([]byte(message))
	}
	conn.Write([]byte("\n"))
}

func DisplayUsers(conn net.Conn) {
	conn.Write([]byte(DisplayAllUsers))
	for _, user := range UserNames {
		message := fmt.Sprintf(DisplayOneUser, user)
		conn.Write([]byte(message))
	}
	conn.Write([]byte("\n"))
}
