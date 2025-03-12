<script setup lang="ts">
import { onMounted, ref } from "vue";
import { useStockStore } from "@/stores/stockStore";

const stockStore = useStockStore();
const isLoadingRecommendation = ref(false);

onMounted(() => {
  stockStore.fetchStocks();
});


const fetchRecommendationOnClick = async () => {
  isLoadingRecommendation.value = true;
  await stockStore.fetchRecommendation();
  isLoadingRecommendation.value = false;
};
</script>


<template>
  <div class="max-w-7xl w-full mx-auto p-8">
    <h1 class="text-red-500 text-4xl font-bold mb-6 text-center flex items-center justify-center gap-3">
      📈 Stock Ratings
    </h1>
    <div class="text-center mb-6">
      <button
        @click="fetchRecommendationOnClick"
        class="bg-blue-500 hover:bg-blue-700 text-white text-lg font-bold py-3 px-8 rounded-lg 
              transition shadow-lg border border-blue-900 focus:outline-none focus:ring-4 focus:ring-blue-300"
        style="display: inline-block; background-color: #2563eb; color: white;"
      >
        🔥 Get Best Investment
      </button>
    </div>



    <div v-if="stockStore.recommendation" class="bg-green-100 p-6 rounded-lg shadow-md mb-6 text-center">
      <h2 class="text-green-700 text-2xl font-bold">🔥 Best Investment Opportunity</h2>
      <p class="text-gray-700 mt-2">
        <strong>{{ stockStore.recommendation.ticker }}</strong> - 
        {{ stockStore.recommendation.company }} ({{ stockStore.recommendation.brokerage }})
      </p>
      <p class="mt-1 text-gray-600">
        Rating Change: <strong>{{ stockStore.recommendation.rating_from }}</strong> → 
        <strong>{{ stockStore.recommendation.rating_to }}</strong>
      </p>
      <p class="mt-1 text-gray-600">
        Target Price: <strong>${{ Number(stockStore.recommendation.target_from).toFixed(2) }}</strong> → 
        <strong>${{ Number(stockStore.recommendation.target_to).toFixed(2) }}</strong>
      </p>
      <p class="mt-2 text-green-700 font-semibold">Investment Score: {{ Number(stockStore.recommendation.score).toFixed(2) }}</p>
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
            <td class="px-5 py-4 text-right font-mono text-gray-700">${{ Number(stock.target_from).toFixed(2) }}</td>
            <td class="px-5 py-4 text-right font-mono text-gray-700">${{ Number(stock.target_to).toFixed(2) }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
