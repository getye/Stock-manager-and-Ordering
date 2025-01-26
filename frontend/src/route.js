
import { createRouter, createWebHistory } from 'vue-router'; 

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('./components/Home.vue') 
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('./components/Login.vue') 
  },
  {
    path: '/user',
    name: 'Users',
    component: () => import('./components/Users/AddUser.vue') 
  },
  {
    path: '/product',
    name: 'Products',
    component: () => import('./components/Products/AddProduct.vue') 
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes, 
});


