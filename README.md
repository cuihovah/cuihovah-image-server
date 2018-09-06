# CUIHOVAH-IMAGE-SERVER
写Markdown文档的时候总是头疼图片放在那里。不只是这样，Blog的时候也是会遇到同样的麻烦。索性写一个自己的图片服务，或者是文件服务用来作为附件使用。

## Usage

```shell

$ go run app.go
```

打开浏览器输入：http://hostname:10991/

- 如图所示打开上传图片
![网页效果](http://47.88.49.197:10991/static/images/e74303118530ac460e883f3217c2276a)

- 预览图片
![网页效果](http://47.88.49.197:10991/static/images/9381ef6c634168838405d8c31128a20a)

## Feature

- 上传图片功能
- 下载图片功能
- 在线预览图片

# API

### 图片上传
    - URI POST / 
    - REQUEST (multipart/form-data)
        + BODY
        
        ```
            fileUpload (file, required)
        ```

    - RESPONSE (application/json)
        + BODY
        
        ```
            {
                "code": "number",
                "msg": "string",
                "data": "string"
            }
        ```

### 图片预览
    - URI GET /static/images/:name 
    - REQUEST 
        + Params
        ```
            name (string, required) - image file name
        ```

    - RESPONSE (image/***)

## LICENSE
MIT

## TODO List

- 图片压缩
- 图片缩略图
- 图片管理
- 用户认证OAuth2.0
- 配置化服务