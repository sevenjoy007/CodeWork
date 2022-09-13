---
description: 介绍下gitbook的使用
---

# GitBook

## 简介

GitBook.com 是一个使用工具链来创建和托管书籍的在线平台。

这个工具链 (GitBook) 是一个使用 Git 和 Markdown 来构建书籍的工具。它可以将你的书输出很多格式：PDF，ePub，mobi，或者输出为静态网页。

[gitbook官网](https://www.gitbook.com/)

[gitbook文档（中文版）](https://chrisniael.gitbooks.io/gitbook-documentation/content/)

[gitbook文档（原版）](https://docs.gitbook.com/integrations/git-sync/enabling-github-sync)——较新

## 安装

#### 常规方法

1. 安装[node.js](https://nodejs.org/en/)，默认安装了npm(node包管理工具）
2.  安装gitbook&#x20;

    <pre class="language-rust"><code class="lang-rust"><strong>npm install -g gitbook-cli</strong></code></pre>

PS：这种方法我在MAC试了很多次都有权限问题，失败。

#### 本人安装成功的方法

> 借鉴：[https://www.jianshu.com/p/ddd1a5edb456](https://www.jianshu.com/p/ddd1a5edb456)

1. 先安装nvm。并且安装提示将东西加到 \~/.zshrc

```bash
brew install nvm
vim ~/.zshrc
source ~/.zshrc
```

2\. 下载node.js**（版本10是可以成功的，10以上好像会报错，反正最新的16版本是会报错的）**

```bash
nvm ls
nvm current
nvm install 10
nvm use 10
```

3\. 安装gitbook

```bash
npm install -g gitbook-cli # -g全局，安装命令行版gitbook-cli
// npm install -g gitbook-cli@2.3.2 --save-dev //安装指定版本的命令行版gitbook-cli
gitbook -V # 查看版本号，看是否安装成功

gitbook fetch 2.6.9 # 再安装2.6.9，用该版本build出来的书籍点击目录可以跳转
gitbook ls    # 查看安装了哪些版本
```

4\. 安装插件

```bash
npm i gitbook-plugin-summary --save 
```

## 使用

### 入门

1. 初始化

```bash
gitbook init    # 创建文件夹
```

会生成`README.md` 和`SUMMARY.md` 两个文件

2\. 编辑目录文件

```markdown
# Summary

* [Introduction](README.md)
* [Read](Read/README1.md)
* [1. 季节](季节/ReadMe2.md)
    * [1.1 春](季节/section0.md)
    * [1.2 夏](季节/section1.md)
* [2. 城市](城市/ReadMe3.md)
    * [1.1 北京](城市/section0.md)
    * [1.2 上海](城市/section1.md)
```

3\. 执行

```bash
gitbook init
```

会创建目录，只支持2级目录

4\. 编译

```bash
 gitbook build # 编译
```

5\. 启动服务

```bash
gitbook serve # 启动服务
```

访问,用浏览器打开 http://localhost:4000/ 或 http://127.0.0.1:4000/ 查看显示书籍的效果。结束预览 ctrl+c

6\. 生成电子书

```bash
gitbook mobi ./ ./MyFirstBook.mobi
```

### 使用git更新

看[英文文档](https://docs.gitbook.com/integrations/git-sync)，需要结合一个github仓库（没有`git.gitbook.com`了，找不到该域名）

1. 先`Enable Github Sync`，选择一个仓库跟分支
2. 初始化可以选择 Github -> Github，或者Github -> Gitbook
3. 之后`Live edit`时，每个改变的MR会在指定分支生成一个commit；如果提commit到github，也会同步到Gitbook空间。

### 结合github page

**`在任意一个 Github 仓库中建立一个名为 gh-pages 的分支。只要 gh-pages 中的内容符合一个静态站点要求，就可以在如下地址中进行访问：https://Github用户名.gitbooks.io/Github 仓库。`**

也就是说，该git项目将有如下两个分支：

* master, 保存书籍的源码。
* gh-pages, 保存书籍编译后的 HTML 文件。即`gitbook build`生成的`_book`目录下的文件。

#### push到**`gh-pages 分支`**

1. `gitbook build`将内容输出到`_book`目录
2. 创建gh-pages分支，将\_book目录内容推送到该分支

```bash
git checkout --orphan source # 创建一个孤儿分支（与原分支内容不同，如果不提交东西，这个分支实际上没有创建）
git rm --cached -r .        # 删除暂存区不需要的东西
git clean -df               # 删除当前目录下没有被track过的文件和文件夹

cp -r _book/* .
git add xxx
git commit -m 'publish'
git push -u origin gh-pages
```

3\. 然后就可以访问**`https://Github用户名.gitbooks.io/Github`**了

#### **`gh-pages`**自动发布 \[WIP]

如果每次都要手动push到gh-pages分支，比较麻烦。可以用gh-pages插件实现自动发布。

1. 安装插件

```bash
npm i -D gh-pages
```

2\. 使用该插件

### 提升

插件

[https://dunwu.gitbooks.io/gitbook-notes/content/advanced/plugins.html](https://dunwu.gitbooks.io/gitbook-notes/content/advanced/plugins.html)

### 参考教程

[https://skyao.gitbooks.io/learning-gitbook/content/creation/add\_chapter.html](https://skyao.gitbooks.io/learning-gitbook/content/creation/add\_chapter.html)\
[https://tonydeng.github.io/gitbook-zh/gitbook-howtouse/howtouse/gitbookinstall.html](https://tonydeng.github.io/gitbook-zh/gitbook-howtouse/howtouse/gitbookinstall.html)

