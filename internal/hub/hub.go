package hub

import "github.com/vbrenister/go-htmx-chat/internal/client"

type Hub struct {
	clients map[*client.Client]bool
}
