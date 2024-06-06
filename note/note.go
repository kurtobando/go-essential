package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note !make sure to capitalize the keys, for them to make it public
type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

func New(title string, content string) *Note {
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().Format(time.DateTime),
	}
}

func (n *Note) PrintNote() {
	fmt.Printf("[TITLE] %s[CONTENT] %s[CREATED_AT]%s\n", n.Title, n.Content, n.CreatedAt)
}

func (n *Note) SaveAsJson() {
	n.Title = strings.TrimSpace(n.Title)
	n.Content = strings.TrimSpace(n.Content)

	b, err := json.Marshal(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Saving %s \n", string(b))

	err = os.WriteFile("tmp/note.json", b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
