package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func buildTestRouter() http.Handler {
	gin.SetMode(gin.TestMode)
	return NewRouter()
}

func performRequest(handler http.Handler, method, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)
	return resp
}

func decodeJSON(t *testing.T, resp *httptest.ResponseRecorder, out any) {
	t.Helper()
	if err := json.Unmarshal(resp.Body.Bytes(), out); err != nil {
		t.Fatalf("响应解析失败: %v", err)
	}
}

func TestHealthz(t *testing.T) {
	handler := buildTestRouter()
	resp := performRequest(handler, http.MethodGet, "/healthz", "")
	if resp.Code != http.StatusOK {
		t.Fatalf("期望状态码为200，实际为%d", resp.Code)
	}

	var payload struct {
		Status string `json:"status"`
	}
	decodeJSON(t, resp, &payload)
	if payload.Status != "ok" {
		t.Fatalf("期望status为ok，实际为%s", payload.Status)
	}
}

func TestListProducts(t *testing.T) {
	handler := buildTestRouter()
	resp := performRequest(handler, http.MethodGet, "/api/products", "")
	if resp.Code != http.StatusOK {
		t.Fatalf("期望状态码为200，实际为%d", resp.Code)
	}

	var payload struct {
		Items []struct {
			ID int64 `json:"id"`
		} `json:"items"`
	}
	decodeJSON(t, resp, &payload)

	if len(payload.Items) != 3 {
		t.Fatalf("期望初始商品数量为3，实际为%d", len(payload.Items))
	}
	if payload.Items[0].ID != 1 || payload.Items[1].ID != 2 || payload.Items[2].ID != 3 {
		t.Fatalf("期望商品ID按升序返回，实际为%+v", payload.Items)
	}
}

func TestGetProduct(t *testing.T) {
	t.Run("查询成功", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodGet, "/api/products/1", "")
		if resp.Code != http.StatusOK {
			t.Fatalf("期望状态码为200，实际为%d", resp.Code)
		}

		var payload struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		}
		decodeJSON(t, resp, &payload)
		if payload.ID != 1 || payload.Name == "" {
			t.Fatalf("返回商品数据不符合预期: %+v", payload)
		}
	})

	t.Run("商品不存在", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodGet, "/api/products/999", "")
		if resp.Code != http.StatusNotFound {
			t.Fatalf("期望状态码为404，实际为%d", resp.Code)
		}
	})

	t.Run("商品ID不合法", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodGet, "/api/products/abc", "")
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("期望状态码为400，实际为%d", resp.Code)
		}
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("创建成功", func(t *testing.T) {
		handler := buildTestRouter()
		createBody := `{"name":"扩展坞","description":"Type-C 扩展坞","price":199,"stock":30}`

		createResp := performRequest(handler, http.MethodPost, "/api/products", createBody)
		if createResp.Code != http.StatusCreated {
			t.Fatalf("期望状态码为201，实际为%d", createResp.Code)
		}

		var created struct {
			ID int64 `json:"id"`
		}
		decodeJSON(t, createResp, &created)
		if created.ID != 4 {
			t.Fatalf("期望新建商品ID为4，实际为%d", created.ID)
		}

		getResp := performRequest(handler, http.MethodGet, "/api/products/4", "")
		if getResp.Code != http.StatusOK {
			t.Fatalf("期望新建商品可查询，实际状态码为%d", getResp.Code)
		}
	})

	t.Run("请求体格式错误", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodPost, "/api/products", `{"name":`) // 非法JSON
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("期望状态码为400，实际为%d", resp.Code)
		}
	})

	t.Run("业务参数不合法", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodPost, "/api/products", `{"name":" ","description":"","price":99,"stock":10}`)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("期望状态码为400，实际为%d", resp.Code)
		}
	})
}

func TestUpdateProduct(t *testing.T) {
	t.Run("更新成功", func(t *testing.T) {
		handler := buildTestRouter()
		updateBody := `{"name":"机械键盘Pro","description":"87键，RGB背光","price":399,"stock":80}`

		resp := performRequest(handler, http.MethodPut, "/api/products/2", updateBody)
		if resp.Code != http.StatusOK {
			t.Fatalf("期望状态码为200，实际为%d", resp.Code)
		}

		var updated struct {
			ID    int64   `json:"id"`
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		}
		decodeJSON(t, resp, &updated)
		if updated.ID != 2 || updated.Name != "机械键盘Pro" || updated.Price != 399 {
			t.Fatalf("更新结果不符合预期: %+v", updated)
		}
	})

	t.Run("商品不存在", func(t *testing.T) {
		handler := buildTestRouter()
		updateBody := `{"name":"机械键盘Pro","description":"87键，RGB背光","price":399,"stock":80}`
		resp := performRequest(handler, http.MethodPut, "/api/products/999", updateBody)
		if resp.Code != http.StatusNotFound {
			t.Fatalf("期望状态码为404，实际为%d", resp.Code)
		}
	})

	t.Run("商品ID不合法", func(t *testing.T) {
		handler := buildTestRouter()
		updateBody := `{"name":"机械键盘Pro","description":"87键，RGB背光","price":399,"stock":80}`
		resp := performRequest(handler, http.MethodPut, "/api/products/0", updateBody)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("期望状态码为400，实际为%d", resp.Code)
		}
	})

	t.Run("请求体格式错误", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodPut, "/api/products/2", `{"name":`)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("期望状态码为400，实际为%d", resp.Code)
		}
	})

	t.Run("业务参数不合法", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodPut, "/api/products/2", `{"name":"键盘","description":"","price":0,"stock":1}`)
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("期望状态码为400，实际为%d", resp.Code)
		}
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("删除成功", func(t *testing.T) {
		handler := buildTestRouter()
		deleteResp := performRequest(handler, http.MethodDelete, "/api/products/3", "")
		if deleteResp.Code != http.StatusNoContent {
			t.Fatalf("期望状态码为204，实际为%d", deleteResp.Code)
		}

		getResp := performRequest(handler, http.MethodGet, "/api/products/3", "")
		if getResp.Code != http.StatusNotFound {
			t.Fatalf("期望删除后状态码为404，实际为%d", getResp.Code)
		}
	})

	t.Run("商品不存在", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodDelete, "/api/products/999", "")
		if resp.Code != http.StatusNotFound {
			t.Fatalf("期望状态码为404，实际为%d", resp.Code)
		}
	})

	t.Run("商品ID不合法", func(t *testing.T) {
		handler := buildTestRouter()
		resp := performRequest(handler, http.MethodDelete, "/api/products/-1", "")
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("期望状态码为400，实际为%d", resp.Code)
		}
	})
}
