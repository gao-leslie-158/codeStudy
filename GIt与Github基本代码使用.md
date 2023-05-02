# GIt与Github基本代码使用

## 1、在本地仓库打开GIt bash

## 2、步骤

- 初始化

  ```
  git init  
  ```

- 建立远程连接

  ```
  git remote add origin https:github.com/....[仓库网址，也可以用SSH]
  ```

- 将文件放入缓存区

  ```
  git add [filename]
  ```

- 提交到本地仓库

  ```
  git commit -m "提交注释"
  ```

- 创建分支

  ```
  gir branch -M 分支名
  ```

- 提交到远程仓库

  ```
  git push -u origin main
  ```


## 3、其他一些git命令

- ### 设置名字和地址

  要设置 Git 的用户名和邮箱地址，可以使用以下命令：

  ```
  git config --global user.name "Your Name"
  git config --global user.email "your_email@example.com"
  ```

  其中 `Your Name` 是你的名字，`your_email@example.com` 是你的邮箱地址。这个命令会将你的名字和邮箱地址存储在 Git 的全局配置中，以便在你提交代码时自动添加作者信息。

  如果你想要为一个特定的 Git 仓库设置不同的用户名和邮箱地址，可以在进入该 Git 仓库的目录后，使用以下命令：

  ```
  git config user.name "Your Name"
  git config user.email "your_email@example.com"
  ```

  这个命令会将你的名字和邮箱地址存储在该 Git 仓库的配置中，以便在你提交代码时自动添加作者信息。

  希望这些信息可以帮助你设置 Git 的用户名和邮箱地址。

- ### 添加SSH Key

  首先在本地创建ssh key。在刚刚新建好的文件夹内点击右键Git Bash Here进入git命令行。

  ```text
  ssh-keygen -t rsa -C "your_email@163.com"
  ```

  “your_email@163.com”改成自己注册github时的邮箱，此处不一定要用163邮箱。

  回车之后会要求确认路径和输入密码，直接一路回车就行。

  成功的话会在~/下生成.ssh文件夹，进去打开id_rsa.pub，复制里面的key。

  ```text
  cat ~/.ssh/id_rsa.pub
  ```

  直接在命令行内输入上面的代码，就会出现key，右键复制key。

  从ssh-rsa开始，复制好后回到Github网页，点击右上角的setting，左侧菜单切换到SSH and GPG keys，点击New SSH key。默认是没有SSH key的。

  为了验证是否成功，在git bash下输入：ssh -T git@github.com如果是第一次的会提示是否continue，输入yes就会看到：

  You’ve successfully authenticated, but GitHub does not provide shell access。

  这就表示已成功连上github了。

- ### 清空当前git缓存区

  ```
  git rm --cached -r .
  ```


- ### 将远程分支的提交记录合并到本地分支：

  ```
  git pull origin main
  ```
