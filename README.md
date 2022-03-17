# <p  align="center" style="font-family:courier;font-size:180%" size=212px> DISCORD ALERTS  </p> 

[![Generic badge](https://img.shields.io/badge/LICENSE-MIT-orange.svg)](LICENSE)
[![Generic badge](https://img.shields.io/badge/DOCKER-HUB-blue.svg)](https://hub.docker.com/repository/docker/dangdancheg/discord_alerts)
[![Generic badge](https://img.shields.io/badge/SWAGGER-API-green.svg)](https://app.swaggerhub.com/apis/Dancheg97/DISCORD_ALERST/1.0.0)


<p align="center">
<img go align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="138px" height="138px" src="https://asterisk-pbx.ru/wiki/_media/asterisk/ari/swaggerlogo360.png" /> 
<img python align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="68px"  height="42px" src="https://thypix.com/wp-content/uploads/2020/04/white-arrow-92.png" />
<img c# align="center" style="padding-left: 10px; padding-right: 10px; padding-bottom: 10px;" width="138px"  height="138px" src="https://cdn.logojoy.com/wp-content/uploads/20210422095037/discord-mascot.png" />
</p>




This project is made to show how to send discord alters from http requests. This can be used to quickly recieve importand messages from log-hubs, or orther sources.

You can connect this alert system to:
- GrayLog
- ELK
- Any application with http library

This bot currently accepts input in `json` format, but it is possible to reconfigure it for another input type.


# Bot setup process:

1) Create discord bot - go to [discord applications](https://discord.com/developers/applications) and create new application

![discord/app](assets/1.png)
![discord/app](assets/2.png)

2) Go to `bot`, and press `add bot`

![discord/app](assets/3.png)
![discord/app](assets/4.png)

3) Press reset token, and save token in some safe place (definetly .txt file, called tokens or passwords :D)

![discord/app](assets/5.png)
![discord/app](assets/6.png)

5) Go to URL-Genrator, and generate URL with required permissions

![discord/app](assets/7.png)
![discord/app](assets/8.png)

6) Open that URL in browser on new page, add bot to your server

![discord/app](assets/9.png)
![discord/app](assets/10.png)
![discord/app](assets/11.png)

7) Go to your discord channel, and extract channel id from channel link

![discord/app](assets/12.png)
![discord/app](assets/13.png)

8) Paste your token and channel id to docker-compose or .env file depending on your build type

![discord/app](assets/14.png)

Launch the bot, and enjoy!

# Docker compose example:

Example of `docker-compose.yml` file:

```yaml
version: '3'
services:
  discord_alerts:
    image: dangdancheg/discord_alerts:latest
    ports:
      - 3458:3458
    environment:
      - TOKEN=PASTE_YOUR_DISCORD_TOKEN
      - CHANNEL_ID=PASTE_CHANNELID_FROM_CHANNEL_LINK
      - PORT=3458
```

# ENV file example:

Example of `.env` file, to build from source:

```python
TOKEN=YOUR_TOKEN
CHANNEL_ID=YOUR_CHANNEL_ID
```

# Test via http:

You can send simple request using http library:


# Test from graylog: