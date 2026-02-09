import { useEffect, useMemo, useState } from "react";
import {
  createProduct,
  deleteProduct,
  listProducts,
  updateProduct,
} from "../api/products";
import PageTitle from "../components/common/PageTitle";

const EMPTY_FORM = {
  name: "",
  price: "",
  stock: "",
  description: "",
};

function getProductId(product) {
  return product?.id ?? product?.productId ?? "";
}

function toFormState(product) {
  return {
    name: product?.name ?? "",
    price: product?.price ?? "",
    stock: product?.stock ?? "",
    description: product?.description ?? "",
  };
}

function toPayload(formData) {
  return {
    name: formData.name.trim(),
    price: Number(formData.price),
    stock: Number(formData.stock),
    description: formData.description.trim(),
  };
}

function ProductListPage() {
  const [products, setProducts] = useState([]);
  const [formData, setFormData] = useState(EMPTY_FORM);
  const [editingId, setEditingId] = useState("");
  const [loading, setLoading] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const isEditing = useMemo(() => editingId !== "", [editingId]);

  // 获取商品列表
  const fetchProducts = async () => {
    setLoading(true);
    setErrorMessage("");
    try {
      const data = await listProducts();
      setProducts(data);
    } catch (error) {
      setErrorMessage(error.message || "加载商品列表失败");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchProducts();
  }, []);

  // 表单输入处理
  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  // 重置表单
  const resetForm = () => {
    setFormData(EMPTY_FORM);
    setEditingId("");
  };

  // 表单验证
  const validateForm = () => {
    if (!formData.name.trim()) return "商品名称不能为空";
    if (formData.price === "" || Number.isNaN(Number(formData.price))) return "商品价格必须是数字";
    if (Number(formData.price) < 0) return "商品价格不能为负数";
    if (formData.stock === "" || Number.isNaN(Number(formData.stock))) return "商品库存必须是数字";
    if (!Number.isInteger(Number(formData.stock))) return "商品库存必须是整数";
    if (Number(formData.stock) < 0) return "商品库存不能为负数";
    return "";
  };

  // 提交表单（创建或更新）
  const handleSubmit = async (event) => {
    event.preventDefault();
    const validationMessage = validateForm();
    if (validationMessage) {
      setErrorMessage(validationMessage);
      return;
    }

    setSubmitting(true);
    setErrorMessage("");
    setSuccessMessage("");

    try {
      const payload = toPayload(formData);
      if (isEditing) {
        await updateProduct(editingId, payload);
        setSuccessMessage("商品更新成功");
      } else {
        await createProduct(payload);
        setSuccessMessage("商品创建成功");
      }
      resetForm();
      await fetchProducts();
    } catch (error) {
      setErrorMessage(error.message || "保存商品失败");
    } finally {
      setSubmitting(false);
    }
  };

  // 编辑商品
  const handleEdit = (product) => {
    setFormData(toFormState(product));
    setEditingId(String(getProductId(product)));
    setErrorMessage("");
    setSuccessMessage("");
  };

  // 删除商品
  const handleDelete = async (product) => {
    const id = getProductId(product);
    if (!id) {
      setErrorMessage("无法删除：商品缺少ID");
      return;
    }
    const confirmed = window.confirm(`确定要删除商品「${product.name}」吗？`);
    if (!confirmed) return;

    setErrorMessage("");
    setSuccessMessage("");
    try {
      await deleteProduct(id);
      setSuccessMessage("商品删除成功");
      if (String(id) === editingId) resetForm();
      await fetchProducts();
    } catch (error) {
      setErrorMessage(error.message || "删除商品失败");
    }
  };

  return (
    <div>
      <PageTitle title="商品列表" description="管理所有商品，支持增删改查操作。" />

      {/* 表单区域 */}
      <section className="panel">
        <h3>{isEditing ? "编辑商品" : "新增商品"}</h3>
        <form className="form" onSubmit={handleSubmit}>
          <div className="form-row">
            <label htmlFor="name">商品名称</label>
            <input id="name" name="name" value={formData.name} onChange={handleInputChange} placeholder="请输入商品名称" required />
          </div>
          <div className="form-row">
            <label htmlFor="price">商品价格</label>
            <input id="price" name="price" type="number" min="0" step="0.01" value={formData.price} onChange={handleInputChange} placeholder="请输入商品价格" required />
          </div>
          <div className="form-row">
            <label htmlFor="stock">商品库存</label>
            <input id="stock" name="stock" type="number" min="0" step="1" value={formData.stock} onChange={handleInputChange} placeholder="请输入商品库存" required />
          </div>
          <div className="form-row">
            <label htmlFor="description">商品描述</label>
            <textarea id="description" name="description" value={formData.description} onChange={handleInputChange} placeholder="请输入商品描述" rows={3} />
          </div>
          <div className="button-row">
            <button type="submit" disabled={submitting}>{submitting ? "提交中..." : isEditing ? "保存修改" : "新增商品"}</button>
            <button type="button" className="button-secondary" onClick={resetForm}>重置</button>
          </div>
        </form>
      </section>

      {/* 消息提示 */}
      {errorMessage && <p className="status error">{errorMessage}</p>}
      {successMessage && <p className="status success">{successMessage}</p>}

      {/* 商品列表 */}
      <section className="panel">
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
          <h3>商品数据</h3>
          <button type="button" className="button-secondary" onClick={fetchProducts} disabled={loading}>
            {loading ? "加载中..." : "刷新列表"}
          </button>
        </div>
        <table className="product-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>名称</th>
              <th>价格</th>
              <th>库存</th>
              <th>描述</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {loading ? (
              <tr><td colSpan={6} style={{ textAlign: "center" }}>正在加载商品列表...</td></tr>
            ) : products.length === 0 ? (
              <tr><td colSpan={6} style={{ textAlign: "center" }}>暂无商品数据</td></tr>
            ) : (
              products.map((product) => {
                const id = getProductId(product);
                return (
                  <tr key={String(id)}>
                    <td>{id}</td>
                    <td>{product.name}</td>
                    <td>{product.price}</td>
                    <td>{product.stock}</td>
                    <td>{product.description}</td>
                    <td>
                      <button type="button" onClick={() => handleEdit(product)}>编辑</button>{" "}
                      <button type="button" className="button-danger" onClick={() => handleDelete(product)}>删除</button>
                    </td>
                  </tr>
                );
              })
            )}
          </tbody>
        </table>
      </section>
    </div>
  );
}

export default ProductListPage;
