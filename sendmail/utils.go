package sendmail

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/resend/resend-go/v3"
)

type Utils interface {
	SendEmail(content string) error
}

type UtilsImpl struct {
	Client *resend.Client
}

// todo 构造函数
func NewUtils(apikey string) *UtilsImpl {

	apikey := os.Getenv(apikey)
	if apikey == "" {
		log.Println("APIKEY为空")
	}
	client := resend.NewClient(apikey)
	return &UtilsImpl{
		Client: client,
	}
}

func (u *UtilsImpl) SendEmail(content string) error {
	params := &resend.SendEmailRequest{
		From:        "Acme <onboarding@resend.dev>",
		To:          []string{"tattoo186225345941@gmail.com"},
		Subject:     "System Status",
		Bcc:         nil,
		Cc:          nil,
		ReplyTo:     "",
		Html:        fmt.Sprintf("<h1>%s</h1>\n<p>%s</p>", content, time.Now().Format("2006-01-02 15:04:05")),
		Text:        "",
		Tags:        nil,
		Attachments: nil,
		Headers:     nil,
		ScheduledAt: "",
		Template:    nil,
	}

	send, err := u.Client.Emails.Send(params)
	if err != nil {
		return err
	}
	fmt.Println(send.Id)
	return nil
}
