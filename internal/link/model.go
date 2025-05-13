package link

import (
	"gorm.io/gorm"
	"math/rand"
	"url/internal/stat"
)

type Link struct {
	gorm.Model
	Url   string      `json:"url"`
	Hash  string      `json:"hash" gorm:"uniqueIndex"`
	Stats []stat.Stat `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	link.GenerateHash()
	return link

}

func (link *Link) GenerateHash() {
	link.Hash = randSeq(10)
}

var letters = []rune("abcdefghijklmnjprstuvwxyzABCDEFGHIJKLMNOPRSTUWVXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
