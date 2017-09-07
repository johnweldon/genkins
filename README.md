# genkins

Toy tool to list current status of jobs on a jenkins server

Right now only has two commands `all` and `bad`, which show status of all jobs and only bad jobs

Default jenkins server is the server for jenkins itself, you can substitute your own server by passing the path to the json api in the `--url`, `-u`, or `$GENKINS_URL` envar


`genkins help`

    NAME:
       genkins - get jenkins updates

    USAGE:
       genkins [global options] command [command options] [arguments...]

    VERSION:
       0.0.0

    COMMANDS:
         all, a   show status of all jobs
         bad, b   show bad jobs
         help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version



`genkins help all`

    NAME:
       genkins all - show status of all jobs

    USAGE:
       genkins all [command options] [arguments...]

    OPTIONS:
       --url value, -u value    url for the jenkins server json api (default: "https://ci.jenkins.io/api/json") [$GENKINS_URL]
       --full, -f               Show full build info [$GENKINS_FULL]
       --id value, -i value     user id for the jenkins server json api [$GENKINS_ID]
       --token value, -t value  user token for the jenkins server json api [$GENKINS_TOKEN]



`genkins help bad`

    NAME:
       genkins bad - show bad jobs

    USAGE:
       genkins bad [command options] [arguments...]

    OPTIONS:
       --url value, -u value    url for the jenkins server json api (default: "https://ci.jenkins.io/api/json") [$GENKINS_URL]
       --full, -f               Show full build info [$GENKINS_FULL]
       --id value, -i value     user id for the jenkins server json api [$GENKINS_ID]
       --token value, -t value  user token for the jenkins server json api [$GENKINS_TOKEN]

