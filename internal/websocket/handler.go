package websocket

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/esgi-challenge/backend/config"
	"github.com/esgi-challenge/backend/internal/chat"
	"github.com/esgi-challenge/backend/internal/models"
	"github.com/esgi-challenge/backend/pkg/logger"
	"github.com/esgi-challenge/backend/pkg/request"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketHandler struct {
	Cfg         *config.Config
	ChatUseCase chat.UseCase
	Logger      logger.Logger
	clients     map[int]map[*websocket.Conn]bool
	mu          sync.Mutex
}

func NewWebSocketHandler(cfg *config.Config, chatUseCase chat.UseCase, logger logger.Logger) *WebSocketHandler {
	return &WebSocketHandler{
		Cfg:         cfg,
		ChatUseCase: chatUseCase,
		Logger:      logger,
		clients:     make(map[int]map[*websocket.Conn]bool),
	}
}

func (h *WebSocketHandler) ChatHandler(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		h.Logger.Errorf("Failed to set websocket upgrade: %+v", err)
		return
	}
	defer conn.Close()

	channelID, err := strconv.Atoi(ctx.Param("channelId"))
	if err != nil {
		h.Logger.Errorf("Invalid channel ID: %+v", err)
		return
	}

	h.mu.Lock()
	// Ensure the main clients map is initialized
	if h.clients == nil {
		h.clients = make(map[int]map[*websocket.Conn]bool)
		h.Logger.Info("Initialized the clients map")
	}

	// Register the client
	if _, ok := h.clients[channelID]; !ok {
		h.clients[channelID] = make(map[*websocket.Conn]bool)
		h.Logger.Infof("Initialized the clients map for channelID %d", channelID)
	}
	h.clients[channelID][conn] = true
	h.mu.Unlock()

	h.Logger.Infof("Client connected to channelID %d", channelID)

	// Handle incoming messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			h.Logger.Errorf("Error reading message: %+v", err)
			break
		}

		var msgData map[string]interface{}
		if err := json.Unmarshal(message, &msgData); err != nil {
			h.Logger.Errorf("Error unmarshalling message: %+v", err)
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "Invalid message format"))
			continue
		}

		token, ok := msgData["jwt"].(string)
		if !ok {
			h.Logger.Errorf("JWT token not found in message")
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "JWT token required"))
			continue
		}

		user, err := request.ValidateRoleWithoutHeader(h.Cfg.JwtSecret, token, models.STUDENT)
		if user == nil || err != nil {
			h.Logger.Errorf("Unauthorized user: %+v", err)
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "Unauthorized"))
			continue
		}

		channel, err := h.ChatUseCase.GetById(uint(channelID))
		if err != nil || (channel.FirstUserId != user.ID && channel.SecondUserId != user.ID) {
			h.Logger.Errorf("User not allowed in this channel: %+v", err)
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "Forbidden"))
			continue
		}

		content, ok := msgData["content"].(string)
		if !ok {
			h.Logger.Errorf("Message content not found in message")
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInvalidFramePayloadData, "Message content required"))
			continue
		}

		msg := &models.Message{
			Content:   content,
			ChannelId: uint(channelID),
			SenderId:  user.ID,
		}

		msg, err = h.ChatUseCase.SaveMessage(msg)
		if err != nil {
			h.Logger.Errorf("Failed to save message: %+v", err)
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "Internal server error"))
			continue
		}

		h.Logger.Infof("Message sent: %+v", msg)

		response := map[string]interface{}{
			"content":   msg.Content,
			"channelId": msg.ChannelId,
			"senderId":  msg.SenderId,
			"createdAt": msg.CreatedAt,
		}
		responseMsg, err := json.Marshal(response)
		if err != nil {
			h.Logger.Errorf("Failed to marshal response message: %+v", err)
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseInternalServerErr, "Internal server error"))
			continue
		}

		h.broadcastMessage(channelID, responseMsg)
	}

	// Unregister the client
	h.mu.Lock()
	delete(h.clients[channelID], conn)
	// Remove the channel if no clients are left
	if len(h.clients[channelID]) == 0 {
		delete(h.clients, channelID)
		h.Logger.Infof("Removed the clients map for channelID %d", channelID)
	}
	h.mu.Unlock()
	h.Logger.Infof("Client disconnected from channelID %d", channelID)
}

func (h *WebSocketHandler) broadcastMessage(channelID int, message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for client := range h.clients[channelID] {
		if err := client.WriteMessage(websocket.TextMessage, message); err != nil {
			h.Logger.Errorf("Error writing message: %+v", err)
			client.Close()
			delete(h.clients[channelID], client)
		}
	}
}
