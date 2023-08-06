#!/bin/bash

#fonts color
Green="\033[32m"
Red="\033[31m"
GreenBG="\033[42;37m"
RedBG="\033[41;37m"
Font="\033[0m"

OK="${Green}[OK]${Font}"
Error="${Red}[错误]${Font}"

cur_path=$(pwd)
COMPOSE="docker-compose"

check_docker() {
    docker --version &> /dev/null
    if [ $? -ne  0 ]; then
        echo -e "${Error} ${RedBG}未安装 Docker!${Font}"
        exit 1
    fi
    docker-compose version &> /dev/null
    if [ $? -ne  0 ]; then
        docker compose version &> /dev/null
        if [ $? -ne  0 ]; then
            echo -e "${Error} ${RedBG}未安装 Docker-compose! ${Font}"
            exit 1
        fi
        COMPOSE="docker compose"
    fi
    if [[ -n `$COMPOSE version | grep -E "\sv*1"` ]]; then
        $COMPOSE version
        echo -e "${Error} ${RedBG}Docker-compose 版本过低，请升级至v2+! ${Font}"
        exit 1
    fi
}

env_get() {
    local key=$1
    local value=`cat ${cur_path}/.env | grep "^$key=" | awk -F '=' '{print $2}'`
    echo "$value"
}

env_set() {
    local key=$1
    local val=$2
    local exist=`cat ${cur_path}/.env | grep "^$key="`
    if [ -z "$exist" ]; then
        echo "$key=$val" >> $cur_path/.env
    else
        if [[ `uname` == 'Linux' ]]; then
            sed -i "/^${key}=/c\\${key}=${val}" ${cur_path}/.env
        else
            docker run -it --rm -v ${cur_path}:/www alpine sh -c "sed -i "/^${key}=/c\\${key}=${val}" /www/.env"
        fi
        if [ $? -ne  0 ]; then
            echo -e "${Error} ${RedBG}设置env参数失败! ${Font}"
            exit 1
        fi
    fi
}

env_init() {
    if [ ! -f "${cur_path}/.env" ]; then
        cp ${cur_path}/.env.example ${cur_path}/.env
    fi
    if [ -z "$(env_get JWT_KEY)" ]; then
        env_set JWT_KEY "$(docker run -it --rm alpine sh -c "date +%s%N | md5sum | cut -c 1-32")"
    fi
    if [ -z "$(env_get DB_PASSWORD)" ]; then
        env_set DB_PASSWORD "$(docker run -it --rm alpine sh -c "date +%s%N | md5sum | cut -c 1-16")"
    fi
    if [ -z "$(env_get REDIS_PASS)" ]; then
        env_set REDIS_PASS "$(docker run -it --rm alpine sh -c "date +%s%N | md5sum | cut -c 1-16")"
    fi
    if [ -z "$(env_get APP_ID)" ]; then
        env_set APP_ID "$(docker run -it --rm alpine sh -c "date +%s%N | md5sum | cut -c 1-6")"
    fi
    if [ -z "$(env_get APP_VERSION)" ]; then
        env_set APP_VERSION "0.0.1"
    fi
}

docker_name() {
    echo `$COMPOSE ps | awk '{print $1}' | grep "\-$1\-"`
}

run_exec() {
    local container=$1
    local args=$2
    local cmd=$3
    local name=`docker_name $container`
    if [ -z "$name" ]; then
        echo -e "${Error} ${RedBG}没有找到 $container 容器! ${Font}"
        exit 1
    fi
    docker exec "$args" "$name" /bin/sh -c "$cmd"
}

check_postgres_up() {
    remaining=10
    while [ ! -d "${cur_path}/docker/postgresql/data/base" ]; do
        ((remaining=$remaining-1))
        if [ $remaining -lt 0 ]; then
            echo -e "${Error} ${RedBG}数据库安装失败! ${Font}"
            exit 1
        fi
        sleep 3
    done
}

start_server() {
    local tmp=$(env_get APP_VERSION)
    while [ -z "$version" ]; do
        read -rp "Please enter the version (Version: $tmp): " version
        [ -z "$version" ] && {
            version=$tmp
        }
    done
    env_set APP_VERSION "$version"
    echo -e "${Font}构建服务${Font}"
    run_exec "golang" "-it" "go build -o release/v$(env_get APP_VERSION)/$(env_get APP_NAME) main.go"
    cp ${cur_path}/.env ${cur_path}/release/v$(env_get APP_VERSION)/.env
    echo -e "${Font}启动服务${Font}"
    run_exec "golang" "-d" "nohup ./release/v$(env_get APP_VERSION)/$(env_get APP_NAME) 1>/dev/null"
}

check_docker

if [ $# -gt 0 ]; then
    if [[ "$1" == "init" ]] || [[ "$1" == "install" ]]; then
        shift 1
        echo -e "${Font}初始化${Font}"
        env_init
        echo -e "${Font}启动容器${Font}"
        $COMPOSE up -d
        echo -e "${Font}检测数据库${Font}"
        check_postgres_up
        start_server
        run_exec "postgres" "-it" "sh /tmp/sh/repassword.sh \"$@\""
        echo -e "${OK} 地址: ${GreenBG}http://IP:$(env_get APP_PORT)${Font}"
    elif [[ "$1" == "update" ]]; then
        shift 1
        $COMPOSE restart golang
        echo -e "${Font}拉取更新${Font}"
        git pull
        start_server
        echo -e "${OK} ${GreenBG}更新完成 ${Font}"
    elif [[ "$1" == "dev" ]]; then
        shift 1
        env_set GIN_MODE "debug"
        if [ ! -f "${cur_path}/bin/air" ]; then
            run_exec "golang" "-it" "go install github.com/cosmtrek/air@latest"
        fi
        $COMPOSE restart golang
        run_exec "golang" "-it" "air"
    elif [[ "$1" == "prod" ]]; then
        shift 1
        env_set GIN_MODE "release"
        $COMPOSE restart golang
        start_server
        echo -e "${OK} ${GreenBG}启动服务完成 ${Font}"
    elif [[ "$1" == "docs" ]]; then
        shift 1
        if [ ! -f "${cur_path}/bin/swag" ]; then
            run_exec "golang" "-it" "go install github.com/swaggo/swag/cmd/swag@latest"
        fi
        run_exec "golang" "-it" "swag i -g routers.go -dir app/$@,core/response --instanceName $@"
    elif [[ "$1" == "repassword" ]]; then
        shift 1
        run_exec "postgres" "-it" "sh /tmp/sh/repassword.sh \"$@\""
    elif [[ "$1" == "go" ]]; then
        shift 1
        e="go $*"
        run_exec "golang" "-it" "$e"
    else
        $COMPOSE "$@"
    fi
else
    $COMPOSE ps
fi