# slack-summary



## Description
`slack-summary` provides a way to send mail, posting of one day of a specific channel
![screenshot](https://raw.github.com/wiki/futabooo/slack-summary/image/screenshot_2016-05-03_16_58_50.png)

## Usage

```
$ slack-summary -h                                                                                                                                                (git)-[master]
Usage of slack-summary:
  -c string
        slack-summary config file path.(Short)
  -conf string
        slack-summary config file path.
  -d string
        Day of history get.(Short) (default "2016/05/04")
  -date string
        Day of history get. (default "2016/05/04")
  -v    Print version information and quit.(Short)
  -version
        Print version information and quit.
```

Fill the item of toml file

```bash
[slack-info]
token = ""
channel_id = ""

[google-account]
name = ""
pass = ""
```
Create executable file ＆ execute

```bash
$ cd path/to/slack-summary
$ go build
$ ./slack-summary -c path/to/conf.toml
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/futabooo/slack-summary
$ cd path/to/slack-sumary/src
$ glide install
```

## Contribution

1. Fork ([https://github.com/futabooo/slack-summary/fork](https://github.com/futabooo/slack-summary/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[futabooo](https://github.com/futabooo)
