# showstart

## 运行说明  

命令格式  
./showstart  [配置文件参数] [模式]  
example
```bash
./showstart -config=./config/config.yaml getActivityList
```
### 参数

1、 配置文件地址

- config=./config/config.yaml

### 运行模式

1、查看搜索演出活动

mode:getActivityList

- keyword [搜索条件]
- pagno [页数]


2、查看指定演出信息  

mode：getActivityDetail

- activityId<演出id> 从查看搜索演出活动中查查询,此参数必传


3、查看指定演出票价  

mode: getTicketList  

- activityId<演出Id> 从查看搜索演出活动中查查询,此参数必传

4、获取观演人列表  

mode: getSpectatorList

5、 下订单  

mode: order

- activityId<演出id> 从查看搜索演出活动中查查询,此参数必传

- ticketId<票价id> 从查看指定演出票价获取

- commonPerformerID[观演人ID] 从观演人列表获取

5、查询剩余是否有票数

mode: getRemainTicket