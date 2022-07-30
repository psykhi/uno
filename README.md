Like `uniq`, but for logs.

`uno` is very useful when one wants to quickly identify the unique or new log statements in large log files. Unlike `uniq`, `uno` uses fuzzy matching
to determine if two lines are similar or not.
While it is certainly possible to write some custom awk/uniq/grep commands to do that on a specific log file,
`uno` does that _well enough_ on pretty much any log file.

_Uno was originally written while working at Unomaly, see this [blog post](https://unomaly.com/blog/its-in-the-anomalies/) for some context!_

# Install

```
go install github.com/psykhi/uno@latest
```

Or download from the [releases page](https://github.com/psykhi/uno/releases).

# Example

Take the following large `sshd` log file

```
Dec 10 11:03:00 LabSZ sshd[25417]: pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=183.62.140.253  user=root
Dec 10 11:03:02 LabSZ sshd[25417]: Failed password for root from 183.62.140.253 port 45648 ssh2
Dec 10 11:03:02 LabSZ sshd[25417]: Received disconnect from 183.62.140.253: 11: Bye Bye [preauth]
Dec 10 11:03:02 LabSZ sshd[25419]: pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=183.62.140.253  user=root
Dec 10 11:03:05 LabSZ sshd[25419]: Failed password for root from 183.62.140.253 port 46059 ssh2
Dec 10 11:03:05 LabSZ sshd[25419]: Received disconnect from 183.62.140.253: 11: Bye Bye [preauth]
Dec 10 11:03:05 LabSZ sshd[25422]: pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=183.62.140.253  user=root
...
Dec 10 11:03:17 LabSZ sshd[25430]: Received disconnect from 183.62.140.253: 11: Bye Bye [preauth]
Dec 10 11:03:17 LabSZ sshd[25432]: pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=183.62.140.253  user=root
Dec 10 11:03:19 LabSZ sshd[25432]: Failed password for root from 183.62.140.253 port 48708 ssh2
Dec 10 11:03:19 LabSZ sshd[25432]: Received disconnect from 183.62.140.253: 11: Bye Bye [preauth]
Dec 10 11:03:19 LabSZ sshd[25434]: pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=183.62.140.253  user=root
Dec 10 11:03:21 LabSZ sshd[25434]: Failed password for root from 183.62.140.253 port 49161 ssh2
Dec 10 11:03:21 LabSZ sshd[25434]: Received disconnect from 183.62.140.253: 11: Bye Bye [preauth]
Dec 10 09:18:33 LabSZ sshd[24641]: reverse mapping checking getaddrinfo for customer-187-141-143-180-sta.uninet-ide.com.mx [187.141.143.180] failed - POSSIBLE BREAK-IN ATTEMPT!
...
Dec 10 11:04:27 LabSZ sshd[25516]: Failed password for root from 183.62.140.253 port 33233 ssh2
Dec 10 11:04:27 LabSZ sshd[25516]: Received disconnect from 183.62.140.253: 11: Bye Bye [preauth]
Dec 10 11:04:27 LabSZ sshd[25513]: Failed password for invalid user admin from 103.99.0.122 port 50289 ssh2
```

Now if we apply `uno` to it
```
~ uno test_file 
Dec 10 11:03:00 LabSZ sshd[25417]: pam_unix(sshd:auth): authentication failure; logname= uid=0 euid=0 tty=ssh ruser= rhost=183.62.140.253  user=root
Dec 10 11:03:02 LabSZ sshd[25417]: Failed password for root from 183.62.140.253 port 45648 ssh2
Dec 10 11:03:02 LabSZ sshd[25417]: Received disconnect from 183.62.140.253: 11: Bye Bye [preauth]
Dec 10 09:18:33 LabSZ sshd[24641]: reverse mapping checking getaddrinfo for customer-187-141-143-180-sta.uninet-ide.com.mx [187.141.143.180] failed - POSSIBLE BREAK-IN ATTEMPT!
Dec 10 11:04:27 LabSZ sshd[25513]: Failed password for invalid user admin from 103.99.0.122 port 50289 ssh2

```

We now have a summary of all the logs that have happened, potentially highlighting "unknown unknowns". In this case,
`uno` highlights the `POSSIBLE BREAK-IN ATTEMPT!` which happened once in the file, surrounded by millions of common statements.


# Options

To see all options, use`uno -h`

### -d

The distance option `-d` can be used to specify how different a new line must be from the others we've seen to be deemed
new/unique. It can take a value between `0` and `1`. The default is `0.2` (20% difference)

### -all

To see all input lines and highlight the new ones in red, use `-all`

```bash
cat my_log_file.txt | uno -all
uno -all my_log_file.txt
```

### -p

Output log patterns (numbers are replaced by `*`)

```bash
cat my_log_file.txt | uno -p
uno -all my_log_file.txt

Jun * *:*:* combo ftpd[*]: connection from * (*-*-*-*.bflony.adelphia.net) at Fri Jun * *:*:* * 
Jun * *:*:* combo cups: cupsd shutdown succeeded
Jul  * *:*:* combo gpm[*]: *** info [mice.c(*)]: 
Jul  * *:*:* combo gpm[*]: imps2: Auto-detected intellimouse PS/*

```
