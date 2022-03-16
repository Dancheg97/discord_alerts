## <p  align="center" style="font-family:courier;font-size:180%" size=212px> DISCORD ALERTS  </p> 

[![Generic badge](https://img.shields.io/badge/LICENSE-MIT-orange.svg)](LICENSE)
[![Generic badge](https://img.shields.io/badge/DOCKER-HUB-blue.svg)](https://hub.docker.com/repository/docker/dangdancheg/discord_alerts)
[![Generic badge](https://img.shields.io/badge/SWAGGER-1.1.0-green.svg)](https://app.swaggerhub.com/apis/Dancheg97/DISCORD_ALERST/1.0.0)


<p align="center">
<a href='https://go.dev/'>
<img go align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="138px" height="138px" src="https://asterisk-pbx.ru/wiki/_media/asterisk/ari/swaggerlogo360.png" /> 
</a>
<a href='https://www.python.org/'>
<img python align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="138px"  height="82px" src="https://thypix.com/wp-content/uploads/2020/04/white-arrow-92.png" />
</a>
<a href='https://docs.microsoft.com/en-us/dotnet/csharp/'>
<img c# align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="138px"  height="138px" src="
https://pnggrid.com/wp-content/uploads/2021/05/Discord-Logo-Circle-1024x1024.png" />
</a>
</p>




This project is made to show how to send discord alters from http requests. This can be used to quickly recieve importand messages from log-hubs, or orther sources.

You can connect this alert system to:
1 - GrayLog
2 - ELK
3 - Any application with http library

This bot currently accepts input in `json` format, but it is possible to reconfigure it for another input type.


# Setup process:

1) Create discord bot - go to [discord applications](https://discord.com/developers/applications) and create new application
2) Go to OAuth2, and create URL, using URL-generator (add permission to write messages)
3) Open generated URL in your browser, and add a bot to your server
4) Copy the token from `Bot` page in discord application, and paste it to your `.env` or `docker-compose` file to field `TOKEN`
5) Copy channel id from your discord channel to send messages to, and paste it to your `CHANNEL_ID` field in `.env` or `docker-compose` file

Launch the app, it is ready to go!

# Docker compose setup:

Example of `docker-compose.yml` file:

```yaml
version: '3'
services:
  discord_alerts:
    image: dangdancheg/discord_alerts:latest
    ports:
      - 8092:8092
    environment:
      - TOKEN=PASTE_YOUR_DISCORD_TOKEN
      - CHANNEL_ID=PASTE_CHANNELID_FROM_CHANNEL_LINK
```

# GrayLog stup:

