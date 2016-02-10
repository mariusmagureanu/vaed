package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"github.com/mariusmagureanu/vaed/plugins"
	"github.com/mariusmagureanu/vaed/response"
)

func socketHandler(ws *websocket.Conn) {
	for {
		receivedStream := make([]byte, 256)
		byteCount, err := ws.Read(receivedStream)

		if err != nil {
			return
		}

		if byteCount > 0 {
			inputHeaders := make(map[string]string)

			decode_err := json.Unmarshal(receivedStream[:byteCount], &inputHeaders)

			if decode_err != nil {
				return
			}
			out := runPlugins("secred_key", "endpoint_XXX", "consumer_XXX")

			jsonOut, encode_err := json.Marshal(out.Headers)
			if encode_err != nil {
				return
			}

			ws.Write(jsonOut)
		}
	}
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	out := runPlugins("secred_key", "endpoint_XXX", "consumer_XXX")

	for k, v := range out.Headers {
		w.Header().Add(k, v)
	}
	w.WriteHeader(http.StatusOK)
}

//runPLugins runs plugins
func runPlugins(apiKey string, endpointId string, consumerId string) response.Out {
	pluginNames := []string{"basic-auth", "metrics", "rate-limit"}

	out := response.Out{}
	out.Headers = make(map[string]string)

	for _, pluginName := range pluginNames {
		plugin, err := plugins.Get(pluginName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		plugin.Run(&out)
	}

	return out
}

func main() {
	http.Handle("/s", websocket.Handler(socketHandler))
	http.HandleFunc("/h", httpHandler)

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic("Error: " + err.Error())
	}
}
