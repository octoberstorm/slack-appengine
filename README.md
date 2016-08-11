# Slack [![GoDoc](https://godoc.org/github.com/rickt/slack?status.png)](https://godoc.org/github.com/rickt/slack)

Golang client for the Slack API, specifically for backend apps running Google App Engine. Forked from https://github.com/bluele/slack.

Modifications by Rick Tait:

* added Slack usergroup support 
* added Slack user presence support 
* HTTP request changes to work w/Google App Engine. HTTP requests are restricted/controlled in App Engine, you have to use an App Engine context-friendly helper function. See https://cloud.google.com/appengine/docs/go/urlfetch/ for more details

## Currently supports:

All the Slack API methods that https://github.com/bluele/slack supports, plus:

Method | Description | Example
--- | --- | ---
usergroups.list | Lists all user groups in a Slack team.  | [#link](https://api.slack.com/methods/usergroups.list)
users.getPresence | Shows user presence/online status. | [#link](https://api.slack.com/methods/users.getPresence)

## Command line tool

If you are looking for slack commandline utility, [vektorlab/slackcat](https://github.com/vektorlab/slackcat) probably suits you.

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>

# Google App Engine & other patches by

**Rick Tait**

* <http://github.com/rickt>
* <rickt@rickt.org>
