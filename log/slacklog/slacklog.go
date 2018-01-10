package slacklog

import (
	slackbot "c3m/common/chatbot/slack"
	"c3m/log"
	"fmt"
	"github.com/nlopes/slack"
)

type Logger struct {
	Token   string
	Channel string
}

var (
	Default = Logger{"xoxb-16475074549-7b5JA4PongYN5vzEBYL3c21B", "#jobfeedslog"}
)

func NewLogger(token, channel string) Logger {
	return Logger{token, channel}
}

func (this Logger) Debug(message string) {
	log.Debug(message)
	this.logPrint("TRACE", message)
}
func (this Logger) Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
	this.logPrintf("TRACE", format, args...)
}
func (this Logger) Info(message string) {
	log.Info(message)
	this.logPrint("INFO", message)
}
func (this Logger) Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
	this.logPrintf("INFO", format, args...)
}
func (this Logger) Warn(message string) {
	log.Warn(message)
	this.logPrint("WARN", message)
}
func (this Logger) Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
	this.logPrintf("WARN", format, args...)
}
func (this Logger) Error(message string) {
	log.Error(message)
	this.logPrint("ERROR", message)
}
func (this Logger) Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
	this.logPrintf("ERROR", format, args...)
}

func (this Logger) Attachment(attachment slack.Attachment) {
	go slackbot.SendAttachment(this.Token, this.Channel, attachment)
}

func (this Logger) logPrint(severity, message string) {
	slackMessage := slack.Attachment{}
	slackMessage.Fallback = message
	slackMessage.Title = fmt.Sprintf("%s %s", getEmoji(severity), severity)
	slackMessage.Text = message

	this.Attachment(slackMessage)
}
func (this Logger) logPrintf(severity, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	this.logPrint(severity, message)
}

func getEmoji(severity string) string {
	var emoji string
	switch severity {
	case "WARN":
		emoji = ":heavy_exclamation_mark:"
	case "ERROR":
		emoji = ":bangbang:"
	default:
		emoji = ":information_source:"
	}

	return emoji
}
