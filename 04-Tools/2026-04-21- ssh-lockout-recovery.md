今天这个SSH登录问题，咱们从“死循环”到彻底打通，踩了不少坑，我帮你把完整的问题脉络和解决方法梳理清楚，以后再遇到就能秒定位👇
 
 
 
🔴 问题核心：SSH 登录云服务器的死循环
 
1. 症状
 
- 重置了云服务器密码，但密码登录被服务器配置禁止，始终提示  Permission denied 
- 本地自己生成的  id_ed25519  密钥登录被拒绝，提示  publickey  认证失败
- 阿里云App端账密登录也进不去，没法通过控制台直连终端
- 一度出现  Connection refused ，以为是网络/端口问题，安全组反复确认也开了22端口
 
2. 根本原因拆解
 
1. 服务器配置锁死：之前手动禁用了SSH密码登录（ PasswordAuthentication no ），且开启了  PermitRootLogin prohibit-password ，导致root用户只能用公钥登录，密码完全失效，重置密码也没用。
2. 本地密钥与服务器不匹配：你本地的  id_ed25519  公钥，没有被添加到服务器的  /root/.ssh/authorized_keys  文件里，服务器不认，所以公钥认证直接被拒绝。
3. 阿里云控制台密钥对没生效：之前绑定的阿里云密钥对，没有重启服务器，公钥没有被正确写入，导致密钥登录也失败。
4. 操作细节错误：一开始用  .pem  密钥登录时，文件没放到  ~/.ssh  目录、权限没设为  600 、命令里文件名没加引号，导致SSH找不到文件，一直报错。
 
 
 
✅ 解决方法：分步骤破局+根治
 
第一步：用阿里云官方密钥破局，先登录进服务器
 
1. 找到阿里云给你的  .pem  密钥文件，复制到本地  ~/.ssh  目录
2. 设置文件权限（SSH强制要求）： chmod 600 ~/.ssh/你的云服务器.pem 
3. 用带引号的正确命令登录： ssh -i ~/.ssh/"你的云服务器.pem" root@你的服务器IP 
4. 成功登录后，先给服务器加个“备用钥匙”
 
第二步：把本地公钥添加到服务器，实现免密登录
 
1. 本地终端查看公钥： cat ~/.ssh/id_ed25519.pub ，复制完整内容
2. 服务器终端执行，把公钥写入授权文件：bash
  
mkdir -p /root/.ssh
echo "你复制的公钥内容" >> /root/.ssh/authorized_keys
chmod 700 /root/.ssh
chmod 600 /root/.ssh/authorized_keys
 
3. 重启SSH服务： systemctl restart sshd 
4. 本地测试： ssh -i ~/.ssh/id_ed25519 root@你的服务器IP ，直接免密登录成功
 
第三步：可选，恢复密码登录（按需选择）
 
如果之后想同时用密码登录，可以修改服务器SSH配置：
 
bash
  
sed -i 's/^PermitRootLogin.*/PermitRootLogin yes/' /etc/ssh/sshd_config
sed -i 's/^PasswordAuthentication.*/PasswordAuthentication yes/' /etc/ssh/sshd_config
systemctl restart sshd
 
 
 
 
💡 关键踩坑复盘（以后避坑）
 
1.  Permission denied  别再瞎重置密码了：先看服务器  sshd_config  配置，如果是  prohibit-password ，密码永远登不上，必须用公钥。
2. 阿里云密钥对绑定后必须重启：不重启服务器，公钥不会被写入系统，绑定了也白搭。
3.  .pem  文件必须满足两个条件：权限600 + 路径正确，少一个都登不上。
4. 公钥必须是服务器信任的：本地生成的密钥，必须把公钥内容加到服务器的  authorized_keys  里，服务器才会认。
 
现在你的服务器已经双密钥都配置好了，本地密钥登录、阿里云官方密钥备用，安全又方便，再也不怕登不进去啦！