# Cloud Foundry Go Cron App
An application for running cron jobs in a Cloud Foundry app.

# Usage
Create a cron-tab.yaml
`cp cron-tab.yaml.example cron-tab.yaml`

Declare jobs in the following format
```yaml
jobs:
  - name: Job name 1
    schedule: "* * * * * *"
    command: "echo Hello World!"
  - name: Job name 2
    schedule: "*/1 * * * * *"
    command: "echo Hello World!"
  - name: Job name 3
    schedule: "*/3 * * * * *"
    command: "echo Hello World!"
```

Push to Cloud Foundry
`cf push`
