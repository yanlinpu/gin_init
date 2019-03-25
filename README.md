开发默认统一为Mac
====

## 项目初始化
1. 命令执行
    
    - 安装包管理工具 `$ brew install dep`
    - 拉取包 `$ dep ensure -v` 
    - 或者一条make命令代替以上两条命令 `$ make proInit`

1. 项目启动

    - `$ make dev`

## 包管理（DEP）

* 初始化

```
    $ dep init
    $ make depInit
```

* 拉取已有包
   
```
    $ dep ensure -v
    $ make proInit
```

* 增加新包
   
```
    ## 每次境加新包之后记得运行 dep ensure，同步记录。
    $ dep ensure -add github.com/xxxx/xxxxx@=0.4.3
    $ make depAdd src=github.com/xxxx/xxxxx@=0.4.3
```
