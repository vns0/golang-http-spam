# golang-http-spam
## This project is a demonstration project. The developer is not responsible for the use of this repository
Example HTTP SPAM code written in GoLang.

Problem: Our project often has phishing attacks that have to be attacked by a simple HTTP attack to overflow the 
database/nginx/cache/etc.

The idea of the project is to give a tool to everyone to clean the market from such stupid scammers

## Setup
Clone the repository and change the working directory:

    git clone https://github.com/nikitavoryet/golang-http-spam.git
    cd golang-http-spam

Build the program:

Linux:

    go build -o attack

Windows:

    go build ./main.go

Run the program:

Linux:

    ./attack -url localhost:8080

Windows:

    ./main.exe -url localhost:8080

## Usage
    ./main

    Usage of ./main:
    -url string
          url for attack with http/https (https://example.com/v1/auth)
    -method string
          method for attack (POST/GET). Default: POST
    -data string
          JSON string for body/query to HTTP requst ({"email" : "n.vtorushin@inbox.ru", "test": "test"})
    -count int
          count requst for attack. Default: 10000 
    -threads int
          count threads for attack. Default: 1
    -proxy string
          path to proxy list. Accept: HTTP(S)/SOCKS(4/5)
    -query string
          query params for URL. Accept: GET/POST request
    -urlPath string
          path to url list. formula: threads * urlPath

    #If u want work with telegram bot replace global var botToken#
        bot supports command /spamm with all arguments above'

## Api auth:
In order to pass authorization, you need to call the /auth command and enter the received code on the website

# In the plans :
```
- [X] GET with query
- [X] add proxy for attack
- [X] use go rutina
- [X] add telegram bot
    - [ ] add stop command
    - [X] add createAdmin command
    - [X] add deleteAdmin command
- [X] mass urls attack
- [X] add db
    - [X] add history table
    - [X] add telegram admins list
    - [ ] add cmd for check history
- [ ] mass urls attack in telegram bot
- [X] add web interface
    - [X] create auth
    - [ ] create dashboard
    - [ ] create admin pannel
    - [ ] create module for start attack in web
```
```
    author: 

    Name:          Nikita
    Mail:          n.vtorushin@inbox.ru
    TG:            @nvtorushin
    About me:      https://vns.guru
