import { NavLink, Outlet } from "react-router-dom";

const menuItems = [
  { label: "首页", path: "/" },
  { label: "商品列表", path: "/products" },
  { label: "新增商品", path: "/products/new" },
];

function MainLayout() {
  return (
    <div className="app-shell">
      <header className="app-header">
        <h1 className="app-title">商品管理系统</h1>
        <nav className="app-nav">
          {menuItems.map((item) => (
            <NavLink
              key={item.path}
              to={item.path}
              end={item.path === "/"}
              className={({ isActive }) =>
                isActive ? "nav-link nav-link-active" : "nav-link"
              }
            >
              {item.label}
            </NavLink>
          ))}
        </nav>
      </header>
      <main className="app-main">
        <Outlet />
      </main>
    </div>
  );
}

export default MainLayout;
