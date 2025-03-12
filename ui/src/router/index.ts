import { createRouter, createWebHistory } from "vue-router";
import StockList from "@/views/StockList.vue";

const routes = [
  { path: "/", component: StockList },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
