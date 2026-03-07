# SuAsk • 速问
![image](https://github.com/user-attachments/assets/68cbe1a7-1d9a-4487-a639-c86c15fbf6e9)

一个匿名提问箱，为了解决~~苏老师~~匿名提问箱经常跑路的问题

软件工程实训项目

---

## 它可以实现什么
![image](https://github.com/user-attachments/assets/a87d6177-c133-469e-8de1-f8d574bc923f)

- 完全匿名，用户可以畅所欲言，后台不会对用户信息进行记录
- 拥有公开论坛，用户可以自由发表提问
- 教师匿名提问箱，用户可以匿名提问教师问题，内置学院教师信息，开箱即用
- 提供提醒功能，你不会错过任何一条有关信息

---

## 如何使用

### 环境依赖
`go 1.24.1`

`goframe v2.9.3`


### 启动命令
```shell
git clone https://github.com/DreamingLri/SuAsk.git
cd SuAsk
gf run main.go
```

除此之外，您还需要配置此项目的[前端](https://github.com/DreamingLri/SuAsk-Web)

---

## 开发导读：知识介绍、代码编写与场景设计

### 1. 知识介绍

SuAsk 是一个基于 `Go + GoFrame + MySQL` 的匿名问答后端，整体结构遵循比较清晰的分层设计：

- `api/`：定义接口入参与返回值，同时声明路由、方法、标签和接口摘要
- `internal/controller/`：负责接收请求、调用 service，并将结果整理后返回前端
- `internal/logic/`：承载核心业务逻辑，例如提问、回答、通知、收藏等流程
- `internal/model/`：定义 controller 与 logic 之间交互所需的输入输出结构
- `internal/model/entity/` 与 `internal/model/do/`：数据库实体与入库对象

如果你是第一次接触这个项目，建议按下面顺序阅读文档：

1. [需求分析](doc/需求分析.md)：快速了解角色权限、匿名机制、提问/回答规则
2. [开发流程](doc/开发流程.md)：了解 `api -> model -> logic -> controller -> cmd` 的开发顺序
3. [代码规范](doc/代码规范.md)：统一命名、错误处理、结构体初始化与注释风格

### 2. 代码编写

在 SuAsk 中新增一个接口时，通常遵循下面的最小流程：

```go
// api/questions/v1/example.go
type ExampleReq struct {
	g.Meta `path:"/questions/example" method:"get" tags:"Questions" summary:"示例接口"`
	ID     int `json:"id" v:"required" dc:"问题 ID"`
}

type ExampleRes struct {
	Title string `json:"title" dc:"问题标题"`
}
```

```go
// internal/controller/questions/example.go
func (c *cQuestions) Example(ctx context.Context, req *v1.ExampleReq) (res *v1.ExampleRes, err error) {
	out, err := service.QuestionDetail().Get(ctx, model.QuestionDetailInput{ID: req.ID})
	if err != nil {
		return nil, err
	}
	return &v1.ExampleRes{Title: out.Title}, nil
}
```

编写代码时可以优先记住三点：

- 先定义 `api` 和 `model`，再补齐 `logic` 与 `controller`
- 尽量让 `controller` 只负责参数转换和结果包装，业务判断放在 `logic`
- 严格遵循已有规范，例如错误单独处理、结构体初始化使用字段名、避免无意义命名

### 3. 场景设计

这个项目最核心的设计不是“页面”，而是“匿名问答场景”本身。当前可以从三个典型场景理解系统设计：

- **场景一：学生问教师**
  - 学生发起匿名提问
  - 可选择公开或私密
  - 私密问题仅提问者与教师可见
  - 教师回复后，公开问题可进入公共问答区展示

- **场景二：学生问大家**
  - 问题默认公开
  - 所有符合权限的用户都可以查看并参与回答
  - 适合课程经验、资源分享、公共讨论等内容

- **场景三：教师持续答疑**
  - 教师可以对重复问题引用过往回答
  - 提问者可以继续追问
  - 系统通过通知提醒提问、回答与回复的状态变化

如果后续要扩展功能，可以继续围绕这些核心场景来设计，例如：

- 新的通知触发条件是否能服务于“追问-回复”闭环
- 新的权限控制是否会破坏匿名或私密提问体验
- 新的页面交互是否能匹配“问教师”和“问大家”两类入口

---

## Q&A

**Q: 为什么要创建这个项目？**
A: 学院的苏老师深受商用匿名提问箱跑路的迫害，~~据他所说已经用跑路了3个~~，便希望我们开发一个属于学院的匿名提问箱

**Q: 为什么使用go-frame？**
A: 这个项目本质上还是本专业的中级实训大作业，当时指明要求使用go语言来编写后端，而go-frame是一个较成熟的框架，便采用了

**Q: 我发现了bug怎么办？**
A: 欢迎提交[Issue](https://github.com/DreamingLri/SuAsk/issues/new)


