package main

import (
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	OutStream, ErrStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		confPath string
		date     string
		version  bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.ErrStream)
	flags.StringVar(&confPath, "conf", "", "slack-summary config file path.")
	flags.StringVar(&confPath, "c", "", "slack-summary config file path.(Short)")
	flags.StringVar(&date, "date", today(), "Day of history get.")
	flags.StringVar(&date, "d", today(), "Day of history get.(Short)")
	flags.BoolVar(&version, "version", false, "Print version information and quit.")
	flags.BoolVar(&version, "v", false, "Print version information and quit.(Short)")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.ErrStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	if confPath == "" {
		fmt.Fprint(cli.ErrStream, "conf file is required\n")
		return ExitCodeError
	}

	conf, err := LoadConfToml(confPath)
	if err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	if err := doSummarySend(conf, date); err != nil {
		fmt.Println(err)
		return ExitCodeError
	}

	_ = conf
	_ = date

	return ExitCodeOK
}

func today() string {
	return time.Now().Format("2006/01/02")
}

func doSummarySend(conf ConfToml, date string) error {
	api := slack.New(conf.SlackInfo.Token)
	history, err := api.GetChannelHistory(conf.SlackInfo.ChannelID, getOneDayHistoryParam(date))
	if err != nil {
		fmt.Println(err)
		return err
	}

	var mailBody string
	for _, message := range history.Messages {
		mailBody += fmt.Sprintln(unixNanoToTimeStr(message.Timestamp))
		mailBody += fmt.Sprintln(message.Text)
	}

	SendMail(conf, mailBody)

	return err
}

func getOneDayHistoryParam(date string) slack.HistoryParameters {
	param := slack.NewHistoryParameters()
	ts, _ := time.ParseInLocation("2006/01/02", date, time.Local)
	te := ts.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
	param.Oldest = strconv.FormatInt(ts.Unix(), 10)
	param.Latest = strconv.FormatInt(te.Unix(), 10)
	return param
}

func unixNanoToTimeStr(unixNano string) string {
	s := strings.Split(unixNano, ".")
	i, _ := strconv.ParseInt(s[0], 10, 64)
	t := time.Unix(i, 0)
	return fmt.Sprintf("[%02d:%02d:%02d]", t.Hour(), t.Minute(), t.Second())
}
