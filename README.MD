# BOT SEND EMAIL AUTOMATICALLY

นี่คือโปรเจคที่ใช้สำหรับการส่งอีเมล์อัตโนมัติ โดยใช้ `Golang` และ `Google API` ในการส่งอีเมล์ โดยจะนำข้อมูลจากไฟล์ Excel มาส่งอีเมล์ไปยังผู้รับ

## Usage
- สร้างไฟล์ `.env` ในโฟลเดอร์ `config` โดยมีรูปแบบดังนี้
```env
smtp_Host=smtp.gmail.com
smtp_Port=587
sender_Email=
sender_Password=
subject_Email=สวัสดีครับนี่คือการทดสอบ
body_Email=สวัสดีครับนี่คือการทดสอบ โปรดอย่าตอบกลับ ขอบคุณครับ
```

- `sender_Password` ที่ได้จาก [AppPasswords](https://myaccount.google.com/apppasswords) ให้ใส่เป็น `sender_Password`


