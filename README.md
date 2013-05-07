runtil
======

A tool to execute a command and stop it at a given time, should it still be running.

Usage
-----

    runtil <end time> <command>

Where <end time> is a time formatted using 24 hours format (1PM = 13:00).
If <end time> is anterior to the current time, runtil will allow the command to run until the asked time *the next day*.

The command is cleanly stopped by a TERM signal.

History
-------

While doing the initial backup over a FUSE file-system, I needed to stop the *rsync* early enough to allow the various caches to flush and *umount* operations to finish before I could shutdown my computer.
I first thought about using the *at* command, but discovered that:

* it was not installed by default on my distribution
* it needs a daemon to run

As I did not want to add yet another daemon on the system and was looking for something to code in *GO*…
