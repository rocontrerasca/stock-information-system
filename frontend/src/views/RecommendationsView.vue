<template>
    <div class="p-8 bg-gray-50 min-h-screen">
      <h1 class="text-3xl font-bold text-center text-gray-800 mb-8">Recommended Stocks</h1>
  
      <!-- Mostrar mensaje de error -->
      <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-6 text-center">
        {{ error }}
      </div>
  
      <!-- Mostrar spinner durante la carga -->
      <div v-if="loading" class="flex justify-center items-center">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-500"></div>
      </div>
  
      <!-- Tabla de acciones -->
      <div v-else class="overflow-x-auto bg-white rounded-lg shadow-lg">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-800">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-white uppercase tracking-wider">Ticker</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-white uppercase tracking-wider">Company</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-white uppercase tracking-wider">Target Price</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-white uppercase tracking-wider">Rating</th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="stock in recommendations" :key="stock.ticker" class="hover:bg-gray-50 transition-colors">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{{ stock.ticker }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{{ stock.company }}</td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span class="bg-blue-100 text-blue-800 px-2 py-1 rounded-full text-xs font-semibold">
                  {{ stock.target_from }} - {{ stock.target_to }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                <span :class="ratingClass(stock.rating_to)">
                  {{ stock.rating_to }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { useStockStore } from "@/store/stockStore";
  import { onMounted } from "vue";
  import { storeToRefs } from 'pinia'
  
  const stockStore = useStockStore();
  const { recommendations, error } = storeToRefs(stockStore);
  
  onMounted(() => stockStore.fetchRecommendations());
  
  // Función para asignar clases dinámicas según el rating
  const ratingClass = (rating: string) => {
    switch (rating.toLowerCase()) {
      case "buy":
        return "bg-green-100 text-green-800 px-2 py-1 rounded-full text-xs font-semibold";
      case "sell":
        return "bg-red-100 text-red-800 px-2 py-1 rounded-full text-xs font-semibold";
      case "neutral":
        return "bg-yellow-100 text-yellow-800 px-2 py-1 rounded-full text-xs font-semibold";
      default:
        return "bg-gray-100 text-gray-800 px-2 py-1 rounded-full text-xs font-semibold";
    }
  };
  </script>