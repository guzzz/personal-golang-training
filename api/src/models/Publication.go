package models

import (
	"errors"
	"strings"
	"time"
)

// Publication representa uma publicação
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	Photo      string    `json:"photo,omitempty"`
	AuthorID   uint64    `json:"authorID,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

// Prepare vai chamar metodos de validar e formatar a publicação
func (publication *Publication) Prepare() error {
	if erro := publication.validate(); erro != nil {
		return erro
	}

	publication.format()
	return nil
}

func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("o título é obrigatório")
	}

	if publication.Content == "" {
		return errors.New("o conteudo é obrigatório")
	}

	return nil
}

func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
