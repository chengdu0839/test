
app.conf配置文件：

开启session
sessionon = true

session存储位置。默认存储在内存中，如果存储在内存中，重启工程时，session则失效
SessionProvider = file

使用文件存储session时的位置
SessionSavePath = ./tmp