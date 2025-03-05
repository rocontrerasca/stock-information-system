import { defineStore } from "pinia";
import { ref } from "vue";

export const useStockStore = defineStore("stocks", () => {
    const stocks = ref<any>([]);
    const recommendations = ref<any>([]);
    const error = ref<string | null>(null);
    const loading = ref(false);

    const fetchStocks = async () => {
        error.value = null;
        loading.value = true;
        try {
            const response = await fetch("http://localhost:8080/stocks");
            if (!response.ok) {
                throw new Error(response.status + response.statusText);
            }
            const data =  await response.json();
            stocks.value = data;
        } catch (err) {
            error.value = "Failed to fetch stocks. Please try again later.";
            console.error(err);
        } finally {
            loading.value = false;        }
    };

    const fetchRecommendations = async () => {
        error.value = null;
        loading.value = true;
        try {
            const response = await fetch("http://localhost:8080/recommendations");

            if (!response.ok) {
                throw new Error(response.status + response.statusText);
            }

            recommendations.value = await response.json();
        } catch (err) {
            error.value = "Failed to fetch recommendations. Please try again later.";
            console.error(err);
        } finally {
            loading.value = false;
        }
    };

    return { stocks, recommendations, error, loading, fetchStocks, fetchRecommendations };
});