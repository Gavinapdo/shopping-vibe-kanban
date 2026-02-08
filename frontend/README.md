# frontend

商品管理系统前端，基于 React + Vite 实现。

## 启动方式

```bash
npm install
npm run dev
```

默认开发端口为 `5173`，并通过 Vite 代理将 `/api` 请求转发至 `http://localhost:8080`。

## 可选环境变量

- `VITE_API_BASE_URL`：后端基础地址，默认为空字符串。
  - 例如：`VITE_API_BASE_URL=http://localhost:8080`

## 页面功能

- 商品列表展示
- 新增商品
- 编辑商品
- 删除商品
- 手动刷新列表
