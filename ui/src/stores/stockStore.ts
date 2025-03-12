import { defineStore } from "pinia";
import { api } from "@/services/api";

export interface Stock {
  ticker: string;
  company: string;
  brokerage: string;
  action: string;
  rating_from: string;
  rating_to: string;
  target_from: number;
  target_to: number;
  time: string;
}

export interface StockRecommendation extends Stock {
  score: number;
}

export const useStockStore = defineStore("stock", {
  state: () => ({
    stocks: [] as Stock[],
    recommendation: null as StockRecommendation | null,
  }),

  actions: {
    async fetchStocks() {
      try {
        const response = await api.get<Stock[]>("/stocks");
        this.stocks = response.data.map(stock => ({
          ...stock,
          target_from: parseFloat(stock.target_from.toString().replace("$", "")), 
          target_to: parseFloat(stock.target_to.toString().replace("$", "")), 
        }));
      } catch (error) {
        console.error("❌ Error fetching stocks:", error);
      }
    },

    async fetchRecommendation() {
      try {
        const response = await api.get<StockRecommendation>("/recommendation");
        this.recommendation = {
          ...response.data,
          target_from: parseFloat(response.data.target_from.toString().replace("$", "")), 
          target_to: parseFloat(response.data.target_to.toString().replace("$", "")),
        };
      } catch (error) {
        console.error("❌ Error fetching recommendation:", error);
      }
    },
  },
});
