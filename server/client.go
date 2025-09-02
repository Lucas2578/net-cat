package server

import (
	"bufio"
	"net"
	"net-cat/utils"
	"strings"
)

func ChooseUsername(conn net.Conn, scanner *bufio.Scanner, userName string) (string, string) {

	for {
		oldUserName := ""
		if userName != "" {
			oldUserName = userName
			conn.Write([]byte(utils.EnterNewPseudo))
			Log(conn, userName, "changepseudo", scanner)
		} else {
			conn.Write([]byte(utils.EnterPseudo))
		}
		exist := false

		if scanner.Scan() {
			userName := strings.TrimSpace(scanner.Text())
			if utils.VerifyIsEmpty(userName) {
				conn.Write([]byte(utils.ErrorFieldIsEmpty))
				continue
			}
			if utils.VerifyIsTooLong(userName) {
				conn.Write([]byte(utils.ErrorDataIsTooLong))
				continue
			}
			for _, user := range utils.UserNames {
				if user == userName {
					exist = true
					break
				}
			}

			if exist {
				conn.Write([]byte(utils.ErrorUserNameAlreadyExist))
				continue
			}

			if userName != "" {
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
				return userName, oldUserName
			}
		} else {
			return "", ""
		}
	}
}
