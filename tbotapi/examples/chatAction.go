// Copyright 2015-2016 mrd0ll4r and contributors. All rights reserved.
// Use of this source code is governed by the MIT license, which can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/mrd0ll4r/tbotapi"
	"github.com/mrd0ll4r/tbotapi/examples/boilerplate"
	"log"
	"strings"
)

func main() {
	apiToken := "323437881:AAEHEPqSX-VKIKSLzs6wC_MiiDw4k8GktQk"

	updateFunc := func(update tbotapi.Update, api *tbotapi.TelegramBotAPI) {
		switch update.Type() {
		case tbotapi.MessageUpdate:
			msg := update.Message
			typ := msg.Type()
			var Text = ""
			Text = strings.TrimRight(update.Message.Text," ")
			if typ != tbotapi.TextMessage {
				// Ignore non-text messages for now.
				fmt.Println("Ignoring non-text message")
				return
			}

			// Note: Bots cannot receive from channels, at least no text messages. So we don't have to distinguish anything here.

			// Display the incoming message.
			// msg.Chat implements fmt.Stringer, so it'll display nicely.
			// We know it's a text message, so we can safely use the Message.Text pointer.


			if Text == "/голос" {
				err := api.NewOutgoingChatAction(tbotapi.NewRecipientFromChat(msg.Chat), tbotapi.ChatActionUploadPhoto).Send()
				if err != nil {
					fmt.Printf("Error sending: %s\n", err)
					return
				}
				outMsg, err := api.NewOutgoingMessage(tbotapi.NewRecipientFromChat(msg.Chat), "Ну, допустим, гав..").Send()
				fmt.Printf("->%d, To:\t%s, Text: %s\n", outMsg.Message.ID, outMsg.Message.Chat, *outMsg.Message.Text)
			}

		case tbotapi.InlineQueryUpdate:
			fmt.Println("Ignoring received inline query: ", update.InlineQuery.Query)
		case tbotapi.ChosenInlineResultUpdate:
			fmt.Println("Ignoring chosen inline query result (ID): ", update.ChosenInlineResult.ID)

		default:
			fmt.Println("Ignoring unknown Update type.")
		}
	}

	// Run the bot, this will block.
	boilerplate.RunBot(apiToken, updateFunc, "ChatAction", "Demonstrates chat actions")
}
