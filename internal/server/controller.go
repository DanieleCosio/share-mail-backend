package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sharemail/internal/config"
	"sharemail/internal/db"
	"sharemail/internal/logs"
	"sharemail/internal/orm"
)

func getEmailLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST.String() {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var messageData GetEmailLinkBody
	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logs.Logger().Error().Err(err).Msg("Failed to decode request body")
		return
	}

	if messageData.RequestAccountOwner == "" || messageData.MessageHtml == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("RequestAccountOwner: %s\n", messageData.RequestAccountOwner)
	log.Printf("MessageHtml: %s\n", messageData.MessageHtml)

	redisClient, err := db.GetRedisConnection()
	if err != nil {
		logs.Logger().Error().Err(err).Msg("Failed to get Redis connection")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	urlHash, err := redisClient.RPop(ctx, config.AppConfig["FREE_REDIS_KEY"]).Result()
	if err != nil {
		logs.Logger().Error().Err(err).Msg("Failed to get URL hash from Redis")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	sqlOrm, err := db.GetOrmConnection()
	if err != nil {
		logs.Logger().Error().Err(err).Msg("Failed to get ORM connection")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	email, err := sqlOrm.CreateEmail(ctx, orm.CreateEmailParams{
		OwnerEmail: messageData.RequestAccountOwner,
		EmailHtml:  messageData.MessageHtml,
		UrlHash:    urlHash,
	})

	for _, attachment := range messageData.Attachments {
		_, err = sqlOrm.CreateAttachment(ctx, orm.CreateAttachmentParams{
			EmailID:       email.ID,
			Name:          attachment.Name,
			MimeType:      attachment.MimeType,
			Size:          attachment.Size,
			AttachmentUrl: attachment.Url,
			PreviewUrl:    &attachment.PreviewUrl,
		})
	}
	sqlOrm = nil

	if err != nil {
		logs.Logger().Error().Err(err).Msg("Failed to create email in database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Email ID: %d\n", email.ID)

	redisClient.LPush(ctx, config.AppConfig["BUSY_REDIS_KEY"], urlHash)
	expirationDate := email.CreatedAt.Time.AddDate(0, 0, 2)

	jsonResponse := GetEmailLinkResponse{
		Url: fmt.Sprintf(
			"%s/%s/%s",
			config.AppConfig["BASE_URL"],
			config.AppConfig["EMAILS_LINKS_PREFIX"],
			urlHash,
		),
		ExpireAt: expirationDate.Format("2006-01-02 15:04:05"),
	}
	jsonResponse.Dispatch(&w, jsonResponse)
}

func ping(w http.ResponseWriter, r *http.Request) {
	jsonResponse := PingResponse{Success: true}
	jsonResponse.Dispatch(&w, jsonResponse)
}
