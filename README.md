1.开发环境设置
    GOPATH
        |-bin
        |-pkg
        |-src
            |-starter-kit
                |-api
                |-cmd
                |-configs
                |-internal
                |-pkg
                |-scripts
                |-vendor

    也就是说，把元代码放到GOPATH的src目录下即可

2.linux上编译运行
    1.  starter-kit/scripts目录下有两个sql脚本，分别对应账号库和游戏库，先创建数据库。
    2.  在starter-kit目录下，执行 make,即可生成可执行文件，可执行文件生成在starter-kit/bin目录下
    3.  cd scripts, 然后执行./run.sh，可以启动服务器
    4.  ./kill.sh停止服务器

3. Makefile提供的命令
   1. make setup。 这个命令会下下载dep包和用来做代码静态检查的 gometalinter，项目的第三方工具可以放在这个命令下面；
   2. make dep. 这个命令会下载项目依赖的第三方包。
   3. make fmt. 这个命令会格式化项目代码。
   4. make lint. 这个会对代码做静态检查
   5. make build. 会编译出可执行文件，生成的可执行文件放在当前目录的bin目录下，也可以使用make，省略build一样的效果
   6. make clean. 删除编译的临时文件，一般情况下用不到。



