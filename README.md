### 存档

不再维护, Makefile 定义常用命令更为简单好用,  而且 make 命令在各大发行里均安装容易, 而且体积也更小, 还自带命令补全, 不过有一点不好就是需要在当前目录才能执行, npm 可以在子目录下执行

### 简介

我喜欢定义一些脚本到 `package.json` 里面， 因为 `npm run` 运行的时候能确定运行目录
是 `pacakge.json`  所在的文件夹，这样的话脚本就不需要判断自己的运行路径是否正确了。
但是为了这个好处去装 `nodejs` 感觉有点太浪费了，所以想写一个兼容 `npm run` 的替代程序

### 安装

可以通过源码安装
`go get -u github.com/shynome/snpm/cmd/snpm`
或者直接在 [release页面](https://github.com/shynome/snpm/releases) 下载可执行文件

因为没做执行环境的处理，依赖于 `sh` 命令，所以没有提供 `windows` 的可执行文件

### 参考

- `https://github.com/legodude17/npm-exec.git`
