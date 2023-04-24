// utils/email.go
package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"mime/multipart"
	"net/smtp"
	"net/textproto"

	"path/filepath"
)

const (
	smtpServer = "smtp.gmail.com"
	smtpPort   = "587"
	emailUser  = "bahatoho.1709@gmail.com"
	emailPass  = "wydkweqvesviquzo"
)

func SendEmail(to, subject, htmlBody string) error {
	auth := smtp.PlainAuth("", emailUser, emailPass, smtpServer)

	buffer := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buffer)

	header := make(textproto.MIMEHeader)
	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("Content-Disposition", "inline")
	htmlWriter, _ := writer.CreatePart(header)
	_, _ = htmlWriter.Write([]byte(htmlBody))

	// Attach logo header image
	logoPath := "utils/templates/images/send_email1.png"
	logoFilename := filepath.Base(logoPath)
	logoFile, err := os.ReadFile(logoPath)
	if err != nil {
		return err
	}
	logoBase64 := base64.StdEncoding.EncodeToString(logoFile)

	header = make(textproto.MIMEHeader)
	header.Set("Content-Type", "image/png")
	header.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", logoFilename))
	header.Set("Content-Transfer-Encoding", "base64")
	header.Set("Content-ID", "<logo>")
	imageWriter, _ := writer.CreatePart(header)
	_, _ = imageWriter.Write([]byte(logoBase64))

	// Attach logo footer image
	logo2Path := "utils/templates/images/send_email2.png"
	logoFilename2 := filepath.Base(logo2Path)
	logoFile2, err := os.ReadFile(logo2Path)
	if err != nil {
		return err
	}
	logoBase642 := base64.StdEncoding.EncodeToString(logoFile2)

	header = make(textproto.MIMEHeader)
	header.Set("Content-Type", "image/png")
	header.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", logoFilename2))
	header.Set("Content-Transfer-Encoding", "base64")
	header.Set("Content-ID", "<logo2>")
	imageWriter2, _ := writer.CreatePart(header)
	_, _ = imageWriter2.Write([]byte(logoBase642))

	// Attach name icon
	nameIconPath := "utils/templates/icon/name-icon.png"
	nameIconFilename := filepath.Base(nameIconPath)
	nameIconFile, err := os.ReadFile(nameIconPath)
	if err != nil {
		return err
	}
	nameIconBase64 := base64.StdEncoding.EncodeToString(nameIconFile)

	header = make(textproto.MIMEHeader)
	header.Set("Content-Type", "image/png")
	header.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", nameIconFilename))
	header.Set("Content-Transfer-Encoding", "base64")
	header.Set("Content-ID", "<nameIcon>")
	imageWriter, _ = writer.CreatePart(header)
	_, _ = imageWriter.Write([]byte(nameIconBase64))

	// Attach phone icon
	phoneIconPath := "utils/templates/icon/phone-icon.png"
	phoneIconFilename := filepath.Base(phoneIconPath)
	phoneIconFile, err := os.ReadFile(phoneIconPath)
	if err != nil {
		return err
	}
	phoneIconBase64 := base64.StdEncoding.EncodeToString(phoneIconFile)

	header = make(textproto.MIMEHeader)
	header.Set("Content-Type", "image/png")
	header.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", phoneIconFilename))
	header.Set("Content-Transfer-Encoding", "base64")
	header.Set("Content-ID", "<phoneIcon>")
	imageWriter, _ = writer.CreatePart(header)
	_, _ = imageWriter.Write([]byte(phoneIconBase64))

	// Attach email icon
	emailIconPath := "utils/templates/icon/email-icon.png"
	emailIconFilename := filepath.Base(emailIconPath)
	emailIconFile, err := os.ReadFile(emailIconPath)
	if err != nil {
		return err
	}
	emailIconBase64 := base64.StdEncoding.EncodeToString(emailIconFile)

	header = make(textproto.MIMEHeader)
	header.Set("Content-Type", "image/png")
	header.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", emailIconFilename))
	header.Set("Content-Transfer-Encoding", "base64")
	header.Set("Content-ID", "<emailIcon>")
	imageWriter, _ = writer.CreatePart(header)
	_, _ = imageWriter.Write([]byte(emailIconBase64))

	// Attach address icon
	addressIconPath := "utils/templates/icon/address-icon.png"
	addressIconFilename := filepath.Base(addressIconPath)
	addressIconFile, err := os.ReadFile(addressIconPath)
	if err != nil {
		return err
	}
	addressIconBase64 := base64.StdEncoding.EncodeToString(addressIconFile)

	header = make(textproto.MIMEHeader)
	header.Set("Content-Type", "image/jpeg")
	header.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", addressIconFilename))
	header.Set("Content-Transfer-Encoding", "base64")
	header.Set("Content-ID", "<addressIcon>")
	imageWriter, _ = writer.CreatePart(header)
	_, _ = imageWriter.Write([]byte(addressIconBase64))
	_ = writer.Close()

	msg := "From: " + emailUser + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: multipart/mixed; boundary=" + writer.Boundary() + "\n\n" +
		buffer.String()

	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, emailUser, []string{to}, []byte(msg))
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return err
	}

	return nil
}

func GetEmailTemplate() (string, error) {
	file, err := os.Open("utils/templates/email_template.html")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content bytes.Buffer
	_, err = io.Copy(&content, file)
	if err != nil {
		return "", err
	}
	return content.String(), nil
}
