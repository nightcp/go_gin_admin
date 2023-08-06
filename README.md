# <p align="center">安装说明</p>

## 安装设置（使用Docker）
> 安装最新版 [Docker](https://www.runoob.com/docker/ubuntu-docker-install.html) 和 [Docker Compose](https://www.runoob.com/docker/docker-compose.html)

#### 1、克隆项目到您的本地或服务器
```shell
cd admin        # 进入项目目录
./cmd init      # 首次安装
```

项目地址为：**`http://ip:port`**（`port`默认为`8080`）。

#### 2、可用命令
```shell
./cmd init          # 首次安装
./cmd update        # 更新程序
./cmd dev           # 开发模式
./cmd prod          # 生成模式
./cmd repassword    # 重置admin密码
./cmd docs admin    # 生成API文档, http://ip:port/swagger/admin/index.html
./cmd go "command"  # 在golang容器中执行命令
```