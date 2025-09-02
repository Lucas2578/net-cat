package utils

import (
	"net"
	"regexp"
	"sync"
)

var (
	Users               []net.Conn
	UserNames           []string
	Mutex               sync.Mutex
	LogsCurrentMessages map[string][]string
	Canal               []string
	UsersCanal          map[net.Conn]string
)

func InitMaps() {
	UsersCanal = make(map[net.Conn]string)
	LogsCurrentMessages = make(map[string][]string)
}

var url string = "./assets/"
var fileNameLogs string = "logs.txt"

var regexRemoveAnsi = regexp.MustCompile(`\x1b\[[0-9;?]*[a-zA-Z]`)

var Black string = "\033[30m"
var Red string = "\033[31m"
var Green string = "\033[32m"
var Yellow string = "\033[33m"
var Blue string = "\033[34m"
var Magenta string = "\033[35m"
var Cyan string = "\033[36m"
var White string = "\033[37m"
var Bold string = "\033[1m"
var Reset string = "\033[0m"

var Alias = map[string]string{
	"/changepseudo": "/changepseudo",
	"/cp":           "/changepseudo",
	"/nick":         "/changepseudo",
	"/pseudo":       "/changepseudo",
	"/help":         "/help",
	"/aide":         "/help",
	"/createcanal":  "/createcanal",
	"/canalcreate":  "/createcanal",
	"/newcanal":     "/createcanal",
	"/logout":       "/logout",
	"/disconnect":   "/logout",
	"/leavecanal":   "/leavecanal",
	"/quit":         "/leavecanal",
	"/changecanal":  "/leavecanal",
	"/canals":       "/canals",
	"/listcanal":    "/canals",
	"/users":        "/users",
	"/userslist":    "/users",
	"/online":       "/users",
}

var Help []string = []string{
	"/help | /aide: Show commands.\n",
	"/changepseudo | /cp | /pseudo | /nick: Change your username.\n",
	"/createcanal | /canalcreate | /newcanal: Create new canal.\n",
	"/logout | /disconnect: Disconnect.\n",
	"/leavecanal | /quit | /changecanal: Change our current canal.\n",
	"/canals | /listcanal: Display the list of canals availables.\n",
	"/users | /userslist | /online: Display the list of canals availables.\n",
}

var PrefixSystem = string(Bold + Yellow + "[NET-CAT] " + Reset)

var ErrorUsage = string(Bold + Red + "[USAGE]:" + Reset + Red + " ./TCPChat $port" + Reset)
var ErrorListener = string(Red + "Error to accept listener" + Reset)
var ErrorCanalNotExist = string(PrefixSystem + Red + "This canal not exist ! You can create with /canalcreate\n" + Reset)
var ErrorCreateCanalAlreadyExist = string(PrefixSystem + Red + "This canal already exist !" + Reset)
var MaxUsersReached = string(PrefixSystem + Red + "The max users reached (10).\nPlease wait for someone disconnect on this chat." + Reset)
var ErrorUserNameAlreadyExist = string(PrefixSystem + Red + "Username already taken.\n" + Reset)
var ErrorFieldIsEmpty = string(PrefixSystem + Red + "Field is empty !\n" + Reset)
var ErrorDataIsTooLong = string(PrefixSystem + Red + "Data is too long. Max 16 characters !\n" + Reset)

var MessageFormatted = string(Black + "%s" + Reset + " | " + Magenta + "[%s]" + Reset + Bold + White + " %s" + Reset + ": %s\n")

var SuccessStart = string(Green + "Listening on port %s\n" + Reset)
var SuccessUserDisconnectCanal = string(PrefixSystem + Green + "You have been disconnected on" + Reset + Bold + Blue + " %s " + Reset + Green + "canal.\n" + Reset)
var SuccessUserJoinCanal = string(PrefixSystem + Green + "You have been joined %s canal.\n" + Reset)
var SuccessNewCanal = string(PrefixSystem + Green + "You are on" + Reset + Bold + Blue + " %s " + Reset + Green + "canal now. Talk to others users !\n" + Reset)
var SuccessNewCanalServer = string(PrefixSystem + Green + "User" + Reset + Bold + Blue + " %s " + Reset + Green + "created new canal:" + Reset + Bold + Blue + " %s " + Reset + Green + "!\n" + Reset)
var DisplayAllCanal = string(PrefixSystem + Green + "Canal available: \n" + Reset)
var DisplayOneCanal = string("- " + Magenta + "%s\n" + Reset)
var DisplayAllUsers = string(PrefixSystem + Green + "Users online: \n" + Reset)
var DisplayOneUser = string("- " + Magenta + "%s\n" + Reset)

var EnterCanal = string(PrefixSystem + Cyan + "Enter your canal: " + Reset)
var EnterNewCanal = string(PrefixSystem + Cyan + "Name your new canal: " + Reset)
var EnterPseudo = string(PrefixSystem + Cyan + "Enter your name: " + Reset)
var EnterNewPseudo = string(PrefixSystem + Cyan + "Enter your new name: " + Reset)

var NotifyChangePseudo = string(PrefixSystem + Blue + "[%s] \"%s\" is now \"%s\"\n" + Reset)
var NotifyUserEnter = string(PrefixSystem + Bold + Blue + "%s" + Reset + Blue + " has joined our canal...\n" + Reset)
var NotifyServerEnter = string(PrefixSystem + Bold + Blue + "%s" + Reset + Blue + " has joined" + Reset + Bold + Blue + " %s " + Reset + Blue + "canal...\n" + Reset)
var NotifyUserLeave = string(PrefixSystem + Bold + Blue + "%s" + Reset + Blue + " has left our canal...\n" + Reset)
var NotifyServerLeave = string(PrefixSystem + Bold + Blue + "%s" + Reset + Blue + " has left" + Reset + Bold + Blue + " %s " + Reset + Blue + "canal...\n" + Reset)
var NotifyUserLogout = string(PrefixSystem + Bold + Red + "%s" + Reset + Yellow + " has disconnected...\n" + Reset)

var Welcome = string(Magenta + " Welcome to TCP-Chat!\n" + Reset + Yellow +
	"          _nnnn_\n" +
	"         dGGGGMMb\n" +
	"        @p~qp~~qMb\n" +
	"        M|@||@) M|\n" +
	"        @,----.JM|\n" +
	"       JS^\\__/  qKL\n" +
	"      dZP        qKRb\n" +
	"     dZP          qKKb\n" +
	"    fZP            SMMb\n" +
	"    HZM            MMMM\n" +
	"    FqM            MMMM\n" +
	"  __| \".        |\\dS\"qML\n" +
	"  |    `.       | `' \\Zq\n" +
	" _)      \\.___.,|     .'\n" +
	" \\____   )MMMMMP|   .'\n" +
	"      `-'       `--'\n" + Reset)
