package main

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/appleboy/drone-template-lib/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type (
	// Config for the plugin.
	Config struct {
		Debug        bool
		Webhook      []string
		Message      string
		TemplateFile string
	}
	// Repo information.
	Repo struct {
		FullName  string
		Link      string
		Namespace string
		Name      string
	}

	// Commit information.
	Commit struct {
		Sha          string
		Ref          string
		Branch       string
		Link         string
		Message      string
		AuthorName   string
		AuthorEmail  string
		AuthorAvatar string
	}

	// Build information.
	Build struct {
		Tag      string
		Event    string
		Number   int
		Status   string
		Link     string
		Started  int64
		Finished int64
		PR       string
		DeployTo string
	}

	// Plugin information.
	Plugin struct {
		Config Config
		Repo   Repo
		Commit Commit
		Build  Build
	}
)

const defaultMessage = `
{
  "msg_type":"interactive",
  "card":{
    "config":{
      "wide_screen_mode":true,
      "enable_forward":true
    },
    "header":{
      "title":{
        "tag":"plain_text",
        "content":"{{ Repo.FullName }}"
      }
    },
    "elements":[
      {
        "tag":"markdown",
        "content":"{{#success Build.Status }}✅{{/success}}{{#failure Build.Status}}❌{{/failure}} Build [#{{ Build.Number }}]({{ Build.Link }}) {{ Build.Status }}.\n📝 Commit by {{ Commit.AuthorName }} on {{ Commit.Branch }}:\n{{ Commit.Message }}"
      },
      {
        "tag":"note",
        "elements":[
          {
            "tag":"plain_text",
            "content":"@{{ datetime Build.Started '2006-01-02 15:04:05' '' }} to @{{ datetime Build.Finished '2006-01-02 15:04:05' '' }}."
          }
        ]
      }
    ]
  }
}`

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func preprocessCommitMessage(commitMessage string) string {
	commitMessage = strings.TrimRight(commitMessage, "\n")
	return trimQuotes(strconv.Quote(commitMessage))
}

func loadMessageFromFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	message, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(message), nil
}

func (p Plugin) extractMessage() (string, error) {
	var message string
	if len(p.Config.TemplateFile) > 0 {
		message, _ = loadMessageFromFile(p.Config.TemplateFile)
	} else if len(p.Config.Message) > 0 {
		message = p.Config.Message
	}

	if len(message) == 0 {
		message = defaultMessage
	}

	var err error
	message, err = template.RenderTrim(message, p)
	if err != nil {
		return "", err
	}

	return message, nil
}

func (p Plugin) send(client http.Client, webhook, message string) error {
	resp, err := client.Post(webhook, "application/json", bytes.NewBuffer([]byte(message)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if p.Config.Debug {
		respBody, _ := ioutil.ReadAll(resp.Body)
		log.Println("================================")
		log.Printf("Request Body: \n%s\n", message)
		log.Printf("Response Body: %#v\n", string(respBody))
		log.Println("================================")
	}

	return nil
}

// Exec executes the plugin.
func (p Plugin) Exec() error {
	if len(p.Config.Webhook) == 0 {
		return errors.New("missing webhook url")
	}

	// initial client
	client := http.Client{}

	// extract message
	message, err := p.extractMessage()
	if err != nil {
		return err
	}

	// send message
	for _, webhook := range p.Config.Webhook {
		if err := p.send(client, webhook, message); err != nil {
			return err
		}
	}

	return nil
}
