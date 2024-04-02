## Websocket Endpoint

Default Websocket endpoints:

Client | URL
:----: | :----:
Go     | ws://localhost:9888/websocket-subscribe 

### Websocket Extension Methods (Websocket-specific)

#### Method Overview

|  #   |            Topic             |                         Description                          |                       NotificationType                       |
| :--: | :--------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
|  1   |      notify_raw_blocks       | Send notifications when a block is connected or disconnected from the best chain. | raw_blocks_connected raw_blocks_disconnected  request_status |
|  2   |    stop_notify_raw_blocks    | Cancel registered notifications for whenever a block is connected or disconnected from the main (best) chain. |                             None                             |
|  3   |   notify_new_transactions    | Send notifications for all new transactions as they are accepted into the mempool. |    new_transaction  request_status |
|  4   | stop_notify_new_transactions | Stop sending either a tx accepted  notification when a new transaction is accepted into the mempool. |                             None                             |

##### Request json format

example: {"topic": "notify_raw_blocks"}



|  #   | field |              description              |
| :--: | :---: | :-----------------------------------: |
|  1   | Topic | Subscribe to the Topic of the message |

##### Respone json format

example:{"notification_type":"raw_blocks_connected","data":"0301ba9505dceb606ad2cffe7379d172ff8fc6c485d25ab9e051b18320deda2675cf6dfc61cc81bddf0540f13bffcbc3c3b76c1a34aba7c95b787c20a2b47366767ed1ce6ac9b5e9db8922c9c377e5192668bc0a367e4a4764f11e7c725ecced1d7b6a492974fab1b6d5bc90968080ea04879c9b83808080801e0107010001010802060038343636360001013effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8099c4d5990101160014d945f4ae7db4cd897f8d77d6d783b84a9d8ed9cc00"}



|  #   |      field       |                         description                          |
| :--: | :--------------: | :----------------------------------------------------------: |
|  1   | NotificationType |             The message type of the notification             |
|  2   |       Data       |                         data content                         |
|  3   |   ErrorDetail    | Error message when returning notification type is request_status |



##### Respone Error json format

example:  {"notification_type":"request_status","data":null,"error_detail":"Websocket Internal error: There is not this topic: notify_raw_blocks1"}



#### Websocket Client Example for go

```
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"

	ws "github.com/eiyaro/net/websocket"
	"github.com/eiyaro/protocol/bc/types"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var addr = flag.String("addr", "localhost:9888", "http service address")

func main() {

	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/websocket-subscribe"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer c.Close()

	// Subscribe to the raw blocks
	req := ws.WSRequest{
		Topic: "notify_raw_blocks",
	}
	msg, _ := json.Marshal(req)
	c.WriteMessage(websocket.TextMessage, msg)

	// Subscribe to the new transactions
	req = ws.WSRequest{
		Topic: "notify_new_transactions",
	}
	msg, _ = json.Marshal(req)
	c.WriteMessage(websocket.TextMessage, msg)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		var rep ws.WSResponse
		err = json.Unmarshal(msg, &rep)
		if err != nil {
			log.Printf("Unmarshal error:", err)
		}

		fmt.Println(rep.NotificationType)
		switch rep.NotificationType {
		case "raw_blocks_connected":
			data := fmt.Sprint(rep.Data)
			block := &types.Block{}
			err = block.UnmarshalText([]byte(data))
			if err != nil {
				fmt.Println(err)
			}
			log.Println(block)
		case "raw_blocks_disconnected":
			data := fmt.Sprint(rep.Data)
			block := &types.Block{}
			err = block.UnmarshalText([]byte(data))
			if err != nil {
				fmt.Println(err)
			}
			log.Println(block)
		case "new_transaction":
			data := fmt.Sprint(rep.Data)
			tx := &types.Tx{}
			err = tx.UnmarshalText([]byte(data))
			if err != nil {
				log.Println(err)
			}
			log.Println(tx)
		case "request_status":
			if rep.ErrorDetail != "" {
				log.Println(rep.ErrorDetail)
			}
		}

	}

}

```

