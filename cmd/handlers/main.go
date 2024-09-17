package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/bsdlp/discord-interactions-go/interactions"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// TODO: BEFORE COMMIT: HIDE THIS KEY
		hexEncodedDiscordPubkey := "{INSERT_PUBLIC_KEY_HERE}"
		discordPubkey, err := hex.DecodeString(hexEncodedDiscordPubkey)
		if err != nil {
			fmt.Printf("error decoding hex encoded discord pubkey: %v\n", err)
		}
		verified := interactions.Verify(r, ed25519.PublicKey(discordPubkey))
		if !verified {
			http.Error(w, "signature mismatch", http.StatusUnauthorized)
			return
		}

		defer r.Body.Close()
		var data interactions.Data
		jsonErr := json.NewDecoder(r.Body).Decode(&data)
		if jsonErr != nil {
			fmt.Printf("error decoding request body: %v\n", jsonErr)
		}

		fmt.Println(data.Type)

		// respond to ping
		if data.Type == interactions.Ping {
			_, err := w.Write([]byte(`{"type":1}`))
			if err != nil {
				fmt.Printf("error writing response: %v\n", err)
			}
			return
		}

		// handle command
		response := &interactions.InteractionResponse{
			Type: interactions.ChannelMessageWithSource,
			Data: &interactions.InteractionApplicationCommandCallbackData{
				Content: "he's ghet!",
			},
		}

		var responsePayload bytes.Buffer
		err = json.NewEncoder(&responsePayload).Encode(response)
		if err != nil {
			fmt.Printf("error encoding response: %v\n", err)
		}

		fmt.Println(data.ResponseURL())
		fmt.Println()

		fmt.Println(responsePayload.String())

		resp, err := http.Post(data.ResponseURL(), "application/json", &responsePayload)
		if err != nil {
			fmt.Printf("error sending response: %v\n", err)
		} else {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("error reading response body: %v\n", err)
			} else {
				fmt.Println(string(body))
			}
		}
		io.WriteString(w, "success")
	})

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
