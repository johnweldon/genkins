# genkins

Toy tool to list current status of jobs on a jenkins server

Right now only has two commands `all` and `bad`, which show status of all jobs and only bad jobs

Default jenkins server is the server for jenkins itself, you can substitute your own server by passing the path to the json api in the `--url`, `-u`, or `$GENKINS_URL` envar

genkins help

    NAME:
    genkins - get jenkins updates

    USAGE:
    genkins [global options] command [command options] [arguments...]

    VERSION:
    0.0.0

    COMMANDS:
    all, a	show status of all jobs
    bad, b	show bad jobs
    help, h	Shows a list of commands or help for one command
    
    GLOBAL OPTIONS:
    --help, -h		show help
    --version, -v	print the version
    

genkins help all

    NAME:
    all - show status of all jobs

    USAGE:
    command all [command options] [arguments...]

    OPTIONS:
    --url, -u 'https://ci.jenkins-ci.org/api/json'	url for the jenkins server json api [$GENKINS_URL]
    --full, -f						Show full build info [$GENKINS_FULL]
    

genkins help bad

    NAME:
    bad - show bad jobs

    USAGE:
    command bad [command options] [arguments...]

    OPTIONS:
    --url, -u 'https://ci.jenkins-ci.org/api/json'	url for the jenkins server json api [$GENKINS_URL]
    --full, -f						Show full build info [$GENKINS_FULL]
    
