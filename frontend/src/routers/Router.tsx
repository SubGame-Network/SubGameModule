import React, { lazy, Suspense } from "react";
import { HashRouter, Switch, Route } from "react-router-dom";
import ScrollToTop from "../utils/ScrollToTop";
import Nav from "../components/Nav";
import Loading from "../components/Loading";

export const routes = [
  {
    pageName: "Home",

    href: "/",
    component: lazy(() => import("../pages/Home")),
  },
  {
    pageName: "ModuleManage",

    href: "/modulemanage",
    component: lazy(() => import("../pages/ModuleManage")),
  },
  {
    pageName: "ContacUs",

    href: "/contactus",
    component: lazy(() => import("../pages/ContacUs")),
  },
  {
    pageName: "ModuleDetail",

    href: "/moduledetail",
    component: lazy(() => import("../pages/ModuleDetail")),
  },
  {
    pageName: "UserInfo",

    href: "/userinfo",
    component: lazy(() => import("../pages/UserInfo")),
  },
  {
    pageName: "SignedUserInfo",

    href: "/signeduserinfo",
    component: lazy(() => import("../pages/SignedUserInfo")),
  },
  {
    pageName: "Module",

    href: "/module",
    component: lazy(() => import("../pages/Module")),
  },
];

const Routes = () => {
  return (
    <Switch>
      {routes.map((route) => {
        return (
          <Route
            path={route.href}
            exact
            component={route.component}
            key={route.pageName}
          />
        );
      })}
    </Switch>
  );
};
function Router() {
  return (
    <HashRouter>
      <Nav />
      <ScrollToTop />
      <Suspense fallback={<Loading />}>
        <Routes />
      </Suspense>
    </HashRouter>
  );
}

export default Router;
