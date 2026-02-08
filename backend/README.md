# backend

商品管理系统后端目录遵循 `golang-standards/project-layout` 组织方式。

## 目录说明

- `cmd/server`：应用程序入口。
- `internal/app`：应用启动与装配逻辑。
- `internal/domain`：领域模型定义。
- `internal/service`：业务服务层。
- `internal/repository`：数据存储实现（当前为内存实现）。
- `internal/transport/http`：HTTP 路由与处理器。
- `configs`：配置文件。
- `docs`：API 文档与设计文档。
- `deployments/kubernetes`：Kubernetes 部署清单。
- `build/package`：镜像构建相关文件。
- `scripts`：开发与运维脚本。
- `test`：集成测试或端到端测试。
- `api`：API 协议与契约文件。
