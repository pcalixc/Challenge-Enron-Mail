<script setup lang="ts">
import ConnectionError from '@/components/indicators/ConnectionError.vue'
import NoData from '@/components/indicators/NoData.vue'
import Loading from '@/components/indicators/Loading.vue'
import Pagination from './Pagination.vue'
import EmailCard from './EmailCard.vue'
import { useEmailsStore } from '@/stores/emails'
const emailsStore = useEmailsStore()

// This function checks if a current SearchTerm is present 'only' in the content of the email.
// const isWordInContent = (subject: string, from: string, to: string, term: string) => {
//   const regex = new RegExp(term, 'gi')
//   if (regex.test(subject) || regex.test(from) || regex.test(to)) {
//     return false
//   }
//   return true
// }
</script>

<template>
  <!-- component -->
  <section v-if="emailsStore.totalResults > 0 && !emailsStore.isLoading">
    <div class="py-5 lg:px-16 sm:px-8 h-[75vh]">
      <div class="mx-auto px-3 max-w-8xl text-gray-500">
        <div class="mt-0 grid sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-3">
          <EmailCard />
        </div>
      </div>
    </div>
  </section>

  <NoData />

  <ConnectionError />

  <Loading v-if="emailsStore.isLoading" />

  <Pagination />
</template>
