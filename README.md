

AdminLTE 2.4.0

### 2017 10 07
index.tpl default.go  
Display raw Post data for testing



### 2017 10 13
开始构建m2m的Dish数据表  
rel(fk)是代表 forgein key, rel(fk)的表cloumn都是xxx\_id的形式  
m2m会自动创建一个中间表  


### 2017 10 18
更换了:
bootstrap.js => bootstrap.min.js
bootstrap.css => bootsrap.min.css
select2.css => select2.css(4.0.4)
新增了:
select2-full.js(4.0.4)
select2-bootstrap.min.css
改变了:
lib.min.js 去除了select2部分
admin.js ,适应新的select2.js的ajax请求格式,params.term方式

