# Send discord alerts from HTTP requests



This project is made to show how to send discord alters from http requests. This can be used to quickly recieve importand messages from log-sources, or that can be triggered in other importand cases.

You can connect this alert system to:
1 - GrayLog
2 - ELK
3 - Any application with http library

This bot currently accepts messages in `.json` format, but it is possible to reconfigure it for another input type.


# Setup process:

1) Create discord bot - go to [discord applications](https://discord.com/developers/applications) and create new application
2) Go to OAuth2, and create URL, using URL-generator (add permission to write messages)
3) Open generated URL in your browser, and add a bot to your server
4) Go to bot's page in discor applicaiton