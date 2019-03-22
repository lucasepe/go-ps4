# Go PS4

> Search the Playstation Store for your favorite PS4 games using the command line.

This tool use [Colly](https://github.com/gocolly/colly) the Lightning Fast and Elegant Scraping Framework for Gophers


## How to use

### List available options

```shell
$ go-ps4 -h
```

```shell
$ PS Store CLI (v.0.5.0)
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
$ go-ps4 -free
3on3 Freestyle...................................................... Gratuito
AirMech® Arena...................................................... Gratuito
A KING'S TALE: FINAL FANTASY XV..................................... Gratuito
anywhereVR.......................................................... Gratuito
APB Reloaded........................................................ Gratuito
...(more games)....
```

### Search for a specific game

```shell
$ go-ps4 "division 2"
Tom Clancy's The Division® 2 - Standard Edition..................... €69,99
Tom Clancy’s The Division®2 - Pacchetto di benvenuto................ €14,99
Pacchetto Valuta premium 500 per Tom Clancy's The Division®2........ €4,99
Pacchetto Valuta premium 1050 per Tom Clancy's The Division®2....... €9,99
Tom Clancy's The Division® 2 - Ultimate Edition..................... €119,99
Pacchetto Valuta premium 6500 per Tom Clancy's The Division®2....... €49,99
```

### Show only weekly deals

```shell
$ go-ps4 -weekly-deals
```

