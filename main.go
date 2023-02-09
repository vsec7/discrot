package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func getMe(b string, t string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", b+"/users/@me", nil)
	req.Header.Set("authorization", t)
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var result Resp
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON Error")
	}

	return result.Username + "#" + result.Discriminator
}

func getMessage(b string, t string, cid string, l string) (string, string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", b+"/channels/"+cid+"/messages?limit="+l, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", t)
	res, err := client.Do(req)
	if err != nil {
		return "", ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Ele
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON Error")
	}
	d, _ := strconv.Atoi(l)
	x := result[d-1]
	return x.Id, x.Content
}

func sendMessage(b string, t string, cid string, txt string) (string, string) {
	client := &http.Client{}
	data := []byte(`{"content":"` + txt + `"}`)
	req, _ := http.NewRequest("POST", b+"/channels/"+cid+"/messages", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", t)
	res, err := client.Do(req)
	if err != nil {
		return "", ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Resp
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON Error")
	}

	return result.Id, result.Content
}

func sendReply(b string, t string, cid string, mid string, txt string) (string, string) {
	client := &http.Client{}
	data := []byte(`{"content":"` + txt + `", "message_reference":{ "message_id": "` + mid + `"} }`)
	req, _ := http.NewRequest("POST", b+"/channels/"+cid+"/messages", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", t)
	res, err := client.Do(req)
	if err != nil {
		return "", ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Resp
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON Error")
	}

	return result.Id, result.Content
}

func deleteMessage(b string, t string, cid string, mid string) {
	client := &http.Client{}
	req, _ := http.NewRequest("DELETE", b+"/channels/"+cid+"/messages/"+mid, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", t)
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
}

func simsimi(lc string, txt string) string {
	client := &http.Client{}
	data := url.Values{}
	data.Set("lc", lc)
	data.Set("text", txt)
	req, _ := http.NewRequest("POST", "https://api.simsimi.vn/v1/simtalk", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Resp
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON Error")
	}

	return result.Msg
}

func rand_quote() string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://gist.githubusercontent.com/vsec7/45b5b362d9dde8bd381dc9b8a3b6331a/raw/7934009fc02655de3ffc4a28af22a6c0c1ecd102/quote-id.json", nil)
	res, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result Ele
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("JSON Error")
	}

	x := result[rand.Intn(len(result))]
	return x.Quote
}

func rand_custom_text(f string) string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines[rand.Intn(len(lines))]
}

func init() {
	flag.StringVar(&ConfigFile, "conf", "config.yaml", "Configuration File")
	flag.StringVar(&Channel_id, "c", "", "Channel ID")
	flag.IntVar(&Delay, "d", 60, "Delay *Seconds")
	flag.StringVar(&Mode, "m", "repost", "Mode")
	flag.StringVar(&Custom, "custom", "custom.txt", "Custom.txt File")
	flag.StringVar(&Last, "l", "50", "Limit Chat History")
	flag.BoolVar(&Del, "del", false, "Delete After Sending")
	flag.StringVar(&Lc, "lc", "id", "Simsimi Language * id/en")
	flag.BoolVar(&Reply, "reply", false, "Reply *For Simsimi Only")
	flag.Usage = func() {
		h := []string{
			"",
			"Discrot (Discord Selfbot)",
			"",
			"Is a simple GO tool for Grinding / Leveling chat discord",
			"",
			"Coded By : github.com/vsec7",
			"",
			"Basic Usage :",
			" ▶ discrot -m quote -d 30",
			"Advanced Usage :",
			" ▶ discrot --conf config.yaml -c 1234567890 -m simsimi -d 30 --reply -del",
			"",
			"Options :",
			"  -conf, --conf <config.yaml>	Set file config.yaml (default: config.yaml) *required",
			"  -c, --c <channel_id>	        Set channel_id *required",
			"  -m, --m <mode>         	Set Mode: repost, quote, simsimi, custom (default: repost)",
			"  -d, --d <delay>         	Set Delay *Seconds (default: 60)",
			"  -l, --l <limit>         	Set Limit last chat history (default: 50)",
			"  -del, --del         		Set Delete after sending message",
			"  -reply, --reply         	Set Reply *For Simsimi Only",
			"  -lc, --lc <id/en>         	Set Simsimi Language *For Simsimi Only (default: id)",
			"  -custom, --custom <file.txt>	Set custom file *required if set custom mode (default: custom.txt)",
			"",
			"",
		}
		fmt.Fprintf(os.Stderr, strings.Join(h, "\n"))
	}
	flag.Parse()
}

var (
	ConfigFile string
	Channel_id string
	Delay      int
	Mode       string
	Last       string
	Custom     string
	Del        bool
	Reply      bool
	Loop       bool
	Lc         string
	config     Config
)

type Config struct {
	Token []string `yaml:"BOT_TOKEN"`
	Chan  []string `yaml:"CHANNEL_ID"`
}

type Resp struct {
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Id            string `json:"id"`
	Content       string `json:"content"`
	Msg           string `json:"message"`
}

type Ele []struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	Quote   string `json:"quote"`
}

func main() {

	b := "https://discord.com/api/v9"
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		fmt.Printf("[ERROR] File %s not found!\n", ConfigFile)
		os.Exit(0)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("[ERROR] File %s could not parsing !\n", ConfigFile)
		os.Exit(0)
	}

	Loop := true

	for Loop {

		for _, t := range config.Token {
			me := getMe(b, t)
			if me == "#" {
				fmt.Printf("[INVALID TOKEN] %s\n", t)
			} else {

				var cid []string
				if len(Channel_id) != 0 {
					cid = append(cid, Channel_id)
				}
				if len(config.Chan) != 0 {
					cid = append(cid, config.Chan...)
				}

				for _, c := range cid {

					switch mode := Mode; mode {
					case "quote":
						q := rand_quote()
						id, msg := sendMessage(b, t, c, q)
						if len(id) != 0 {
							fmt.Printf("[%s][QUOTE][%s] %s\n", me, id, msg)
						} else {
							fmt.Printf("[%s][QUOTE][FAILED] Maybe muted / kicked from Channel: %s\n", me, c)
						}

						if Del {
							fmt.Printf("[%s][QUOTE][DELETE][%s] %s\n", me, id, msg)
							deleteMessage(b, t, c, id)
						}
					case "repost":

						_, m := getMessage(b, t, c, Last)

						id, msg := sendMessage(b, t, c, m)
						if len(id) != 0 {
							fmt.Printf("[%s][REPOST][%s] %s\n", me, id, msg)
						} else {
							fmt.Printf("[%s][REPOST][FAILED] Maybe muted / kicked from Channel: %s\n", me, c)
						}

						if Del {
							fmt.Printf("[%s][REPOST][DELETE][%s] %s\n", me, id, msg)
							deleteMessage(b, t, c, id)
						}
					case "simsimi":

						if Reply {
							i, m := getMessage(b, t, c, Last)

							simi := simsimi(Lc, m)
							id, msg := sendReply(b, t, c, i, simi)
							if len(id) != 0 {
								fmt.Printf("[%s][SIMSIMI][REPLY][%s] %s\n", me, id, msg)
							} else {
								fmt.Printf("[%s][SIMSIMI][REPLY][FAILED] Maybe muted / kicked from Channel: %s\n", me, c)
							}

							if Del {
								fmt.Printf("[%s][SIMSIMI][REPLY][DELETE][%s] %s\n", me, id, msg)
								deleteMessage(b, t, c, id)
							}
						} else {
							_, m := getMessage(b, t, c, Last)

							simi := simsimi(Lc, m)
							id, msg := sendMessage(b, t, c, simi)
							if len(id) != 0 {
								fmt.Printf("[%s][SIMSIMI][%s] %s\n", me, id, msg)
							} else {
								fmt.Printf("[%s][SIMSIMI][FAILED] Maybe muted / kicked from Channel: %s\n", me, c)
							}

							if Del {
								fmt.Printf("[%s][SIMSIMI][DELETE][%s] %s\n", me, id, msg)
								deleteMessage(b, t, c, id)
							}
						}

					case "custom":

						txt := rand_custom_text(Custom)
						id, msg := sendMessage(b, t, c, txt)
						if len(id) != 0 {
							fmt.Printf("[%s][CUSTOM][%s] %s\n", me, id, msg)
						} else {
							fmt.Printf("[%s][CUSTOM][FAILED] Maybe muted / kicked from Channel: %s\n", me, c)
						}

						if Del {
							fmt.Printf("[%s][CUSTOM][DELETE][%s] %s\n", me, id, msg)
							deleteMessage(b, t, c, id)
						}

					}

				}
			}
		}

		fmt.Printf("----------[ Delay for %d Seconds]----------\n", Delay)
		time.Sleep(time.Duration(Delay) * time.Second)
	}
}
