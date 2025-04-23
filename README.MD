# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/quick

# gf命令使用
## gf gen ctrl：
1. 编写/api/模块/版本/定义文件.go，并定义请求响应结构体: CreateReq、CreateRes
2. gf gen ctrl: 自动生成Controller层文件
3. Controller层职责
    - 参数校验
    - 实现接口逻辑：在Controller层直接实现接口逻辑，或者调用Service层接口/第三方接口实现接口逻辑
    - 生成返回数据

## gf gen dao：
1. internal/dao
    - 数据操作对象：通过对象方式访问底层数据源，底层基于ORM实现。
    - 文件可根据需要扩展。
2. internal/model/do
    - 数据转换模型：业务模型到数据模型的转换。
3. internal/model/entity
    - 数据模型：由工具维护，用户不能修改。