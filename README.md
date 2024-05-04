# 秒杀系统
利用golang实现的一个简易商品秒杀系统后端项目

框架：kitex,gin
数据库：mysql,gorm
缓存：redis
分布式锁：redis
其他：Token,bcrypt

结构体：商品good{name，price，keep//库存，timeout}
​	     用户（买家）user{name，passwd，token，book//订单}
  	   商家（卖家）hoster{name，passwd，token，flag//验证商家身份，good//管理商品}

kitex-server：用户和商家的注册登录和Token验证
kitex-client：处理商品逻辑

1：活动创建：CreatKill：一个商家验证身份后创建一种商品
2：商家查看活动（订单）：ShowKill: 一个商家验证身份后查看自己的商品情况
3：用户参与活动：
​	JoinKill：用户参与秒杀活动
​	利用redis分布式锁保持秩序
​	下订单后商品库存-1
​	释放锁后，如果还未付款，则库存+1
