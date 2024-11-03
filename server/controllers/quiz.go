package controllers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/real-time-vocab-quiz/server/messages"
	"github.com/real-time-vocab-quiz/server/services"
)

type QuizController struct {
	QuizService *services.QuizService
	mu          sync.Mutex
}

var quizController *QuizController

var clients map[string]map[*websocket.Conn]string

func newQuizController(quizService *services.QuizService) *QuizController {
	if quizController == nil {
		quizController = &QuizController{
			QuizService: quizService,
		}
		clients = make(map[string]map[*websocket.Conn]string)
	}
	return quizController
}

func (qc *QuizController) JoinQuizByCode(c *gin.Context) {
	// Upgrade HTTP request to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.Error(c.Writer, "Failed to upgrade to websocket", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	// TODO: hardcode here first, check quiz session and userId and quiz Id later
	code := c.Param("code")
	_, err = qc.QuizService.GetQuizByCode(code)
	if err != nil {
		conn.WriteJSON(gin.H{"error": "Quiz not found"})
		return
	}

	// TODO: create new session in DB

	var joinMessage messages.JoinMessage
	if err := conn.ReadJSON(&joinMessage); err != nil {
		conn.WriteJSON(gin.H{"type": "error", "message": "Invalid join parameters"})
		return
	}

	// Add client to the quiz's client list
	qc.mu.Lock()
	if clients[code] == nil {
		clients[code] = make(map[*websocket.Conn]string)
	}
	clients[code][conn] = joinMessage.Data.UserId
	participantsMessage := qc.getParticipantsMessage(code)
	qc.mu.Unlock()

	// Broadcast updated participants list
	qc.broadcastParticipants(code, participantsMessage)

	for {
		var message messages.Base[interface{}]
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Println("Failed to read message", err)
			break
		}

		// TODO: handle scoring and updating leaderboard if any quiz update

	}

	delete(clients[code], conn)
	participantsMessage = qc.getParticipantsMessage(code)

	qc.broadcastParticipants(code, participantsMessage)
}

func (qc *QuizController) getParticipantsMessage(code string) messages.ParticipantsMessage {
	participants := []messages.Participant{}
	for _, userId := range clients[code] {
		participants = append(participants, messages.Participant{UserId: userId})
	}
	return messages.ParticipantsMessage{
		Base: messages.Base[messages.ParticipantList]{
			EventType: "participants",
			Data: messages.ParticipantList{
				Participants: participants,
			},
		},
	}
}

func (qc *QuizController) broadcastParticipants(code string, participantsMessage messages.ParticipantsMessage) {
	for conn := range clients[code] {
		conn.WriteJSON(participantsMessage)
	}
}
