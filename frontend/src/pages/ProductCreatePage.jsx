import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { createProduct } from "../api/products";
import PageTitle from "../components/common/PageTitle";

const EMPTY_FORM = {
  name: "",
  price: "",
  stock: "",
  description: "",
};

function ProductCreatePage() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState(EMPTY_FORM);
  const [submitting, setSubmitting] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  // 表单输入处理
  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
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

  // 提交表单
  const handleSubmit = async (event) => {
    event.preventDefault();
    const validationMessage = validateForm();
    if (validationMessage) {
      setErrorMessage(validationMessage);
      return;
    }

    setSubmitting(true);
    setErrorMessage("");

    try {
      await createProduct({
        name: formData.name.trim(),
        price: Number(formData.price),
        stock: Number(formData.stock),
        description: formData.description.trim(),
      });
      // 创建成功后跳转到商品列表
      navigate("/products");
    } catch (error) {
      setErrorMessage(error.message || "创建商品失败");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <div>
      <PageTitle title="新增商品" description="填写商品信息并提交到后端。" />
      <section className="panel">
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
          {errorMessage && <p className="status error">{errorMessage}</p>}
          <div className="button-row">
            <button type="submit" disabled={submitting}>{submitting ? "提交中..." : "创建商品"}</button>
            <button type="button" className="button-secondary" onClick={() => navigate("/products")}>返回列表</button>
          </div>
        </form>
      </section>
    </div>
  );
}

export default ProductCreatePage;
