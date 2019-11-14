新建项目后使用以下命令初始化go mod
go mod init [项目名]

下载库的命令
go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]



旧的项目迁移到go mod:
go mod init [项目名]
go build ./...

如果存在就的vendor目录以及glide.ymal 都可以进行删除
