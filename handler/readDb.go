package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (h handler) ReadDb(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		log.Println("Ignoring message from self")
		return
	}

	if m.Content != "!stats" {
		return
	}

	rows, err := h.Db.Query("SELECT word_value, says.word_count, users.name FROM word INNER JOIN says ON says.word_id=word.id INNER JOIN users ON users.id=$1 AND users.id=says.user_id", m.Author.ID)
	if err != nil {
		log.Panicln(err)
	}

	for rows.Next() {
		var wordValue string
		var wordCount int
		var userName string
		err = rows.Scan(&wordValue, &wordCount, &userName)
		if err != nil {
			log.Panicln(err)
		}
	}

	// TODO
	rows.Close()
}
