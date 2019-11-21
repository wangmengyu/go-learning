安装最新版本的命令从google 搜索关键字：
Install Elasticsearch with Docker
可以找这个网址：
https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html

命令如下：
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.4.2
运行单独节点：
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.4.2



IDE 内置工具位于
tools->HttpClient->Test RestFUL web service

HOST/PORT: http://localhost:9200

1.查看ES状态
HTTP method = GET
PATH: 空， 执行。

2.新增记录
 method 选择 PUT 或者 POST
  用POST可以增加之前不存在的字段
  而PUT不可以
  
 
 path 填写 index/type/id    
  意义： 库名/表名/主键
  举例： imooc/course/1
  (id如果不给，系统会自定义一个id)
  
 切到Request标签。右侧Request body 选择Text,点击文件夹编写需要提交的JSON内容：
 例如：
 {   "name": "golang",   "instructor" : "ccmouse"  }
 
 点击运行执行

3. 查询单个记录
 method = GET
 path = 库名/表名/id
 实例 ：imooc/course/1
 
4.搜索全部记录
method = GET
path = 库名/表名/_search
(注意request body.text需要清空)
实例：imooc/course/_search 
 
5. 搜索包含关键字golang的记录
method = GET
path =  imooc/course/_search
Request标签下， Request paramters 添加两个字段：
q=golang
pretty=

6.修改记录
method = POST
path = 库名/表名/id 
举例：imooc/course/2

request.body 选择text
填写完整数据：
{   "name": "interview",   "instructor": "ccmouse",   "url": "https://coding.imooc.com/class/132.html"    }

提交执行




 
 
 
 
 
  
  




