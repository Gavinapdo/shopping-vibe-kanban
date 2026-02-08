import PageTitle from "../components/common/PageTitle";

function DashboardPage() {
  return (
    <div>
      <PageTitle
        title="欢迎使用商品管理系统"
        description="该页面用于演示前端路由和组件结构初始化。"
      />
      <section className="panel">
        <p>你可以通过顶部导航进入商品列表页与新增商品页。</p>
      </section>
    </div>
  );
}

export default DashboardPage;
