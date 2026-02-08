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

func TestHealthz(t *testing.T) {
	handler := buildTestRouter()
	resp := performRequest(handler, http.MethodGet, "/healthz", "")
	if resp.Code != http.StatusOK {
		t.Fatalf("期望状态码为200，实际为%d", resp.Code)
	}
}

func TestListProducts(t *testing.T) {
	handler := buildTestRouter()
	resp := performRequest(handler, http.MethodGet, "/api/v1/products", "")
	if resp.Code != http.StatusOK {
		t.Fatalf("期望状态码为200，实际为%d", resp.Code)
	}

	var payload struct {
		Items []struct {
			ID int64 `json:"id"`
		} `json:"items"`
	}
	if err := json.Unmarshal(resp.Body.Bytes(), &payload); err != nil {
		t.Fatalf("响应解析失败: %v", err)
	}
	if len(payload.Items) != 3 {
		t.Fatalf("期望初始商品数量为3，实际为%d", len(payload.Items))
	}
}

func TestGetProduct(t *testing.T) {
	handler := buildTestRouter()
	resp := performRequest(handler, http.MethodGet, "/api/v1/products/1", "")
	if resp.Code != http.StatusOK {
		t.Fatalf("期望状态码为200，实际为%d", resp.Code)
	}

	var payload struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(resp.Body.Bytes(), &payload); err != nil {
		t.Fatalf("响应解析失败: %v", err)
	}
	if payload.ID != 1 || payload.Name == "" {
		t.Fatalf("返回商品数据不符合预期: %+v", payload)
	}
}

func TestCreateProduct(t *testing.T) {
	handler := buildTestRouter()
	createBody := `{"name":"扩展坞","description":"Type-C 扩展坞","price":199,"stock":30}`

	createResp := performRequest(handler, http.MethodPost, "/api/v1/products", createBody)
	if createResp.Code != http.StatusCreated {
		t.Fatalf("期望状态码为201，实际为%d", createResp.Code)
	}

	var created struct {
		ID int64 `json:"id"`
	}
	if err := json.Unmarshal(createResp.Body.Bytes(), &created); err != nil {
		t.Fatalf("响应解析失败: %v", err)
	}
	if created.ID == 0 {
		t.Fatalf("新建商品ID不应为0")
	}

	getResp := performRequest(handler, http.MethodGet, "/api/v1/products/4", "")
	if getResp.Code != http.StatusOK {
		t.Fatalf("期望新建商品可查询，实际状态码为%d", getResp.Code)
	}
}

func TestUpdateProduct(t *testing.T) {
	handler := buildTestRouter()
	updateBody := `{"name":"机械键盘Pro","description":"87键，RGB背光","price":399,"stock":80}`

	updateResp := performRequest(handler, http.MethodPut, "/api/v1/products/2", updateBody)
	if updateResp.Code != http.StatusOK {
		t.Fatalf("期望状态码为200，实际为%d", updateResp.Code)
	}

	var updated struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	if err := json.Unmarshal(updateResp.Body.Bytes(), &updated); err != nil {
		t.Fatalf("响应解析失败: %v", err)
	}
	if updated.Name != "机械键盘Pro" || updated.Price != 399 {
		t.Fatalf("更新结果不符合预期: %+v", updated)
	}
}

func TestDeleteProduct(t *testing.T) {
	handler := buildTestRouter()
	deleteResp := performRequest(handler, http.MethodDelete, "/api/v1/products/3", "")
	if deleteResp.Code != http.StatusNoContent {
		t.Fatalf("期望状态码为204，实际为%d", deleteResp.Code)
	}

	getResp := performRequest(handler, http.MethodGet, "/api/v1/products/3", "")
	if getResp.Code != http.StatusNotFound {
		t.Fatalf("期望删除后状态码为404，实际为%d", getResp.Code)
	}
}
