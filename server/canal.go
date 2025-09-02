package server

import (
	"bufio"
	"fmt"
	"net"
	"net-cat/utils"
	"strings"
)

func ChooseCanal(conn net.Conn, scanner *bufio.Scanner) string {

	for {
		utils.DisplayCanal(conn)
		conn.Write([]byte(utils.EnterCanal))

		if scanner.Scan() {
			canal := strings.TrimSpace(scanner.Text())
			if utils.VerifyIsEmpty(canal) {
				conn.Write([]byte(utils.ErrorFieldIsEmpty))
				continue
			}
			if utils.VerifyIsTooLong(canal) {
				conn.Write([]byte(utils.ErrorDataIsTooLong))
				continue
			}
			for _, existCanal := range utils.Canal {
				if existCanal == canal {
					utils.UsersCanal[conn] = canal
					message := fmt.Sprintf(utils.SuccessUserJoinCanal, canal)
					conn.Write([]byte(message))
					return canal
				}
			}
			conn.Write([]byte(utils.ErrorCanalNotExist))
		} else {
			return ""
		}
	}
}

func CreateCanal(conn net.Conn, scanner *bufio.Scanner, oldCanal string, userName string) string {
	for {
		conn.Write([]byte(utils.EnterNewCanal))

		if scanner.Scan() {
			newCanal := strings.TrimSpace(scanner.Text())
			if utils.VerifyIsEmpty(newCanal) {
				conn.Write([]byte(utils.ErrorFieldIsEmpty))
				continue
			}
			if utils.VerifyIsTooLong(newCanal) {
				conn.Write([]byte(utils.ErrorDataIsTooLong))
				continue
			}
			for _, canal := range utils.Canal {
				if canal == newCanal {
					conn.Write([]byte(utils.ErrorCreateCanalAlreadyExist))
					break
				}
			}
			utils.NotifyLog(conn, userName, "logout", oldCanal)
			utils.Canal = append(utils.Canal, newCanal)
			utils.UsersCanal[conn] = newCanal
			message := fmt.Sprintf(utils.SuccessNewCanal, newCanal)
			conn.Write([]byte(message))
			message = fmt.Sprintf(utils.SuccessNewCanalServer, userName, newCanal)
			fmt.Print(message)
			return newCanal
		} else {
			return ""
		}
	}
}
