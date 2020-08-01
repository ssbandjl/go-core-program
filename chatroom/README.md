# 简介



# 调试

* 步骤

  ```
  打开Reids服务器：
  D:\Redis>redis-server.exe
  
  编译：
  D:\go\src
  go build -o client.exe go_code/chatroom/client/main
  go build -o server.exe go_code/chatroom/server/main
  
  运行服务端：
  D:\go\src\backup>server1.exe
  运行客户端：
  D:\go\src\backup>client1.exe
  
  进行调试
  ```