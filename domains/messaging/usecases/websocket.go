package usecases

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/entities"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/repositories"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	Clients map[uint]*entities.Client = make(map[uint]*entities.Client)
	Mutex   sync.Mutex                = sync.Mutex{}
)

type IWebsocketUseCase interface {
	HandleNewWebsocketClient(ctx context.Context, conn *websocket.Conn)
}

type WebsocketUseCase struct {
	userRepository  repositories.IUserRepository
	groupRepository repositories.IGroupRepository
}

func NewWebsocketUseCase() IWebsocketUseCase {
	return &WebsocketUseCase{
		userRepository:  repositories.NewUserRepository(),
		groupRepository: repositories.NewGroupRepository(),
	}
}

func addNewClient(client *entities.Client) {
	Mutex.Lock()
	defer Mutex.Unlock()

	Clients[client.ID] = client
}

func (w *WebsocketUseCase) HandleNewWebsocketClient(ctx context.Context, conn *websocket.Conn) {
	userID, ok := util.GetUserIDFromContext(ctx)
	if !ok {
		return
	}

	logger := util.GetLoggerFromContext(ctx)

	user, err := w.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		logger.Error("get user by id error", zap.Error(err))
		conn.Close()
		return
	}

	client := entities.NewClient(userID, conn)
	client.Start(ctx)

	addNewClient(client)

	for {
		event := <-client.ReadChannel
		switch event.EventType {
		case entities.EventType_SendMessage:
			var msg entities.EventSendMessage
			err := json.Unmarshal(event.EventData, &msg)
			if err != nil {
				logger.Error("unmarshal event send message error", zap.Error(err))
				continue
			}

			if msg.GroupID != nil {
				group, err := w.groupRepository.GetGroupByID(ctx, *msg.GroupID)
				if err != nil {
					logger.Error("get group by id error", zap.Error(err))
					continue
				}

				ok, err := w.groupRepository.IsUserInGroup(ctx, userID, group.ID)
				if err != nil {
					logger.Error("check user in group error", zap.Error(err))
					continue
				}

				if !ok {
					logger.Error("user not in group", zap.Uint("user_id", userID), zap.Uint("group_id", *msg.GroupID))
					continue
				}

				groupMemberIDs, err := w.groupRepository.GetGroupMemberIDs(ctx, *msg.GroupID)
				if err != nil {
					logger.Error("get group member ids error", zap.Error(err))
					continue
				}

				for _, memberID := range groupMemberIDs {
					client, ok := Clients[memberID]
					if ok {
						client.SendEvent(entities.NewEventReceiveMessageUser(msg.Content, user))
					}
				}
			} else if msg.TargetUserID != nil {
				targetClient, ok := Clients[*msg.TargetUserID]
				if !ok {
					logger.Error("target client not found", zap.Uint("target_id", *msg.TargetUserID))
					continue
				}

				targetClient.SendEvent(entities.NewEventReceiveMessageUser(msg.Content, user))
			}
		}
	}
}
