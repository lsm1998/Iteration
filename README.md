# Iteration

娱乐项目，版本迭代工具，自动发包 \
由于项目开发阶段频繁发包，每次手动连接服务器比较麻烦，所以使用golang完成一个简单的自动化发包工具 \
原理就是通过ftp将jar包传输到服务器，再调用执行shell脚本，编译并运行docker容器

处理半包和粘包
1. https://www.jb51.net/article/165021.htm
2. https://segmentfault.com/a/1190000013493942

使用方式：
1. 服务端使用build.bat编译，发在jar包和dockerfile同级目录，服务端json配置只需要port即可；
2. 客户端需要指定json配置各项，输入pag命令即是发包指令；
3. 支持运行命令，不用连接服务器即可使用基本命令；

文件说明：
1. run.sh 运行服务端程序的脚本；
2. build.sh 结束旧版本程序，并编译运行新版本程序的脚本；
3. config.json 配置文件，服务端只需要port项；