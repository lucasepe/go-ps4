# Go PS4

[![Go Report Card](https://goreportcard.com/badge/github.com/lucasepe/go-ps4)](https://goreportcard.com/report/github.com/lucasepe/go-ps4) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/gojp/goreportcard/blob/master/LICENSE)

> Search the Playstation Store for your favorite PS4 games using the command line.

This tool use [Colly](https://github.com/gocolly/colly) the Lightning Fast and Elegant Scraping Framework for Gophers


- Free [open source](https://github.com/lucasepe/go-ps4) software
- Works on [Linux](https://github.com/lucasepe/go-ps4/releases/download/v1.0.0/ps4-linux-amd64), [Mac OSX](https://github.com/lucasepe/go-ps4/releases/download/v1.0.0/ps4-darwin-amd64), [Windows](https://github.com/lucasepe/go-ps4/releases/download/v1.0.0/ps4-windows-amd64.exe)
- Just a single portable binary file


## How to use

### List available options

```shell
$ ps4 -h
PS Store CLI (v.0.5.0)
Usage: go-ps4 [OPTIONS] [Game Title]
  -addons
        show also extra contents
  -free
        show only free titles
  -lang string
        language (it, en) (default "it")
  -weekly-deals
        show only weekly deals titles
```

### Search only for free games

```shell
$ ps4 -free
3on3 Freestyle...................................................... Gratuito
AirMech® Arena...................................................... Gratuito
A KING'S TALE: FINAL FANTASY XV..................................... Gratuito
anywhereVR.......................................................... Gratuito
APB Reloaded........................................................ Gratuito
...(more games)....
```

### Search for a specific game

```shell
$ ps4 "division 2"
Tom Clancy's The Division® 2 - Standard Edition..................... €69,99
Tom Clancy’s The Division®2 - Pacchetto di benvenuto................ €14,99
Pacchetto Valuta premium 500 per Tom Clancy's The Division®2........ €4,99
Pacchetto Valuta premium 1050 per Tom Clancy's The Division®2....... €9,99
Tom Clancy's The Division® 2 - Ultimate Edition..................... €119,99
Pacchetto Valuta premium 6500 per Tom Clancy's The Division®2....... €49,99
```

### Show only weekly deals

```shell
$ ps4 -weekly-deals
```

