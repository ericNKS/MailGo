package service

import (
	"fmt"
	"sendMail/internal/model"

	gomail "gopkg.in/mail.v2"
)

func ExecuteSendMail(r *model.Mail, m *model.Message, destinatario string) {
	message := gomail.NewMessage()

	message.SetHeader("From", r.From)
	message.SetHeader("To", destinatario)
	message.SetHeader("Subject", m.Subject)

	// message.AddAlternative("text/html", `
	//     <html>
	//         <body>
	//             <h1>This is a Test Email</h1>
	//             <p><b>Hello!</b> This is a test email with HTML formatting.</p>
	//             <p>Thanks,<br>Mailtrap</p>
	//         </body>
	//     </html>
	// `)

	message.AddAlternative("text/html", m.Body)

	dialer := gomail.NewDialer(r.SmtpHost, int(r.SmtpPort), r.From, r.Password)

	// Send the email
	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
	}
}
