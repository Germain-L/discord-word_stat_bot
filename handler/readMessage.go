package handler

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h handler) ReadMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		log.Println("Ignoring message from self")
		return
	}

	userId := m.Author.ID
	userName := m.Author.Username
	message := m.Content

	// keep only words from message, remove all non-alphanumeric characters
	words := strings.FieldsFunc(message, func(r rune) bool {
		return !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r >= 'Ç' && r <= 'Ñ')
	})

	// convert to lowercase
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	woordCount := make(map[string]int)

	// count the times each word appears in the message
	for _, word := range words {
		woordCount[word]++
	}

	queryUserId := 0
	// create user in database if not exists
	h.Db.QueryRow("SELECT id FROM users WHERE id = $1", userId).Scan(&queryUserId)

	if queryUserId == 0 {
		_, err := h.Db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", userId, userName)
		if err != nil {
			log.Println(err)
		}
	}

	queryWordId := 0
	saysExists := 0

	// insert word count into database
	for word, count := range woordCount {
		// insert word in database if not exists and get id
		queryWordId = 0
		h.Db.QueryRow("SELECT id FROM word WHERE word_value = $1", word).Scan(&queryWordId)
		if queryWordId == 0 {
			_, err := h.Db.Exec("INSERT INTO word (word_value) VALUES ($1)", word)

			if err != nil {
				log.Println(err)
			}

			h.Db.QueryRow("SELECT id FROM word WHERE word_value = $1", word).Scan(&queryWordId)
		}

		// check if says exists
		h.Db.QueryRow("SELECT SELECT COUNT(DISTINCT word_id) FROM says WHERE user_id = $1 AND word_id = $2", userId, queryWordId).Scan(&saysExists)

		// if says id is 0, insert new row
		if saysExists == 0 {
			_, err := h.Db.Exec("INSERT INTO says (user_id, word_id, word_count) VALUES ($1, $2, $3)", userId, queryWordId, count)
			if err != nil {
				log.Println(err)
			}
		} else {
			_, err := h.Db.Exec("UPDATE says SET word_count = word_count + $1 WHERE id = $2", count, saysExists)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
