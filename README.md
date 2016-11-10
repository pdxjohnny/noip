# No-IP Dynamic DNS Client [![Build Status](https://travis-ci.org/pdxjohnny/noip.svg?branch=master)](https://travis-ci.org/pdxjohnny/noip)

No-IPs client is really old from what I can tell so it probably has a ton of
CVEs in it. On a brighter note this uses go so if you've got the latest version
of go then you should be ok using this client. It looks for environment
variables and then sends them to No-IPs api every ten minutes to keep your
domain pointing in the right place.

## Environment Variables - In docker-compose file for ease of use

```yaml
version: '2'

services:

  noip:
    image: pdxjohnny/noip
    restart: always
    environment:
      email: username
      username: username
      password: password
      hostname: comma.seperated,list.of,host.names
```
