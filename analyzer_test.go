package mail_analyzer_test

import (
  . "MailAnalyzer"
  "time"
  "testing"
  . "gopkg.in/check.v1"
)

func TestAnalyzer(t *testing.T) { TestingT(t) }

type AnalyzerTests struct {
}

type AnalyzerMock struct {
  Analyzer
}

func (a *AnalyzerMock) Connect(email string, secretFile string) {
  a.Client = &MailClientMock{}
}

type MailClientMock struct {
  MailClient
}

func (c *MailClientMock) GetList() []ListItem {
  return  []ListItem  {
    ListItem {id: "12345678", threadId: "abc"},
    ListItem {id: "87654321", threadId: "cba"},
  }
}

func (c *MailClientMock) GetMailDataById(id string) Mail {
  return Mail {
    id: id,
    sender: "test@gmail.com",
    subject: "Testing",
    content: "Just a Test",
  }
}

var _ = Suite(&AnalyzerTests{})

func (t *AnalyzerTests) SetUpTest(c *C) {
  t.analyzer = &AnalyzerMock{}
}

func (t *AnalyzerTests) TestConnect(c *C) {
  t.analyzer.Connect("jeffersongarzonm@gmail.com", "client_secret.json")
  c.Check(t.analyzer.client, Not(Equals), nil)
}

func (t *AnalyzerTests) TestAddMail (c *C) {
  now := time.Now()
  mail := Mail{
    sender: "leorock64@gmail.com",
    subject: "Almojabanas a mil",
    content: "vendo empanadas",
    date: now,
  }

  t.analyzer.AddMail(mail)

  c.Check(len(t.analyzer.mails), Equals, 1)

  response := t.analyzer.mails[0]

  c.Check(response.sender, Equals, "leorock64@gmail.com")
  c.Check(response.subject, Equals, "Almojabanas a mil")
  c.Check(response.content, Equals, "vendo empanadas")
  c.Check(response.date, Equals, now)

}

func (t *AnalyzerTests) TestLoadMails(c *C) {
  t.analyzer.Connect()
  t.analyzer.LoadMails()

  c.Check(len(t.analyzer.email), Equals, 2)

  c.Check(t.analyzer.emails[0].id, Equals, "12345678")
  c.Check(t.analyzer.emails[1].id, Equals, "87654321")
}

func (t *AnalyzerTests) TestGetStringCSV(c *C) {
  sender := "leo@gmail.com"
  subject := "hay tareas?"
  content := "mano hay tarea pa mañana"
  time := time.Now()

  mail := Mail{sender: sender, subject: subject, content: content, date: time}

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

