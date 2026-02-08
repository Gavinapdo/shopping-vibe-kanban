# 商品管理 API 文档

本文档描述当前后端提供的全部 HTTP 接口，服务基础地址示例：`http://localhost:8080`。

## 通用约定

- 接口前缀：`/api/v1`
- 数据格式：请求与响应均为 `application/json`
- 错误响应格式：

```json
{
  "error": "错误信息"
}
```

## 1. 健康检查

- 方法：`GET`
- 路径：`/healthz`
- 用途：探活检查

成功响应（`200`）：

```json
{
  "status": "ok"
}
```

---

## 2. 查询商品列表

- 方法：`GET`
- 路径：`/api/v1/products`

成功响应（`200`）：

```json
{
  "items": [
    {
      "id": 1,
      "name": "无线鼠标",
      "description": "静音按键，支持蓝牙与2.4G",
      "price": 89,
      "stock": 120
    }
  ]
}
```

---

## 3. 查询单个商品

- 方法：`GET`
- 路径：`/api/v1/products/{id}`
- 路径参数：
  - `id`：商品 ID，正整数

成功响应（`200`）：

```json
{
  "id": 1,
  "name": "无线鼠标",
  "description": "静音按键，支持蓝牙与2.4G",
  "price": 89,
  "stock": 120
}
```

失败响应：

- `400`（ID 非法）

```json
{
  "error": "商品ID不合法"
}
```

- `404`（商品不存在）

```json
{
  "error": "商品不存在"
}
```

---

## 4. 创建商品

- 方法：`POST`
- 路径：`/api/v1/products`

请求体：

```json
{
  "name": "扩展坞",
  "description": "Type-C 扩展坞",
  "price": 199,
  "stock": 30
}
```

字段说明：

- `name`：必填，非空字符串
- `description`：可选，字符串
- `price`：必填，大于 0
- `stock`：必填，大于等于 0

成功响应（`201`）：

```json
{
  "id": 4,
  "name": "扩展坞",
  "description": "Type-C 扩展坞",
  "price": 199,
  "stock": 30
}
```

失败响应：

- `400`（JSON 格式错误）

```json
{
  "error": "请求参数格式错误"
}
```

- `400`（业务参数不合法）

```json
{
  "error": "商品参数不合法"
}
```

---

## 5. 更新商品

- 方法：`PUT`
- 路径：`/api/v1/products/{id}`
- 路径参数：
  - `id`：商品 ID，正整数

请求体：

```json
{
  "name": "机械键盘Pro",
  "description": "87键，RGB背光",
  "price": 399,
  "stock": 80
}
```

成功响应（`200`）：

```json
{
  "id": 2,
  "name": "机械键盘Pro",
  "description": "87键，RGB背光",
  "price": 399,
  "stock": 80
}
```

失败响应：

- `400`（ID 非法）

```json
{
  "error": "商品ID不合法"
}
```

- `400`（JSON 格式错误）

```json
{
  "error": "请求参数格式错误"
}
```

- `400`（业务参数不合法）

```json
{
  "error": "商品参数不合法"
}
```

- `404`（商品不存在）

```json
{
  "error": "商品不存在"
}
```

---

## 6. 删除商品

- 方法：`DELETE`
- 路径：`/api/v1/products/{id}`
- 路径参数：
  - `id`：商品 ID，正整数

成功响应：`204 No Content`

失败响应：

- `400`（ID 非法）

```json
{
  "error": "商品ID不合法"
}
```

- `404`（商品不存在）

```json
{
  "error": "商品不存在"
}
```
