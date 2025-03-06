import { setActivePinia, createPinia } from 'pinia';
import { describe, it, expect, beforeEach } from 'vitest';
import { useStockStore } from "@/store/stockStore";

describe('Stock Store', () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });

    it('fetches stocks', async () => {
        const store = useStockStore();
        await store.fetchStocks();
        expect(store.stocks.length).toBeGreaterThan(0);
    });

    it('fetches recommendations', async () => {
        const store = useStockStore();
        await store.fetchRecommendations();
        expect(store.recommendations.length).toBeGreaterThan(0);
    });
});