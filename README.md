# showstart

## 运行说明  

命令格式  
./showstart  [参数] [模式]  
example
```bash
./showstart -config=./config/config.yaml -keyword='元气' getActivityList
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