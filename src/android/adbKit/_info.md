## problem

### 执行 adb connect 连接局域网内的安卓时报错: failed to connect to '192.168.60.205:16384': No route to host

原因很可能是 macOS 的「本地网络」隐私权限。 新版 macOS 对 App 访问局域网设备（192.168.60.x）有更严格的控制，GoLand 启动的子进程继承
GoLand 的权限，而 GoLand 可能没有局域网访问权限。

先做个快速验证——在终端直接运行（不通过 GoLand）：

cd /Users/richelieu/GolandProjects/chimera && go run test/test.go

如果终端运行成功（前提是终端有"本地网络"的权限），就确认是 GoLand 的权限问题。

然后去 系统设置 → 隐私与安全性 → 本地网络，看看 GoLand 是否在列表里，没有的话需要授权。