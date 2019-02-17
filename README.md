# pusher-channel-terminal-web-sync

How to make use of Pusher Channels to show the output of your terminal in realtime.
Think of something similar to how your CI server would display the progress of a build in realtime.

#### Getting started

- Clone this repository `git clone git@github.com:adelowo/pusher-channel-terminal-web-sync.git`.
- Update `.env` to contain your original credentials.
- Update `PUSHER_KEY` and `PUSHER_CLUSTER` in L2 - L3 of [app.js](https://github.com/adelowo/pusher-encrypted-channels/blob/master/app.js#L2-L3)
- Run any command on the command line such as `yoursampleprogram | go run main.go`

## Built With

- [Pusher Channels](https://pusher.com/channels) - APIs to enable devs building realtime features.
- Golang
