import { createWebHistory, createRouter } from 'vue-router';

const routes = [

  {
    path: '/index',
    name: 'index',
    component: () => import('./Public/IndexHeader.vue'),
  },
  {
    path: '/welcome',
    alias: '/',
    name: 'welcome',
    component: () => import('./Public/Views/Welcome.vue'),
  },

  /*
  * Admin panel
  */

  {
    path: '/auth/login',
    name: 'login',
    component: () => import('./Admin/Auth/LoginAuth.vue'),
  },
  {
    path: '/auth/register',
    name: 'register',
    component: () => import('./Admin/Auth/RegisterAuth.vue'),
  },
  {
    path: '/login/forgot_password',
    name: 'forgot_password',
    component: () => import('./Admin/Auth/ForgotPassword.vue'),
  },
  {
    path: '/login/pass/:url',
    name: 'reset_password',
    component: () => import('./Admin/Auth/ResetPassword.vue'),
  },
  {
    path: '/home',
    name: 'home',
    component: () => import('./Admin/Views/DashboardAdmin.vue'),
  },
  {
    path: '/users/list',
    name: 'users_list',
    component: () => import('./Admin/Views/Users/UsersList.vue'),
  },
  {
    path: '/users/add',
    name: 'users_add',
    component: () => import('./Admin/Views/Users/UsersAdd.vue'),
  },
  {
    path: '/users/list/:id',
    name: 'users_list_prev',
    component: () => import('./Admin/Views/Users/UsersListPrev.vue'),
  },
  {
    path: '/role/list',
    name: 'role_list',
    component: () => import('./Admin/Views/Casbinrole/CasbinroleList.vue'),
  },
  {
    path: '/role/add',
    name: 'role_add',
    component: () => import('./Admin/Views/Casbinrole/CasbinroleAdd.vue'),
  },
  {
    path: '/example',
    name: 'example',
    component: () => import('./Admin/Views/ExampleComponent.vue'),
  },

];

const router = createRouter({

  history: createWebHistory(import.meta.env.BASE_URL),
  routes,

});

export default router;
