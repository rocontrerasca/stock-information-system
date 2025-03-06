import { mount } from '@vue/test-utils';
import RecommendationsView from "@/views/RecommendationsView.vue";
import { describe, it, expect, beforeEach } from 'vitest';
import { createPinia, setActivePinia } from 'pinia';

describe('RecommendationsView', () => {
    beforeEach(() => {
        setActivePinia(createPinia());
    });
    
    it('renders the component', () => {
        const wrapper = mount(RecommendationsView);
        expect(wrapper.text()).toContain('Recommended Stocks');
    });
});