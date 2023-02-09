# discrot
Discrot is a simple GO tool for Grinding / Leveling chat discord

Created By viloid (github.com/vsec7)

> *** NOTE : USE AT YOUR OWN RISK! ***

> DONT SELL THIS SCRIPT! YOU'RE REALLY FUCKING POOR STUPID DOG!

```bash
Discrot (Discord Selfbot)

Is a simple GO tool for Grinding / Leveling chat discord

Coded By : github.com/vsec7

Basic Usage :
 ▶ discrot -m quote -d 30
Advanced Usage :
 ▶ discrot --conf config.yaml -c 1234567890 -m simsimi -d 30 --reply -del

Options :
  -conf, --conf <config.yaml>   Set file config.yaml (default: config.yaml) *required
  -c, --c <channel_id>          Set channel_id *required
  -m, --m <mode>                Set Mode: repost, quote, simsimi, custom (default: repost)
  -d, --d <delay>               Set Delay *Seconds (default: 60)
  -l, --l <limit>               Set Limit last chat history (default: 50)
  -del, --del                   Set Delete after sending message
  -reply, --reply               Set Reply *For Simsimi Only
  -lc, --lc <id/en>             Set Simsimi Language *For Simsimi Only (default: id)
  -custom, --custom <file.txt>  Set custom file *required if set custom mode (default: custom.txt)

```

## • Features
- Send Quote message
- Send Response Simsimi message
- Send Repost message from channel chat history
- Send select random line from custom.txt
- Auto Delete message

## • Requirement
> go version: go1.18+ 

## • Installation
```
go install -v github.com/vsec7/discrot@latest
```

## • Configuration Template
```yaml
BOT_TOKEN: 
 - ODE3MzkzOTQ4MjM4xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
 - ODE3MzkzOTQ4MjM4xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
CHANNEL_ID: 
 - 103009xxxxxxxxxxxxxxx
 - 103009xxxxxxxxxxxxxxx
```

## • How to get Discord SelfBot Token?

```
javascript:(()=>{var t=document.body.appendChild(document.createElement`iframe`).contentWindow.localStorage.token.replace(/["]+/g, '');prompt('Get Selfbot Discord Token by github.com/vsec7', t)})();
```

[<kbd>DETAILS CLICK HERE</kbd>](https://gist.github.com/vsec7/12066af3f704bd337c52c30f4c492ba2)

## • Donate

SOL Address : viloid.sol

BSC Address : 0xd3de361b186cc2Fc0C77764E30103F104a6d6D07
