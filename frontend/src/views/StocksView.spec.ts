import { mount } from '@vue/test-utils';
import StocksView from "./StocksView.vue";
import { describe, it, expect, beforeEach } from 'vitest';
import { setActivePinia, createPinia } from 'pinia';

describe('AllStocksView', () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });
    it('renders the component', () => {
        const wrapper = mount(StocksView);
        expect(wrapper.text()).toContain('Stock Information');
    });
});