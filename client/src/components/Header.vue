<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import ToggleTheme from './ToggleTheme.vue';
import LogoSection from './LogoSection.vue'
import SearchBar from './SearchBar.vue'

</script>

<template>
  <nav
    id="header"
    class="container mx-auto h-20 dark:bg-primary_dark bg-white rounded-lg p-4 mt-0 pt-2 dark:shadow-md dark:shadow-ligth_blue shadow-violet-100 shadow-md">
    <div class="w-full responsive-container flex justify-between px-6 py-2">
      <LogoSection/>

      <SearchBar/>

      <ToggleTheme/>

    </div>
  </nav>
</template>


<style scoped>
@media (max-width: 768px) {
  .responsive-container {
    display: flex;
    justify-content: center; 
    align-items: center;
    gap: 3rem;
  }
}
</style>
