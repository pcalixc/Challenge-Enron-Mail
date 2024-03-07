<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { onBeforeMount } from 'vue'
import Header from '@/components/Header.vue'
import EmailDetails from '@/components/EmailDetails.vue'
import TotalResults from '@/components/TotalResults.vue'
import EmailsTable from '@/components/EmailsTable.vue'
import { useEmailsStore } from '@/stores/emails';
const emailsStore = useEmailsStore()

onBeforeMount(() => {
  emailsStore.getData(1)
})

</script>

<template>

  <Header/>

  <EmailsTable
    :getData="emailsStore.getData" 
    :emails ="emailsStore.emails"
    :asigneSelectedContent = "emailsStore.asigneSelectedContent"
    :currentSearchTerm="emailsStore.currentSearchTerm"
    :isLoading="emailsStore.isLoading"
    :totalResults="emailsStore.totalResults"
    :conectionError="emailsStore.conectionError"
    :totalPages="emailsStore.totalPages"
    :currentPage="emailsStore.currentPage"
    :suggestions="emailsStore.suggestion"
    />

    <TotalResults v-if="!emailsStore.conectionError && emailsStore.totalResults !=0"
    :totalResults="emailsStore.totalResults"
    :isLoading="emailsStore.isLoading"
    :currentSearchTerm="emailsStore.currentSearchTerm"
    :currentPage="emailsStore.currentPage"
    :sugerencias="emailsStore.suggestion"
    :getData="emailsStore.getData"
    />


  <EmailDetails v-if="emailsStore.modalOpen"
    :selectedEmail="emailsStore.selectedEmail"
    :asigneSelectedContent="emailsStore.asigneSelectedContent"
    :currentSearchTerm="emailsStore.currentSearchTerm"
    :selectedEmailIndex="emailsStore.selectedEmailIndex"
    @close="emailsStore.modalOpen=false"
    />

</template>

<style>

::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}


::-webkit-scrollbar-thumb {
  background-color: #9292FF;
  border-radius: 5px;
}

::-webkit-scrollbar-thumb:hover {
  background-color: #555;
}

::-webkit-scrollbar-track {
  background-color: #f1f1f1;
}
</style>
