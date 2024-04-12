package tools

import (
	"config"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	mathRand "math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
	"utils"
)

func GenerateToken(id int, role int, duration int) (string, string) {
	access_token, _, _ := utils.CreateToken("access", config.ConfigList.Jwt.Secret, id, role, duration)
	refresh_token, _, _ := utils.CreateToken("refresh", config.ConfigList.Jwt.Secret, id, role, duration)
	// 保存token models.user_token

	return access_token, refresh_token
}

func GenerateAccessToken(id int, role int, duration int) string {
	access_token, _, _ := utils.CreateToken("access", config.ConfigList.Jwt.Secret, id, role, duration)
	return access_token
}

func GenerateKey() (string, error) {
	key := make([]byte, 16) // 16 bytes = 128 bits = 32 hex characters
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}

func GenerateRandomNumberString(length int) string {
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(900000) + 100000

	return fmt.Sprintf("%d", randomNumber)
}

func SendMail(to string, code *string, duration *int) error {
	// 设置认证信息。
	Email := config.GetConfigMap("email").(*config.Email)
	EmailReg := config.GetConfigMap("email_reg").(*config.EmailReg)
	emailUser := Email.EmailUser
	emailPass := Email.EmailPass
	emailTLS := Email.EmailTLS
	emailPort := Email.EmailPort
	emailServer := Email.EmailServer
	emailFrom := EmailReg.EmailFrom
	emailSubject := EmailReg.EmailSubject
	emailFormat := EmailReg.EmailFormat
	// pass := EmailPass
	*code = GenerateRandomNumberString(6)
	dur, err := strconv.Atoi(config.ConfigMap["email_reg"].(*config.EmailReg).Duration)
	if err != nil {
		return err
	}
	*duration = dur
	emailFormat = strings.Replace(emailFormat, "[code]", *code, 1)
	msg := "From: " + emailFrom + "\n" +
		"To: " + to + "\n" +
		"Subject: " + emailSubject + "\n\n" +
		"" + emailFormat + ""

	auth := smtp.PlainAuth("", emailUser, emailPass, emailServer)
	if emailTLS == "1" {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         emailServer,
		}

		conn, err := tls.Dial("tcp", emailServer+":"+emailPort, tlsconfig)
		if err != nil {
			return err
		}

		client, err := smtp.NewClient(conn, emailServer)
		if err != nil {
			return err
		}

		if err = client.Auth(auth); err != nil {
			return err
		}

		if err = client.Mail(emailUser); err != nil {
			return err
		}

		if err = client.Rcpt(to); err != nil {
			return err
		}

		w, err := client.Data()
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(msg))
		if err != nil {
			return err
		}

		err = w.Close()
		if err != nil {
			return err
		}

		err = client.Quit()
		if err != nil {
			return err
		}
		return nil

	} else {
		err := smtp.SendMail(emailServer+":"+emailPort, auth, emailUser, []string{to}, []byte(msg))
		return err

	}

}
