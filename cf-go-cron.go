package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/robfig/cron"
	"gopkg.in/yaml.v2"
)

type cronjob struct {
	Name     string `yaml:"name"`
	Schedule string `yaml:"schedule"`
	Command  string `yaml:"command"`
}

// CronTab is a collection of jobs
type CronTab struct {
	Jobs []cronjob `yaml:"jobs"`
}

func loadCronJobs() CronTab {
	var jobs CronTab
	data, err := ioutil.ReadFile("cron-tab.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	if yaml.Unmarshal(data, &jobs) != nil {
		log.Fatalf("error: %v", err)
	}
	return jobs
}

func stringToCommand(stringCommand string) *exec.Cmd {
	parts := strings.Fields(stringCommand)
	cmd := exec.Command(parts[0], parts[1:]...)
	return cmd
}

func main() {
	jobs := loadCronJobs()
	c := cron.New()
	for _, job := range jobs.Jobs {
		func(job cronjob) {
			c.AddFunc(job.Schedule, func() {
				log.Println("Running: " + job.Name)
				cmd := stringToCommand(job.Command)
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			})
		}(job)
	}
	c.Start()

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
