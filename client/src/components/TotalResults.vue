<script setup lang="ts">
import { defineProps } from 'vue'
interface Props {
    totalResults : number
}

const props = defineProps<Props>()

</script>

<template>
      <div  class="flex flex-row align-middle justify-center mt-5">
    <p class="text-slate-700 text-xs dark:text-slate-200 underline">
      {{ props.totalResults }} emails founded
    </p>
  </div>

</template>