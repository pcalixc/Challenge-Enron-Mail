<script setup lang="ts">
import { ref, watch } from 'vue'
const searchTerm = ref('')
const errorMessage = ref('')
import { useEmailsStore } from '@/stores/emails';
const emailsStore = useEmailsStore()

watch(() => emailsStore.currentSearchTerm, (newValue) => {
  searchTerm.value = newValue
})

const handleSubmit = () => {
  const regex = /[?/><>,!@#$%^&*()_=+{}|[\]\\:;"'`~]/
  if (regex.test(searchTerm.value)) {
    errorMessage.value = 'Input must not contain special characters.'
    return
  }
  emailsStore.getData(1, searchTerm.value)
  errorMessage.value = ''
}
</script>

<template>
  <div class="flex items-center m-0   flex-col justify-center max-w-md">
    <form
      @submit.prevent="handleSubmit()"
      class="relative text-slate-800 dark:text-slate-50"
    >
      <input
        type="search"
        placeholder="Enter a keyword..."
        v-model="searchTerm"
        class="peer lg:w-[28vw] cursor-pointer relative shadow-inner z-10 h-12 rounded-xl border-2 border-[#2525ff52] dark:bg-[#28287549] bg-transparent pl-16 pr-4 outline-none focus:cursor-text focus:border-[#4a4aff5c] focus:border-3"
      />
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="absolute inset-y-0 my-auto h-8 w-12 border-r border-transparent stroke-gray-500 px-3.5 dark:peer-focus:border-slate-600 peer-focus:border-slate-300 dark:peer-focus:stroke-slate-200 peer-focus:stroke-slate-500"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        stroke-width="2"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
        />
      </svg>
    </form>
    <span class="flex absolute mt-16 items-center tracking-wide text-red-500 text-xs">{{
      errorMessage
    }}</span>
  </div>
</template>
