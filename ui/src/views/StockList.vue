<script setup lang="ts">
import { onMounted } from "vue";
import { useStockStore } from "@/stores/stockStore";

const stockStore = useStockStore();

onMounted(() => {
  stockStore.fetchStocks();
});
</script>

<template>
  <div class="max-w-7xl w-full mx-auto p-8">
    <h1 class="text-red-500 text-4xl font-bold mb-6 text-center">📈 Stock Ratings</h1>

    <div v-if="stockStore.stocks.length === 0" class="text-gray-500 text-center">
      No data available.
    </div>

    <div class="overflow-x-auto bg-white shadow-lg rounded-xl p-4">
      <table class="w-full border-collapse rounded-lg overflow-hidden">
        <thead>
          <tr class="bg-gray-800 text-white text-lg">
            <th class="px-5 py-3 text-left">Ticker</th>
            <th class="px-5 py-3 text-left">Company</th>
            <th class="px-5 py-3 text-left">Brokerage</th>
            <th class="px-5 py-3 text-left">Action</th>
            <th class="px-5 py-3 text-center">Rating From</th>
            <th class="px-5 py-3 text-center">Rating To</th>
            <th class="px-5 py-3 text-right">Target From</th>
            <th class="px-5 py-3 text-right">Target To</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="stock in stockStore.stocks" :key="stock.ticker" class="border-b transition hover:bg-gray-100">
            <td class="px-5 py-4 text-left font-medium text-gray-700">{{ stock.ticker }}</td>
            <td class="px-5 py-4 text-left text-gray-600">{{ stock.company }}</td>
            <td class="px-5 py-4 text-left text-gray-600">{{ stock.brokerage }}</td>
            <td class="px-5 py-4 text-left text-gray-600">{{ stock.action }}</td>
            <td class="px-5 py-4 text-center font-semibold text-gray-700">{{ stock.rating_from }}</td>
            <td class="px-5 py-4 text-center font-semibold text-gray-700">{{ stock.rating_to }}</td>
            <td class="px-5 py-4 text-right font-mono text-gray-700">{{ stock.target_from }}</td>
            <td class="px-5 py-4 text-right font-mono text-gray-700">{{ stock.target_to }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>