import { useEffect, useState } from "react";
import { listProducts } from "../../api/products";

// 商品列表表格组件，从后端 API 获取数据
function ProductTable() {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await listProducts();
        setProducts(data);
      } catch (err) {
        setError(err.message || "加载商品列表失败");
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);

  if (loading) return <p>正在加载商品列表...</p>;
  if (error) return <p className="status error">{error}</p>;

  return (
    <table className="product-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>商品名称</th>
          <th>价格（元）</th>
          <th>库存</th>
          <th>描述</th>
        </tr>
      </thead>
      <tbody>
        {products.length === 0 ? (
          <tr><td colSpan={5} style={{ textAlign: "center" }}>暂无商品数据</td></tr>
        ) : (
          products.map((product) => (
            <tr key={product.id}>
              <td>{product.id}</td>
              <td>{product.name}</td>
              <td>{product.price}</td>
              <td>{product.stock}</td>
              <td>{product.description}</td>
            </tr>
          ))
        )}
      </tbody>
    </table>
  );
}

export default ProductTable;
