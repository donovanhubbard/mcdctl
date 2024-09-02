# mcdctl
A memcached TUI

## Todo

* ~~Get a rudimentary UI~~
* ~~Connect to memcached~~
* Make the UI pretty
* ~~For some reason the model's memcache client is being set to nil.~~
* ~~Refactor the huge list of `strings.HasPrefix` to something better in the client.go file~~
* Make the client aware of commands so you can avoid having to type the size of the commands
* Have shift+enter place a newline and not send the commands
* press up to repeat the last command 
* Scroll back to see previous commands
* Delete text once there are too many lines in the command history
* ~~Create a makefile~~
* gracefully close the memcached connection
