package model

import (
	"auth/utils/constants"
	"auth/utils/helper"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

// all notifications messages must have following methods,
// which prepare message to be sent
type InterMessage interface {
	// this indicates the sender address
	GetFrom() string
	// indicates email addresses of receivers
	GetToList() []string
	// the subject of the email message
	GetSubject() string
	// body of the message in html format
	GetHtml() string
	// body of the message in plain text format
	GetPlainText() string
}

type InterDialer interface {
	DialAndSend(m... *gomail.Message) error
}

// errors
var errorMessageInvalidFrom = errors.New("invalid sender address")
var errorMessageInvalidToList = errors.New("invalid receiver address")
var errorMessageInvalidSubject = errors.New("invalid subject")
var errorMessageInvalidBody = errors.New("invalid body of a message")

// validate
func MessageOnlyValidateOne(message InterMessage) error {
	switch {
	case message.GetFrom() == "":
		return errorMessageInvalidFrom
	case len(message.GetToList()) == 0:
		return errorMessageInvalidToList
	case message.GetSubject() == "":
		return errorMessageInvalidSubject
	case len(message.GetHtml()) + len(message.GetPlainText()) == 0:
		return errorMessageInvalidBody
	}

	return nil
}

// prepare message to send
func MessageOnlyPrepareMail(message InterMessage) (*gomail.Message, error) {
	// validate message
	if err := MessageOnlyValidateOne(message); err != nil {
		return &gomail.Message{}, err
	}

	// prepare mail message to send
	msgToSend := gomail.NewMessage()

	// set headers here
	msgToSend.SetHeader("From", message.GetFrom())
	msgToSend.SetHeader("To", message.GetToList()...)
	msgToSend.SetHeader("Subject", message.GetSubject())

	// set body
	if len(message.GetHtml()) > 0 {
		msgToSend.SetBody("text/html", message.GetHtml())
	}

	if len(message.GetPlainText()) > 0 {
		msgToSend.SetBody("text/plain", message.GetPlainText())
	}

	return msgToSend, nil
}

// send message
func MessageOnlySend(dialer InterDialer, message *gomail.Message) error {
	startTime := helper.OnlyGetCurrentTime()
	err := dialer.DialAndSend(message)
	fmt.Println("Message took: ", helper.OnlyGetCurrentTime().Sub(startTime).Seconds(), " s")
	return err
}

// establish a connection & send a message
func MessageDialAndSend(message *gomail.Message) error {
	/*
		TODO: get smtp credential from db
	 */

	port, err := strconv.Atoi(os.Getenv(constants.CONST_SMTP_SERVER_PORT))
	if err != nil {
		return err
	}

	// create dialer (establish connection)
	// 'smtp-relay.sendinblue.com', 587, 'yerassyl.danay@nu.edu.kz', 'pkhRjzw93cBFI6NE'
	var d = &gomail.Dialer{
		Host:      os.Getenv(constants.CONST_SMTP_SERVER_HOST),
		Port:      port,
		Username:  os.Getenv(constants.CONST_SMTP_SERVER_USERNAME),
		Password:  os.Getenv(constants.CONST_SMTP_SERVER_PASSWORD),
	}

	err = MessageOnlySend(d, message)
	return err
}
