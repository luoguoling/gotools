# gotools
record.go主要用于记录日志
需要配合
export PROMPT_COMMAND='{ $(history 1 | { read x cmd;date=$(date "+%Y-%m-%d %T");/usr/bin/record "$date ### IP:$SSH_CLIENT ### PS:$SSH_TTY ### USER:$USER ### $cmd"; });} >& /dev/null'

