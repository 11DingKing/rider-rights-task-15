# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

调用审计查询时把 page_size 传成一万，服务就一次读出一万条；offset 传负数时也没有统一处理。请给分页参数补上默认值、上限和负偏移保护，避免一次请求拖垮服务。测试文件不要新增或改写。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-15
- 仓库地址：https://github.com/11DingKing/rider-rights-task-15.git
- parent SHA：6ba7e5b8188ec37d8df78bbe8723b96939737b48

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-15.git bug-repro
cd bug-repro
git checkout --detach 6ba7e5b8188ec37d8df78bbe8723b96939737b48
go test ./internal/domain -run "^TestAuditPageSizeIsCapped$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestAuditPageSizeIsCapped$" -count=1
--- FAIL: TestAuditPageSizeIsCapped (0.00s)
    task15_test.go:8: page size was not capped: 10000
FAIL
FAIL	riderguard/internal/domain	0.053s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestAuditPageSizeIsCapped$" -count=1
--- FAIL: TestAuditPageSizeIsCapped (0.00s)
    task15_test.go:8: page size was not capped: 10000
FAIL
FAIL	riderguard/internal/domain	0.003s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，审计查询未传页大小时使用默认值，超过最大值时限制到 200，负数偏移归零；正常小页查询和总数统计保持不变。定向测试、相关包测试及全量回归必须通过，不得删除、跳过或削弱测试。
