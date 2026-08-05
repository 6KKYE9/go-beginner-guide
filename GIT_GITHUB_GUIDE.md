# 🔗 Git 与 GitHub 完整连接指南（新手版）

> 本指南带你从“什么都没装”到“代码成功上传 GitHub”。
> 每步都带有 **【检查点】**，做完请对照验证，确保不出错再继续。

---

## 第 1 步：环境检查 —— 确认是否安装了 Git

打开 PowerShell（按 Win 键，输入 `powershell` 回车），输入：

```powershell
git --version
```

- ✅ **已安装**：显示类似 `git version 2.45.0.windows.1`
- ❌ **未安装**：提示“无法识别”之类的错误（就像本次环境一样）

### 如果没装 Git，按下面安装：

1. 打开官网下载页：https://git-scm.com/download/win
2. 下载 **64-bit Git for Windows Setup**
3. 双击安装，**一路点 Next**（默认选项即可，会自动加入 PATH）
4. 安装完成后，**重新打开** PowerShell，再执行 `git --version` 验证

> ⚠️ 关键：安装完必须**重新打开**命令行窗口，否则系统还找不到 git。

**【检查点 1】** 执行 `git --version` 能显示版本号 ✅

---

## 第 2 步：配置 Git 用户信息（用户名 + 邮箱）

Git 每次提交代码都要记录“是谁提交的”，所以需要先告诉它你的身份。
这里的用户名/邮箱**不要求**和 GitHub 完全一致，但建议用同一个，方便辨认。

```powershell
# 设置全局用户名（把 "你的名字" 换成你喜欢的，比如 ASUS）
git config --global user.name "你的名字"

# 设置全局邮箱（换成你的邮箱）
git config --global user.email "你的邮箱@example.com"
```

### 检查当前配置状态：

```powershell
git config --global --list
```

应能看到刚才设置的 `user.name` 和 `user.email`。

**【检查点 2】** `git config --global --list` 中能看到正确的 user.name 和 user.email ✅

> 💡 小知识：`--global` 表示“对这台电脑所有仓库生效”。如果只想对单个项目设置，进到项目目录后去掉 `--global` 即可。

---

## 第 3 步：生成 SSH 密钥并添加到 GitHub

**为什么需要 SSH 密钥？**
简单说，它是一种“身份证”，让你电脑和 GitHub 之间安全地通信，免去了每次推送都输密码的麻烦。

### 3.1 生成密钥

```powershell
ssh-keygen -t ed25519 -C "你的邮箱@example.com"
```

执行后会出现提示：
```
Generating public/private ed25519 key pair.
Enter file in which to save the key (C:\Users\ASUS/.ssh/id_ed25519):
```
直接**按回车**（使用默认路径）。接着会问是否设置密码，也直接**按回车两次**（不设密码，最省事）。

### 3.2 确认密钥文件位置

默认会生成两个文件：
- `C:\Users\ASUS\.ssh\id_ed25519` —— 私钥（**绝不要给别人看**）
- `C:\Users\ASUS\.ssh\id_ed25519.pub` —— 公钥（要复制到 GitHub 的就是它）

验证文件存在：
```powershell
ls C:\Users\ASUS\.ssh\
```
应看到 `id_ed25519` 和 `id_ed25519.pub` 两个文件。

### 3.3 复制公钥内容

```powershell
# 这条命令会把公钥内容直接输出，全选复制
cat C:\Users\ASUS\.ssh\id_ed25519.pub
```
把输出的整段文字（以 `ssh-ed25519` 开头，以你的邮箱结尾）复制下来。

### 3.4 把公钥添加到 GitHub

1. 打开 https://github.com 并登录
2. 点击右上角头像 → **Settings**（设置）
3. 左侧找到 **SSH and GPG keys** → 点击 **New SSH key**
4. Title 随便写（如 `我的笔记本`），Key 类型选 `Authentication Key`
5. 把刚才复制的公钥粘贴到 Key 框里，点 **Add SSH key**

### 3.5 验证连接

```powershell
ssh -T git@github.com
```
首次会问 `Are you sure...`，输入 `yes` 回车。
成功会显示：`Hi 你的用户名! You've successfully authenticated...`

**【检查点 3】** `ssh -T git@github.com` 显示成功认证 ✅

---

## 第 4 步：创建 GitHub 仓库

1. 打开 https://github.com → 右上角 **+** → **New repository**
2. Repository name 填：`go-beginner-guide`
3. 选择 **Public**（公开，方便学习分享）
4. **不要**勾选 “Add a README”（我们本地已有）
5. 点击 **Create repository**
6. 创建后会看到一个页面，先**不要急着照它给的命令全跑**，记下仓库的 SSH 地址，形如：
   `git@github.com:你的用户名/go-beginner-guide.git`

---

## 第 5 步：把本地项目连接到远程仓库并推送

在**项目目录** `go-beginner-guide` 里打开 PowerShell，依次执行：

```powershell
# 1) 初始化本地 Git 仓库（会在目录下生成 .git 文件夹）
git init

# 2) 把远程仓库地址记下来，名字叫 origin（惯例叫法）
git remote add origin git@github.com:你的用户名/go-beginner-guide.git

# 3) 确认远程地址添加成功
git remote -v

# 4) 把当前所有文件加入“暂存区”（准备提交）
git add .

# 5) 提交，并写一段说明
git commit -m "初次提交：Go 初学者学习项目"

# 6) 推送到 GitHub（第一次推送需要 -u 关联分支）
git push -u origin main
```

> ⚠️ 注意：新版 Git 默认分支名是 `main`。如果你的本地分支叫 `master`，把上面 `main` 换成 `master`，或先执行 `git branch -M main` 改名。

**【检查点 4】** 打开 GitHub 仓库页面，能看到你上传的代码文件 ✅

---

## 第 6 步：日常推送（push）与拉取（pull）

以后每次改完代码，固定三步走：

```powershell
git add .                          # 1. 添加改动
git commit -m "描述你改了什么"      # 2. 提交
git push                           # 3. 推送到 GitHub（已关联后只需 git push）
```

**从 GitHub 拉取别人/其他设备的改动：**
```powershell
git pull
```

---

## 常见问题与解决方法

| 问题 | 原因 | 解决方法 |
|------|------|----------|
| `git: command not found` | Git 未安装或未加入 PATH | 重装 Git，并重新打开命令行 |
| `Permission denied (publickey)` | SSH 密钥没配对 | 回到第 3 步，确认公钥已添加到 GitHub，且用 SSH 地址（非 HTTPS） |
| `failed to push some refs` | 远程有本地没有的提交（如 GitHub 自动建了 README） | 先 `git pull --rebase origin main` 再 `git push` |
| `everything up-to-date` | 还没 `git add` 和 `git commit` | 确认已执行 add 和 commit 再 push |
| 推送要反复输密码 | 用了 HTTPS 地址 | 改用 SSH 地址：`git remote set-url origin git@github.com:用户名/仓库.git` |

---

## 关键概念速记（用人话解释）

- **仓库（Repository）**：装你项目代码的“文件夹”，在 GitHub 上就是一个项目页。
- **提交（Commit）**：给当前代码拍一张“快照”，并写说明，方便以后找回。
- **推送（Push）**：把你本地的快照上传到 GitHub。
- **拉取（Pull）**：把 GitHub 上最新的代码下载到本地。
- **分支（Branch）**：代码的“平行世界”，默认叫 `main`，可以开新分支试用改动而不影响主线。

恭喜！你已经掌握了 Git + GitHub 的核心流程 🎉
