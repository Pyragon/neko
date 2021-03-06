<div align="center">
  <a href="https://n.eko.moe/#/" ><img src="https://raw.githubusercontent.com/nurdism/neko/master/docs/_media/logo.png" width="450" height="auto"/></a>
  <p align="center">
    <img src="https://img.shields.io/github/v/release/nurdism/neko" alt="release">
    <img src="https://img.shields.io/github/license/nurdism/neko" alt="license">
    <img src="https://img.shields.io/docker/pulls/nurdism/neko" alt="pulls">
    <img src="https://img.shields.io/github/issues/nurdism/neko" alt="issues">
    <a href="https://discord.gg/3U6hWpC" ><img src="https://discordapp.com/api/guilds/665851821906067466/widget.png" alt="Chat on discord"><a/>
    <a href="https://github.com/nurdism/neko/actions" ><img src="https://github.com/nurdism/neko/workflows/deploy/badge.svg" alt="build"><a/>
  </p>
  <br/>
  <br/>
  <img src="https://i.imgur.com/ZSzbQr7.gif" width="650" height="auto"/>
  <br/>
  <br/>
</div>

# n.eko (m1k1o fork)
This app uses Web RTC to stream a desktop inside of a docker container. This is fork of https://github.com/nurdism/neko.

## Differences to original repository.

### New Features
- Clipboard button with text area - for browsers, that don't support clipboard syncing or for HTTP.
- Keyboard modifier state synchronization (Num Lock, Caps Lock, Scroll Lock) for each hosting.
- Added chromium ungoogled (with h265 support) an kept up to date by @whalehub.
- Added Picture in Picture button (only for watching screen, controlling not possible).
- Added RTMP broadcast. Enables broadcasting neko screen to local RTMP server, YouTube or Twitch.
- Stereo sound (works properly only in Firefox host).

### Bugs
- Fixed minor gst pipeline bug.
- Locked screen only for users, admins can still join.

### Misc
- Custom docker workflow.
- Based on debian buster instead of stretch.
- Custom avatars without any 3rd party depenency.
- Ignore duplicate notify bars.
- No pointer events for notify bars.
- Disable debug mode by default.

# Getting started & FAQ

Use following docker images:
- `m1k1o/neko:latest` - for Firefox.
- `m1k1o/neko:chromium` - for Chromium Ungoogled (needs `--cap-add=SYS_ADMIN`).
- `m1k1o/neko:base` - for custom base.

Networking:
- If you want to use n.eko in **external** network, you can omit `NEKO_NAT1TO1`. It will automatically get your Public IP.
- If you want to use n.eko in **internal** network, set `NEKO_NAT1TO1` to your local IP address (e.g. `NEKO_NAT1TO1: 192.168.1.20`)-

Why so many ports?
- WebRTC needs UDP ports for each channel it creates towards users.
- Every user will need 2 UDP ports (for getting audio/video and sending mouse positions).
- You can freely limit number of UDP ports. But you can't map them to diferent ports.
  - This **WONT** work: `32000-32100:52000-52100/udp`
- You can change API port (8080).
  - This **WILL** work: `3000:8080`

Behind reverse proxy?
- Nginx configuration: https://github.com/nurdism/neko/issues/111#issuecomment-742656957
- Apache configuration: https://github.com/nurdism/neko/blob/cad98a62a5bd7f1daf2c11980631bb14ba81a1f6/docs/apache-proxypass-config.md#example-apache-config
- Traefik configuration: https://github.com/m1k1o/neko-vpn/blob/a1b934515dcf597992a515d61d307c2450a11002/docker-compose.yml#L38-L43

Want to use VPN for your neko browsing?
- Check this out: https://github.com/m1k1o/neko-vpn

Accounts:
- There are no accounts, display name (a.k.a. username) can be freely chosen. Only paword needs to match. Depeding on which password matches, visitor gets its privilege:
  - Anyone, who enters with `NEKO_PASSWORD` will be **user**.
  - Anyone, who enters with `NEKO_PASSWORD_ADMIN` will be **admin**.

Screen size
- Only admins can change screen size.
- You can set default screen size, but this size **MUST** be one from list, that your server supports.
- You will get this list in frontend, where you can choose from.

## Firefox

```yaml
version: "3.4"
services:
  neko:
    image: "m1k1o/neko:latest"
    restart: "unless-stopped"
    shm_size: "2gb"
    ports:
      - "8080:8080"
      - "52000-52100:52000-52100/udp"
    environment:
      NEKO_SCREEN: '1920x1080@30'
      NEKO_PASSWORD: neko
      NEKO_PASSWORD_ADMIN: admin
      NEKO_EPR: 52000-52100
      NEKO_NAT1TO1: <your-IP>
```

## Chromium Ungoogled

```yaml
version: "3.4"
services:
  neko:
    image: "m1k1o/neko:chromium"
    restart: "unless-stopped"
    shm_size: "2gb"
    ports:
      - "8080:8080"
      - "52000-52100:52000-52100/udp"
    cap_add:
      - SYS_ADMIN
    environment:
      NEKO_SCREEN: '1920x1080@30'
      NEKO_PASSWORD: neko
      NEKO_PASSWORD_ADMIN: admin
      NEKO_EPR: 52000-52100
      NEKO_NAT1TO1: <your-IP>
```
