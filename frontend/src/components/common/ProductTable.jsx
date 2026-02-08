const mockProducts = [
  { id: 1, name: "笔记本电脑", price: 5999, stock: 12 },
  { id: 2, name: "机械键盘", price: 499, stock: 37 },
  { id: 3, name: "无线鼠标", price: 199, stock: 48 },
];

function ProductTable() {
  return (
    <table className="product-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>商品名称</th>
          <th>价格（元）</th>
          <th>库存</th>
        </tr>
      </thead>
      <tbody>
        {mockProducts.map((product) => (
          <tr key={product.id}>
            <td>{product.id}</td>
            <td>{product.name}</td>
            <td>{product.price}</td>
            <td>{product.stock}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

export default ProductTable;
