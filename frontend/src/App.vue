<template>
   <nav class="bg-blue-600 p-4 center">
        <router-link to="/" class="text-white mr-4">All Stocks</router-link>
        <router-link to="/recommendations" class="text-white">Recommendations</router-link>
    </nav>
  <router-view />
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>