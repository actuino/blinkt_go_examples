## Blink Unicorn Client

A Go client for the [Unicorn Display Server](https://github.com/actuino/unicorn-display/tree/master/server)

Gets commands from the server in real-time via WebSockets and updates it's Blinkt! display.

## How to run via Docker

* Run a Unicorn server on your network : see [Readme](https://github.com/actuino/unicorn-display/blob/master/server/README.md)
* Pull and run the Docker image on the Pi:
```
docker pull actuino/blinkt-unicorn-client:1
docker run --privileged -d -e DISPLAY_SERVER_HOST=192.168.7.3 actuino/blinkt-unicorn-client:1
```
(replace `192.168.7.3` by the actual IP or hostname of the display server)
If your server is listening on another port than the default 80, use `-e DISPLAY_SERVER_PORT=port_number` to specify the right one.

* Point your browser to the IP of the display server. The first of the 8 lines will be pushed to the Blinkt!

## Protocol

Still WIP. See [Unicorn Display Protocol](https://github.com/actuino/unicorn-display/blob/master/doc/PROTOCOL.md)
This client only supports the 'StaticId' format for now.

## More info

You'll find a [blog post on Actuino](http://www.actuino.fr/raspi/).
