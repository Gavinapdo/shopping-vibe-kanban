import { Link } from "react-router-dom";
import PageTitle from "../components/common/PageTitle";

function NotFoundPage() {
  return (
    <div>
      <PageTitle title="页面不存在" description="请检查访问路径是否正确。" />
      <section className="panel">
        <p>
          你访问的页面不存在，请返回<Link to="/">首页</Link>继续操作。
        </p>
      </section>
    </div>
  );
}

export default NotFoundPage;
