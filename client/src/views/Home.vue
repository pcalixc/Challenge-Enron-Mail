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

  <EmailsTable/>

  <TotalResults/>

  <EmailDetails/>

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
