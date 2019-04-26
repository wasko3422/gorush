package gorush

import (
	"github.com/wasko3422/gorush/config"
	"github.com/wasko3422/gorush/storage"

	"github.com/appleboy/go-fcm"
	"github.com/sideshow/apns2"
	"github.com/sirupsen/logrus"
)

var (
	// PushConf is gorush config
	PushConf config.ConfYaml
	// QueueNotification is chan type
	QueueNotification chan PushNotification
	// ApnsClient is apns client
	ApnsClient *apns2.Client
	// VOIPClient is Voice Over IP client
	VOIPClient *apns2.Client
	// SafariClient is safari client
	SafariClient *apns2.Client
	// FCMClient is apns client 321
	FCMClient *fcm.Client
	// LogAccess is log server request log
	LogAccess *logrus.Logger
	// LogError is log server error log
	LogError *logrus.Logger
	// StatStorage implements the storage interface
	StatStorage storage.Storage
)
