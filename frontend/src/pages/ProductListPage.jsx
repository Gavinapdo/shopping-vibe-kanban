import PageTitle from "../components/common/PageTitle";
import ProductTable from "../components/common/ProductTable";

function ProductListPage() {
  return (
    <div>
      <PageTitle title="商品列表" description="此处展示内置的静态商品数据。" />
      <section className="panel">
        <ProductTable />
      </section>
    </div>
  );
}

export default ProductListPage;
