package test

import (
	"fmt"
	"net"
	"net/smtp"
	"time"
)

func Main2() {
	// ตั้งค่า IP หรือ hostname ของเครื่อง PC ที่ต้องการตรวจสอบ
	host := "172.22.97.55"
	port := "8080"
	address := fmt.Sprintf("%s:%s", host, port)

	// ตั้งค่าอีเมลสำหรับการแจ้งเตือน
	smtpServer := "smtp.example.com"
	smtpPort := "587"
	sender := "your-email@example.com"
	password := "your-email-password"
	receiver := "receiver-email@example.com"

	// ตั้งค่าเวลาตรวจสอบ (เช่น ทุก 5 นาที)
	checkInterval := 1 * time.Minute

	for {
		if !isOnline(address) {
			sendEmail(smtpServer, smtpPort, sender, password, receiver, host)
		}
		time.Sleep(checkInterval)
	}
}

// ฟังก์ชันตรวจสอบการเชื่อมต่อ
func isOnline(address string) bool {
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Printf("เครื่อง %s ไม่ออนไลน์\n", address)
		return false
	}
	conn.Close()
	fmt.Printf("เครื่อง %s ออนไลน์\n", address)
	return true
}

// ฟังก์ชันส่งอีเมลแจ้งเตือน
func sendEmail(server, port, sender, password, receiver, host string) {
	auth := smtp.PlainAuth("", sender, password, server)
	msg := []byte("To: " + receiver + "\r\n" +
		"Subject: แจ้งเตือนเครื่อง PC ไม่ออนไลน์\r\n" +
		"\r\n" +
		"เครื่อง " + host + " ไม่ออนไลน์.\r\n")

	err := smtp.SendMail(server+":"+port, auth, sender, []string{receiver}, msg)
	if err != nil {
		fmt.Printf("ไม่สามารถส่งอีเมลแจ้งเตือนได้: %v\n", err)
	} else {
		fmt.Println("ส่งอีเมลแจ้งเตือนสำเร็จ")
	}
}
