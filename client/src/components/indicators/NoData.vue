<script setup lang="ts">
import { LottieAnimation } from 'lottie-web-vue'
import animationNotData from '@/assets/NoData.json'
import { useEmailsStore } from '@/stores/emails';
const emailsStore = useEmailsStore()

</script>
<template>
  <div v-if="emailsStore.totalResults == 0 && !emailsStore.isLoading && !emailsStore.ServerErrorResponse.errorStatus" class="flex flex-col align-middle justify-center m-auto">
    <div
      class="flex bg-yellow-100 rounded-lg dark:bg-deep_blue dark:text-white p-4 mb-4 mt-9 w-auto text-sm text-yellow-700"
      role="alert"
    >
      <svg
        class="w-5 h-5 inline mr-3"
        fill="currentColor"
        viewBox="0 0 20 20"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z"
          clip-rule="evenodd"
        ></path>
      </svg>
      <div>
        <span class="font-medium">No emails matched the term "{{ emailsStore.currentSearchTerm }}" </span>

        <span v-if="emailsStore.searchSuggestion.length>0 && emailsStore.searchSuggestion[0] != emailsStore.currentSearchTerm">
          Did you mean
          <span
            @click="emailsStore.getData(1, emailsStore.searchSuggestion[0])"
            class="underline font-semibold cursor-pointer"
            >{{ emailsStore.searchSuggestion[0]}}</span>
          ?</span>
      </div>
    </div>
    <div class="h-full flex justify-center align-middle mt-12">
      <lottie-animation :animation-data="animationNotData" :loop="true" class="w-72 h-72" />
    </div>
  </div>
</template>
