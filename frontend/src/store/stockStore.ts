import { defineStore } from "pinia";
import { ref } from "vue";
import type { Stock } from "@/types/Stock";
import { apiFetch } from "@/utils/api"; // Importamos apiFetch

export const useStockStore = defineStore("stocks", () => {
    const stocks = ref<Stock[]>([]);
    const recommendations = ref<Stock[]>([]);
    const error = ref<string | null>(null);
    const loading = ref(false);

    const fetchStocks = async () => {
        loading.value = true;
        const data = await apiFetch<Stock[]>("/stocks");
        if (data) stocks.value = data;
        loading.value = false;
    };

    const fetchRecommendations = async () => {
        loading.value = true;
        const data = await apiFetch<Stock[]>("/recommendations");
        if (data) recommendations.value = data;
        loading.value = false;
    };

    return { stocks, recommendations, error, loading, fetchStocks, fetchRecommendations };
});