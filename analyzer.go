package mail_analyzer

import (
  "time"
)

type Mail struct {
  sender string
  subject string
  content string
  date time.Time
  recipients [] string
  cc [] string
  cco [] string
}

type MailAnalyzer struct {
  email string
  authfilename string
  mails [] Mail
}

