package mail_analyzer

import (
  "time"
  "testing"
  . "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type AnalyzerTests struct {
  analyzer *MailAnalyzer
}

var _ = Suite(&AnalyzerTests{})

func (t *AnalyzerTests) SetUpTest(c *C) {
  t.analyzer = &MailAnalyzer{"jeffersongarzonm@gmail.com", "auth.json", []Mail{}}
}

func (t *AnalyzerTests) TestGetMailFromData(c *C) {
  now := time.Now()
  mail := t.analyzer.GetMailFromData("leorock64@gmail.com", "Almojabanas a mil", "vendo empanadas", now)

  c.Check(t.analyzer.mails.length, Equals, 1)

  c.Check(mail.sender, Equals, "leorock64@gmail.com")
  c.Check(mail.subject, Equals, "Almojabanas a mil")
  c.Check(mail.content, Equals, "vendo empanadas")
  c.Check(mail.time, Equals, now)
}

func (t *AnalyzerTests) TestGetStringCSV(c *C) {
  sender := "leo@gmail.com"
  subject := "hay tareas?"
  content := "mano hay tarea pa mañana"
  time := time.Now()

  mail := Mail{sender, subject, content, time, []string{}, []string{}, []string{}}

  mail.AddRecipient("jeffersongarzonm@gmail.com")
  mail.AddRecipient("minenequerido@gmail.com")
  mail.AddCC("juliansito@fluvip.com")
  mail.AddCCO("mariita@fluvip.com")

  t.analyzer.AddMail(mail)

  csv := t.analyzer.GetStringCSV()
  realCSV := "Sender,Subject,Content,Date,Recipients,CC,CCO\n" +
              mail.sender + "," + mail.subject + "," + mail.content + "," + mail.date.String() +
              ",\"jeffersongarzon@gmail,minenequerido@gmail.com,\",\"juliansito@fluvip.com,\",\"marrita@fluvip.com\"\n"

  c.Assert(csv, Equals, realCSV)
}

