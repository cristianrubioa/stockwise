import { defineStore } from "pinia";
import { api } from "@/services/api";

export interface Stock {
  ticker: string;
  company: string;
  brokerage: string;
  action: string;
  rating_from: string;
  rating_to: string;
  target_from: string;
  target_to: string;
  time: string;
}

export const useStockStore = defineStore("stock", {
  state: () => ({
    stocks: [] as Stock[],
  }),

  actions: {
    async fetchStocks() {
      try {
        const response = await api.get<Stock[]>("/stocks");
        this.stocks = response.data;
      } catch (error) {
        console.error("❌ Error fetching stocks:", error);
      }
    },
  },
});
