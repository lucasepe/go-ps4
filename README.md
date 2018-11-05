# Go PS4

> Search the Playstation Store for your favorite PS4 games using the command line.

This tool use [Colly](https://github.com/gocolly/colly) the Lightning Fast and Elegant Scraping Framework for Gophers

## How to build

### Windows

```shell
go build -ldflags "-X main.version=0.0.1 -X main.date=%date:~-4%-%date:~3,2%-%date:~6,2%T%time:~0,2%:%time:~3,2%:%time:~6,2%"
```

### Linux

```shell
go build -ldflags "-X main.version=0.0.1 -X main.date=%date:~10,4%-%date:~4,2%-%date:~7,2%T%time:~0,2%:%time:~3,2%:%time:~6,2%"
```

## How to use

### List available options

```shell
go-ps4 -help
```

```shell
Usage of go-ps4:
  -addons
        show also extra contents
  -free
        show only free titles
  -lang string
        language (it, en) (default "it")
  -search string
        search for specified title
  -version
        show app version
  -weekly-deals
        show only weekly deals titles
```

### Search only for free games

```shell
go-ps4 -free
```

### Search for a specific game

```shell
go-ps4 -search "strange brigade"
```

### Show only weekly deals

```shell
go-ps4 -weekly-deals
```

