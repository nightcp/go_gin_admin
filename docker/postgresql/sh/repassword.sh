#!/bin/sh

new_password=$1

Green="\033[32m"
RedBG="\033[41;37m"
GreenBG="\033[42;37m"
Font="\033[0m"
OK="${Green}[OK]${Font}"
Error="${Red}[错误]${Font}"

new_encrypt=$(date +%s%N | md5sum | awk '{print $1}' | cut -c 1-8)
if [ -z "$new_password" ]; then
    new_password=$(date +%s%N | md5sum | awk '{print $1}' | cut -c 1-16)
fi

md5_password=$(echo -n `echo -n $new_password$new_encrypt | md5sum | awk '{print $1}'`$new_encrypt | md5sum | awk '{print $1}')
content=$(echo "select username from admin_users where id=1;" | psql -U $POSTGRES_USER -d $POSTGRES_DB)
account=$(echo "$content" | sed -n '3p')

if [ -z "$account" ]; then
    echo "${Error} ${RedBG}账号不存在! ${Font}"
    exit 1
fi

psql -w -U $POSTGRES_USER -d $POSTGRES_DB <<EOF
update admin_users set salt='${new_encrypt}',password='${md5_password}' where id=1;
EOF

echo "${OK} 账号: ${GreenBG}${account} ${Font}"
echo "${OK} 密码: ${GreenBG}${new_password} ${Font}"