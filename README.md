
##### 版权声明:
> `github.com/googx/service` forked自 `github.com/kardianos/service`

    github.com/googx/service 项目forked自 github.com/kardianos/service 项目,该次提交之前的代码版权(包括不限于著作权)归为原作者所有,
    本次提交(以git提交hash:4df36c9fc1c6ac86231851ad6fa5627e184c94e5为分界点)之后的修改的代码版权(包括不限于著作权)为googx作者所有.
    特此声明!

service will install / un-install, start / stop, and run a program as a service (daemon).
Currently supports Windows XP+, Linux/(systemd | Upstart | SysV), and OSX/Launchd.

Windows controls services by setting up callbacks that is non-trivial. This
is very different then other systems. This package provides the same API
despite the substantial differences.
It also can be used to detect how a program is called, from an interactive
terminal or from a service manager.

## BUGS
 * Dependencies field is not implemented for Linux systems and Launchd.
 * OS X when running as a UserService Interactive will not be accurate.
