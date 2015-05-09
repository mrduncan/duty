# Duty

Duty is a command line [PagerDuty](http://pagerduty.com) client.

## Usage

You'll need to have the following environment variables set:
- `PAGERDUTY_SUBDOMAIN` is the subdomain of your account (http://<subdomain>.pagerduty.com).
- `PAGERDUTY_API_KEY` is your API key, create a new one at https://<subdomain>.pagerduty.com/api_keys.

``` sh
$ duty
usage: duty <command>

Available commands are:
  help        Show usage
  incidents   List incidents
  schedules   List schedules
  users       List users
```

``` sh
$ duty users
ABCDEFG
Name:      Jane Doe
Email:     jane@example.com
Role:      owner
Timezone:  Pacific Time (US & Canada)

HIJKLMN
Name:      John Doe
Email:     john@example.com
Role:      user
Timezone:  Pacific Time (US & Canada)
```

## Disclaimer

I originally threw this together in about an hour so it's very incomplete.  Pull requests welcome!
