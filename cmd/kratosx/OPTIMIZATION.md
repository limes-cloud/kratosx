# cmd/kratosx/ 优化清单

## 一、代码质量问题

| # | 文件 | 问题 | 严重度 |
|---|------|------|--------|
| 1 | internal/proto/client/client.go:268-306 | 大段注释掉的旧 `generate` 函数未清理 | 中 |
| 2 | internal/proto/add/add.go:79-87 | 注释掉的 `modName` 函数未清理 | 低 |
| 3 | internal/proto/client/client.go:1-3 | import 未按标准分组（标准库/第三方/本项目混排） | 低 |
| 4 | internal/base/mod.go:48-82 | `KratosMod` 和 `KratosxMod` 两个函数逻辑几乎完全相同（获取 GOMODCACHE/GOPATH），可抽取公共辅助函数 | 中 |
| 5 | internal/base/mod_test.go:8-13 | `TestModuleVersion` 参数签名不匹配（缺少 `proDir`），测试不可能通过编译 | 高 |
| 6 | internal/pkg/pkg.go:3-5 | `ListType` 接口定义了但 `InList` 使用的是 `comparable` 约束，`ListType` 接口本身未被使用 | 中 |

## 二、Bug / 逻辑缺陷

| # | 文件 | 问题 | 严重度 |
|---|------|------|--------|
| 7 | internal/run/run.go:135 | `findCMD` 中 `filepath.Join(base, "..")` 结果未赋值给 `base`，导致 for 循环实际上不会向上遍历目录，是一个明显的 bug | 高 |
| 8 | internal/run/run.go:71-74 | `go run` 的 `fd.Dir = dir` 设置了绝对路径作为工作目录，但紧接着 `changeWorkingDirectory` 可能再次覆盖它，逻辑语义不清 | 中 |
| 9 | internal/change/get.go:65 | GitHub API 参数名错误：`pre_page` 应为 `per_page` | 高 |
| 10 | internal/proto/client/client.go:95 | `overlapIndex` 初始化为 0 但后面判断 `if overlapIndex != -1` 永远为 true，逻辑有误 | 中 |
| 11 | internal/change/get_test.go:11 | 测试用例期望值与实际 URL 不匹配（URL 是 `limes-cloud/kratosx` 但期望 owner 是 `go-kratos`），测试数据有误 | 高 |

## 三、代码重复 / 可维护性

| # | 文件 | 问题 | 严重度 |
|---|------|------|--------|
| 12 | internal/project/new.go & internal/project/add.go | `New` 和 `Add` 方法中"目录存在确认+删除+创建+重命名+打印提示"逻辑高度重复，可抽取公共函数 | 中 |
| 13 | internal/run/run.go / internal/change/change.go | 错误输出格式 `fmt.Fprintf(os.Stderr, "\033[31mERROR:..."` 散落多处，应统一为一个 error 输出工具函数 | 低 |
| 14 | internal/base/mod.go:84-116 | `KratosxCliMod` 获取 GOMODCACHE/GOPATH 的逻辑与 `KratosMod`/`KratosxMod` 重复 | 中 |

## 四、安全 / 健壮性

| # | 文件 | 问题 | 严重度 |
|---|------|------|--------|
| 15 | internal/project/new.go:38 / internal/project/add.go:36 | `os.RemoveAll(to)` 未检查错误返回值 | 低 |
| 16 | internal/proto/client/client.go:215 | `os.ReadFile` 错误被忽略后继续使用 `protoBytes`，如果读取失败 `protoBytes` 为 nil，后续逻辑可能异常 | 低 |
| 17 | internal/project/project.go:122-125 | `processProjectParams` 中 `projectName[2:]` 切片索引假设 `~` 后面紧跟 `/`，如果用户输入 `~name` 会出错 | 中 |

## 五、过时 / 无用代码

| # | 文件 | 问题 | 严重度 |
|---|------|------|--------|
| 18 | internal/base/install_compatible.go | `go.mod` 已声明 `go 1.24.6`，`!go1.17` 的构建约束永远不会生效，该文件可以删除 | 中 |
| 19 | internal/change/change.go:25 | 默认 repoURL 指向 `limes-cloud/kratos.git`（而非 `kratosx`），changelog 获取的是 kratos 而非 kratosx 的变更日志，可能不符合预期 | 中 |
| 20 | internal/proto/client/client.go:170-192 | `findProtoFiles` 函数通过 Walk 收集目录但结果顺序不确定（map迭代），且与 `walk` 函数功能部分重叠 | 低 |

## 六、体验 / 规范优化

| # | 文件 | 问题 | 严重度 |
|---|------|------|--------|
| 21 | version.go | 版本号硬编码，建议通过 `ldflags` 注入或读取 go module 版本 | 低 |
| 22 | internal/change/change.go:15 | Long 描述中写的是 `kratos changelog` 而非 `kratosx changelog` | 低 |
| 23 | internal/change/get.go:99 | GitHub API Authorization header 格式应为 `Bearer <token>` 或 `token <token>`，当前直接传 token 可能认证失败 | 中 |

---

## 优先级建议

### 必须修复（高严重度）
- #5 — TestModuleVersion 参数签名不匹配，编译失败
- #7 — findCMD 循环不会向上遍历，功能失效
- #9 — GitHub API 参数名拼写错误 pre_page → per_page
- #11 — 测试数据与 URL 不匹配

### 建议修复（中严重度）
- #1, #4, #6, #10, #12, #14, #17, #18, #19, #23

### 可选改进（低严重度）
- #2, #3, #8, #13, #15, #16, #20, #21, #22
