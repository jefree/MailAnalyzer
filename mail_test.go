package mail_analyzer_test

import (
  . "MailAnalyzer"
  "time"
  "testing"
  . "gopkg.in/check.v1"
)

func TestMail(t *testing.T) { TestingT(t) }

type MailTests struct {
  mail Mail
}

var _ = Suite(&MailTests{})

func (t *MailTests) SetUpTest(c *C) {
  sender := "leo@gmail.com"
  subject := "hay tareas?"
  content := "mano hay tarea pa mañana"
  time := time.Now()

  t.mail = Mail{sender: sender, subject: subject, content: content, date: time}
}

func (t *MailTests) TestCreateMail(c *C) {
  sender := "jeffersongarzonm@gmail.com"
  subject := "Almojabanas a mil"
  content := "Vendo empanadas"
  date := time.Now()

  mail := CreateMail(sender, subject, content, date)

  c.Assert(mail.sender, Equals, sender)
  c.Assert(mail.subject, Equals, subject)
  c.Assert(mail.content, Equals, content)
  c.Assert(mail.date, Equals, date)
  c.Assert(len(mail.recipients), Equals, 0)
  c.Assert(len(mail.cc), Equals, 0)
  c.Assert(len(mail.cco), Equals, 0)
}

func (t *MailTests) TestAddRecipient(c *C) {
  mail := t.mail

  mail.AddRecipient("yonosoydeverdad@fake.com")
  mail.AddRecipient("yotampoco@fake.com")

  c.Assert(len(mail.recipients), Equals, 2)
  c.Assert(mail.recipients[0], Equals, "yonosoydeverdad@fake.com")
  c.Assert(mail.recipients[1], Equals, "yotampoco@fake.com")
}

func (t *MailTests) TestAddCC(c *C) {
  mail := t.mail

  mail.AddCC("yonosoydeverdad@fake.com")
  mail.AddCC("yotampoco@fake.com")

  c.Assert(len(mail.cc), Equals, 2)
  c.Assert(mail.cc[0], Equals, "yonosoydeverdad@fake.com")
  c.Assert(mail.cc[1], Equals, "yotampoco@fake.com")
}

func (t *MailTests) TestAddCCO(c *C) {
  mail := t.mail

  mail.AddCCO("yonosoydeverdad@fake.com")
  mail.AddCCO("yotampoco@fake.com")

  c.Assert(len(mail.cco), Equals, 2)
  c.Assert(mail.cco[0], Equals, "yonosoydeverdad@fake.com")
  c.Assert(mail.cco[1], Equals, "yotampoco@fake.com")
}
