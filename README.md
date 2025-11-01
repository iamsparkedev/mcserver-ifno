##     :globe_with_meridians:  MCserver-Info
  A lightweight API service built in **Go** to fetch real-time information about any Minecraft server — including version, player count, MOTD, and online status.

![Contributors](https://img.shields.io/github/contributors/iamsparkedev/mcserver-info)
![Forks](https://img.shields.io/github/forks/iamsparkedev/mcserver-info?style=social)
![Stars](https://img.shields.io/github/stars/iamsparkedev/mcserver-info?style=social)
![License](https://img.shields.io/github/license/iamsparkedev/mcserver-info)
-----
[![Discord](https://discordapp.com/api/guilds/1418891584732266669/widget.png?style=banner2)](https://discord.gg/UgQHxrCk2z)

> [!CAUTION]
>
> This Project Is Under Development
> , Some features might break


# API Endpoints & Usage
### Server Lookup
``` /api/server?ip=ip:port```

Example : `/api/server/?ip=mc.hypixel.net:25565`

Output :``` {
    "ip": "mc.hypixel.net",
    "version": "Requires MC 1.8 / 1.21",
    "online": true,
    "players": {
        "online": 28074,
        "max": 200000
    },
    "motd": "           §aHypixel Network §c[1.8/1.21]\n           §d§lSB 0.23.5 §7- §6§lHALLOWEEN EVENT"
} ```

### Server List
``` /api/list?id=randomid```

Example : `api/list?id=3`


Output : ``` {
    "server_id": "3",
    "name": "Hypixel Network",
    "ip_address": "mc.hypixel.net",
    "gamemodes": [
        "Bedwars",
        "SkyWars",
        "SkyBlock",
        "Murder Mystery",
        "Minigames"
    ],
    "description": "One of the largest and most famous Minecraft servers, offering a massive selection of original minigames."
} ```

### 

## BUILT WITH
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=fff)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=fff)
