goose
=====

Maverick: I can't reach the ejection handle. Eject.  
Goose: I'm trying!  
Maverick: Eject eject eject eject!   

What's it for?
--------------

goose is an event listener for [supervisor]: http://supervisord.org/
It is intended for use with [docker] https://www.docker.com/
Using supervisor within the context of a docker container it may be necessary to kill the container if one of supervisor's sub-processes goes fatal. Goose listens for fatal events and sends SIGTERM to all processes if it detects one. After 3 seconds it will then issue SIGKILL to all processes.

WARNING!
--------

Do not send SIGTERM to goose unless you want to kill every running process you have with SIGKILL. 
