# Kiwi Clients

Kiwi 客户端，目前仅支持SSH的本地拉起。

## 安装

### Windows

下载Win安装包，双击KIWI-install.msi文件（需要等待10几秒钟）。

## 卸载

### Windows

控制面板直接卸载即可。


## 测试
1. 编写json
```json
{
	"protocol": "ssh",
	"args": {
		"ip": "10.1.1.32",
		"port": 22,
		"username": "test",
		"password": "pwd"
	}
}
```
2. 编码成base64:`ewoJInByb3RvY29sIjogInNzaCIsCgkiYXJncyI6IHsKCQkiaXAiOiAiMTAuMS4xLjMyIiwKCQkicG9ydCI6IDIyLAoJCSJ1c2VybmFtZSI6ICJ0ZXN0IiwKCQkicGFzc3dvcmQiOiAicHdkIgoJfQp9`

3. Chrome浏览器打开`kiwi://ewoJInByb3RvY29sIjogInNzaCIsCgkiYXJncyI6IHsKCQkiaXAiOiAiMTAuMS4xLjMyIiwKCQkicG9ydCI6IDIyLAoJCSJ1c2VybmFtZSI6ICJ0ZXN0IiwKCQkicGFzc3dvcmQiOiAicHdkIgoJfQp9`

## 打包


## 参考

- [浏览器拉起app原理](https://juejin.cn/post/6844903989155217421)
