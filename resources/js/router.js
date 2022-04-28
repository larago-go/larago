import { createWebHistory, createRouter } from "vue-router";

const routes =  [

  {
    path: "/index",
    name: "index",
    component: () => import("./Public/Index")
  },

  {
    path: "/welcome",
    alias: "/",
    name: "welcome",
    component: () => import("./Public/Views/Welcome")
  },

  /**
  *Admin panel
  **/

  {
    path: "/auth/login",
    name: "login",
    component: () => import("./Admin/Auth/Login")
  },

  {
    path: "/auth/register",
    name: "register",
    component: () => import("./Admin/Auth/Register")
  },
  
  {
    path: "/login/forgot_password",
    name: "forgot_password",
    component: () => import("./Admin/Auth/ForgotPassword")
  },
  
  {
    path: "/login/pass/:url",
    name: "reset_password",
    component: () => import("./Admin/Auth/ResetPassword")
  },

  {
    path: "/home",
    name: "home",
    component: () => import("./Admin/Views/Dashboard")
  },

  {
    path: "/users/add",
    name: "users_add",
    component: () => import("./Admin/Views/Users/Users_add")
  },

  {
    path: "/users/list",
    name: "users_list",
    component: () => import("./Admin/Views/Users/Users_list")
  },
  
  {
    path: "/users/list/:id",
    name: "users_list_prev",
    component: () => import("./Admin/Views/Users/Users_list_prev")
  },

  {
    path: "/role/add",
    name: "role_add",
    component: () => import("./Admin/Views/Casbinrole/Casbinrole_add")
  },

  {
    path: "/role/list",
    name: "role_list",
    component: () => import("./Admin/Views/Casbinrole/Casbinrole_list")
  },
  {
    path: "/example",
    name: "example",
    component: () => import("./Admin/Views/ExampleComponent")
  },

];

const router = createRouter({

  history: createWebHistory(),
  
  routes,

});

export default router;
