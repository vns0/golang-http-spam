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
          method for attack (POST/GET). def: POST
    - data string
          JSON string for body/query to HTTP requst ({"email" : "n.vtorushin@inbox.ru", "test": "test"})
    - count int
          count requst for attack. Default: 10000

# In the plans :
```
- [ ] GET with query
- [ ] add proxy for attack
- [ ] use go rutina
- [ ] mass urls attack
```

```
Donate:

    BTC:  192TC7d7ZRYJQbQnAWvMpkccnBNQN1ae6R
    ETH:  0x7d1082d952f4d584ae2910e14018f4dce7495c74
    LTC:  MLx6wmFjXfBTKj6JfB5NXaiKjNLeEntRoZ
    DOGE: DHCjW71EWBzvv43XPXVJc491brcBJXXq88
```
    author: 
    
    Name:          Nikita
    Company:       SmartWorld
    Position:      TeamLead
    Mail:          n.vtorushin@inbox.ru
    TG:            @nikitavoryet
    Year of birth: 1999
    FullStack:     JS/GO