### 简介

我喜欢定义一些脚本到 `package.json` 里面， 因为 `npm run` 运行的时候能确定运行目录
是 `pacakge.json`  所在的文件夹，这样的话脚本就不需要判断自己的运行路径是否正确了。
但是为了这个好处去装 `nodejs` 感觉有点太浪费了，所以想写一个兼容 `npm run` 的替代程序

是基于这个仓库改的 `https://github.com/legodude17/npm-exec.git`
