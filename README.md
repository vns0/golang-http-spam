# golang-http-spam

Example HTTP SPAM code written in GoLang.

Problem: Our project often has phishing attacks that have to be attacked by a simple HTTP attack to overflow the 
database/nginx/cache/etc.

The idea of the project is to give a tool to everyone to clean the market from such stupid scammers

## Setup
Clone the repository and change the working directory:

    git clone https://github.com/nikitavoryet/golang-http-spam.git
    cd golang-http-spam

Build and run the program:

    go build -o attack
    ./attack -url localhost:8080

## Usage
    ./main -url string [-method string POST | GET] [-data string JSON] [-count int]

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

# In the plans :
```
- [ ] GET with query
- [X] add proxy for attack
- [X] use go rutina
- [ ] mass urls attack
```
```
    author: 
    
    Name:          Nikita
    Company:       OG1
    Position:      TeamLead
    Mail:          n.vtorushin@inbox.ru
    TG:            @nikitavoryet
    Year of birth: 1999
    FullStack:     JS/GO
