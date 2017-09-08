package main

import (
	// Includes padrões pro Bot
	"fmt"
	"os"
	"log"
	"regexp"
	"bufio"
	"strings"
	// Bot package
	"./ircBot"
	// Tools package
	"./tools"
	// Includes dos comandos criados
	"time"
)

func main(){

	// "Try catch" do Go
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	// Cria o bot
	bot := ircBot.NewBot(
		"irc.chat.twitch.tv", // Host (default)
		"6667", // Porta (default)
		"nick-here", // Nick do Bot/ContaNormal
		"channel-here", // Canal para se auto conectar
		"oauth:token-here", // Token do Bot/ContaNormal
	)

	// Conecta o Bot
	conn, _ := bot.Connect()
	defer conn.Close()

	// Cria o buffer
	bot.Maker()

	// Inicia a função "PingLoop()" em uma nova thread
	go bot.PingLoop()

	// Pede ao servidor para enviar informação completa de cada mensagem
	bot.Send("CAP REQ :twitch.tv/tags")
	bot.Send("CAP REQ :twitch.tv/membership")

	// Expressão Regular default para tratamento da resposta do servidor
	// reg, _ := regexp.Compile(":([a-z0-9_]*)!([a-z0-9_]*)@([a-z0-9.-]*) ([A-Z]*) #([a-z0-9_]*) :(![a-z0-9_]*)?(.*)")

	// 
	inf, _ := regexp.Compile(`([az\-\w]+)=([a-zA-Z0-9_\:\,\-\#\/]*|\s*)`)
	reg, _ := regexp.Compile(`.*? :([a-zA-Z0-9_\-]+)?@?.*[twitch.tv]\s([a-zA-Z]+)?\s?#([a-zA-Z0-9_\-]+)?\s:(![a-zA-Z0-9_\-]+)?\s?(.*)?`)
	for {
		// Recebe os dados do servidor
		line, err := bot.Buffer.ReadLine()
		if err != nil {
			log.Fatal("Unable to receive data from IRC server ", err)
			os.Exit(1)
		}

		// Função para segunda checagem do "Eu to vivo"
		isPing, _ := regexp.MatchString("PING", line)
		if isPing == true {
			data := strings.Split(line, "PING ")
			bot.Send("PONG "+data[1]+"")
			time.Sleep(50 * time.Millisecond)
			continue
		}

		// Passa para o próximo loop evitando comparação de mensagens do sistema
		// INICIO
		if strings.Contains(line, ":tmi.twitch.tv PONG ") == true {
			continue
		}
		if strings.Contains(line, "JOIN #") == true {
			// user joined
			continue
		}
		if strings.Contains(line, ":jtv MODE") == true {
			// privilegios
			continue
		}
		// FIM

		// Aplica a Expressão Regular
		t := reg.FindStringSubmatch(line)
		// Se a Expressão Regular retornar 8 argumentos
		if len(t) == 1 {
			username := t[0][1]
			tipo := t[0][2]
			channel := t[0][3]
			command := t[0][4]
			text := t[0][5]
			var info = make(map[string]string)
			for _, v := range i {
				info[v[1]] = v[2]
			}

			// Exemplo de receive de bits
			if info["bits"] != "" {
				bot.Msg(channel, "Usuário "+username+" doou "+info["bits"]+" bits.")
				time.Sleep(1000 * time.Millisecond)
				continue
			}

			// Comandos
			if command == "!gold" {
				request, _ := http.Get("https://wowtoken.info/snapshot.json")
				r_body, _ := ioutil.ReadAll(request.Body)
				request.Body.Close()
				body := string(r_body[:])
				reg2, _ := regexp.Compile("{\"NA\":{\"timestamp\":([0-9]*),\"raw\":{\"buy\":([0-9]*),\"")
				json := reg2.FindStringSubmatch(body)
				gold_ := json[2]
				gold := gold_[:len(gold_)-3] + "," + gold_[len(gold_)-3:]
				// fmt.Printf("%s\n", json[2])
				bot.Msg(channel, "[Bolsa de Azeroth] informa:")
				time.Sleep(1500 * time.Millisecond)
				bot.Msg(channel, "[NA] Cotação do OURO: " + gold + "g")
				continue
			}

			// Comandos
			if command == "!dolar" {
				price := tools.Get("http://api.dolarhoje.com")
				time.Sleep(1000 * time.Millisecond)
				bot.Msg(channel, "[DOLAR hoje] R$ " + price)
				continue
			}

			// Imprime a mensagem formatada recebida de um canal
			// 	"[#<canal>] <username>:<comando> <mensagem>"
			fmt.Printf("[#%s] %s: %s %s\n", channel, username, command, text)
		} else {
			// Impreme outras mensagens recebidas do servidor que
			//	não sejam referentes à mensagens de usuários
			fmt.Printf("%s\n", line)
		}

	}

	// Bye bye
	os.Exit(0)
}
	reg, _ := regexp.Compile(":([a-z0-9_]*)!([a-z0-9_]*)@([a-z0-9.-]*) ([A-Z]*) #([a-z0-9_]*) :(![a-z0-9_]*)?(.*)")
		if len(t) >= 8 {
			username := t[1]
			channel := t[5]
			command := t[6]
			text := t[7]
			// Comandos
			if command == "!gold" {
				request, _ := http.Get("https://wowtoken.info/snapshot.json")
				r_body, _ := ioutil.ReadAll(request.Body)
				request.Body.Close()
				body := string(r_body[:])
				reg2, _ := regexp.Compile("{\"NA\":{\"timestamp\":([0-9]*),\"raw\":{\"buy\":([0-9]*),\"")
				json := reg2.FindStringSubmatch(body)
				gold_ := json[2]
				gold := gold_[:len(gold_)-3] + "," + gold_[len(gold_)-3:]
				// fmt.Printf("%s\n", json[2])
				bot.Msg(channel, "[Bolsa de Azeroth] informa:")
				time.Sleep(1500 * time.Millisecond)
				bot.Msg(channel, "[NA] Cotação do OURO: " + gold + "g")