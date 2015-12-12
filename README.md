# Could Foundry Go Cron App
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

If the job has other language dependencies make sure to list their buildpacks in the `.buildpacks` yaml
```yaml
# An app the run an npm package
https://github.com/cloudfoundry/nodejs-buildpack
https://github.com/cloudfoundry/go-buildpack.git
```

Push to Cloud Foundry
`cf push`
