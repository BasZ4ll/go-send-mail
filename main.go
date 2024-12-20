package main

import (
	"fmt"
	"github/BasZ4ll/go-send-email/test"
	"log"
	"net/smtp"
	"os"

	"path/filepath"
	//"net/smtp"

	"github.com/joho/godotenv"
	"github.com/tealeg/xlsx"
)

func main() {
	// เรียกใช้ func ที่ต้องการ test
	test.Main2()
	SendMail()
}

func SendMail() {

	godotenv.Load()

	// ดึงข้อมูลจากไฟล์ Excel
	files, err := filepath.Glob("*.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	for _, excelFileName := range files {
		xlFile, errD := xlsx.OpenFile(excelFileName)
		if errD != nil {
			log.Fatalf("ไม่สามารถเปิดไฟล์ Excel: %s", errD)
		}

		// ข้อมูลเข้าสู่ระบบ SMTP ของ Google
		smtpHost := os.Getenv("smtp_Host")             //"smtp.gmail.com
		smtpPort := os.Getenv("smtp_Port")             //587
		senderEmail := os.Getenv("sender_Email")       // แทนที่ด้วยอีเมลของคุณ
		senderPassword := os.Getenv("sender_Password") // แทนที่ด้วยรหัสผ่านของคุณ

		// ดึงข้อมูลจากแต่ละแถวและคอลัมน์
		for _, sheet := range xlFile.Sheets {
			for _, row := range sheet.Rows[1:] {
				fmt.Println("Row: ", row)
				email := row.Cells[0].String()
				fmt.Println("email:", email)
				name := row.Cells[1].String()
				fmt.Println("name:", name)

				recipients := []string{email}
				fmt.Println("recipients:", recipients)
				// ส่งอีเมลผ่าน SMTP
				subject := os.Getenv("subject_Email") // แทนที่ด้วยหัวข้อของอีเมลของคุณ
				body := os.Getenv("body_Email")       // แทนที่ด้วยเนื้อหาของอีเมลของคุณ

				// กำหนดการเชื่อมต่อ SMTP
				auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

				// กำหนดเนื้อหาของอีเมล
				message := []byte("To: " + email + "\r\n" + "Subject: " + subject + "\r\n" + "\r\n" + body)

				// ส่งอีเมลผ่าน SMTP
				err := smtp.SendMail(smtpHost+":"+fmt.Sprint(smtpPort), auth, senderEmail, recipients, message)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println()
			}
		}
		fmt.Println("Email sent for file: ", excelFileName)
	}
	fmt.Println("Email sent successfully!")
}
