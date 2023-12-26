# Xlog 项目

这是一个基于 Gin 框架的简单博客项目。

## 特性

- 用户可以浏览博客文章
- 用户可以通过搜索功能查找文章
- 基于JWT实现登录认证功能
- 超级用户可以进行后台管理

## 技术栈

- [Gin](https://github.com/gin-gonic/gin): Go 语言的 Web 框架
- [Gorm](https://gorm.io/): Go 语言的 ORM 库
- [jwt-go](https://github.com/dgrijalva/jwt-go): Go语言的jwt库

## 快速开始

1. 克隆项目：

    ```bash
    git clone https://github.com/Illusionboy/blog.git
    ```

2. 进入项目目录：

    ```bash
    cd blog
    ```

3. 安装依赖：

    ```bash
    go mod tidy
    ```

4. 建立mysql数据库
5. 建立配置文件config.yaml
    例:
    ```yaml
    # 启动端口号
    server:
    post: 8001

    # MySQL数据源配置
    mysql:
    username: root          # 数据库用户名
    password: 12345678  # 密码
    url: tcp(localhost:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local
    ```

6. 找到合适的前端模板，主要文件如下：

   - templates/
   - ├── backend/
   - ├───── index.html
   - ├───── login.html
   - ├───── layouts/
   - ├──────── master.html
   - ├───── channel/
   - ├──────── view.html
   - ├──────── list.html
   - ├───── post/
   - ├──────── view.html
   - ├──────── list.html
   - ├── frontend/
   - ├───── index.html
   - ├───── login.html
   - ├───── layouts/
   - ├──────── master.html
   - ├───── channel/
   - ├──────── list.html
   - ├───── post/
   - ├──────── detail.html
   
   同时创建static文件夹
   - static/
   - ├── images/
   - ├── thumbnails/
7. 运行项目：

    ```bash
    go run main.go
    ```

5. 访问 [8001端口](http://localhost:8001) [^1]查看项目。

## 配置

根据你使用的数据库修改 `config.yaml` 目录下可以找到项目的配置文件，你可以根据需要进行修改。

## 示例

[Xlog(https)](https://www.xlw-xlog.xyz/)

[Xlog(http)](http://www.xlw-xlog.xyz/)

[^1]: 8001端口为上面config.yaml中给出
