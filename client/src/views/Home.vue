<!-- eslint-disable vue/multi-word-component-names -->
<script setup lang="ts">
import { ref,computed, onBeforeMount } from 'vue'
import Fuse from 'fuse.js';
import Header from '@/components/Header.vue'
import  type {Email, Hit} from '@/types/index'
import EmailDetails from '@/components/EmailDetails.vue'
import TotalResults from '@/components/TotalResults.vue'
import EmailsTable from '@/components/EmailsTable.vue'
import palabrasJson from '@/assets/words_dictionary.json'; 
const open = ref<boolean>(false)
const selectedEmail = ref<Email | undefined>()
const emails = ref<Hit[]>([])
const selectedEmailIndex = ref<number>(0)
const currentSearchTerm = ref<string>('')
const totalResults = ref<number>(0)
const currentPage = ref<number>(1)
const amountEmailsByPage = ref<number>(8)
const isLoading = ref<boolean>(false)
const conectionError = ref<boolean>(false)
const palabrasEnBaseDeDatos: string[] = Object.keys(palabrasJson);
const sugerencias= ref()
const totalPages = computed(() => Math.ceil(totalResults.value / amountEmailsByPage.value))


const getData = async (pageNumber: number, searchTerm?: string) => {
  searchTerm = typeof searchTerm === 'undefined' ? '' : searchTerm

  let strPageNumber = pageNumber.toString()
  currentSearchTerm.value = searchTerm
  isLoading.value = true

  try {
    const response = await fetch(
      `http://${import.meta.env.VITE_API_URL}/emails?page=${strPageNumber}&max=${amountEmailsByPage.value}&term=${searchTerm}`
    )

    const data = await response.json()
    emails.value = data.hits.hits
    totalResults.value = data.hits.Total.value

    if (totalResults.value===0){
      const fuse = new Fuse(palabrasEnBaseDeDatos, { threshold: 0.2 });
      const resultados = fuse.search(searchTerm);

      if (resultados.length > 0) {
        sugerencias.value = resultados
    .filter((resultado) => resultado.item.startsWith(searchTerm.slice(0, 2)))
    .map((resultado) => resultado.item);
        console.log("sugg",sugerencias.value)
      }else{
        sugerencias.value = [];
      }
    }

  } catch (error) {
    console.error('Error fetching data:', error)
    conectionError.value=true
  }
  isLoading.value = false
  currentPage.value = pageNumber
}

const asigneSelectedContent = (index: number) => {
  selectedEmail.value = emails.value[index]._source
  selectedEmailIndex.value = index
  open.value=true
}

onBeforeMount(() => {
  getData(1)
})
</script>

<template>

  <Header 
  :getData="getData"
 />



  <EmailsTable
    :getData="getData" 
    :emails ="emails"
    :asigneSelectedContent = "asigneSelectedContent"
    :currentSearchTerm="currentSearchTerm"
    :isLoading="isLoading"
    :totalResults="totalResults"
    :conectionError="conectionError"
    :totalPages="totalPages"
    :currentPage="currentPage"
    :suggestions="sugerencias"
    />

    <TotalResults v-if="!conectionError && totalResults !=0"
    :totalResults="totalResults"
    :isLoading="isLoading"
    :currentSearchTerm="currentSearchTerm"
    :currentPage="currentPage"
    :sugerencias="sugerencias"
    :getData="getData"
    />


  <EmailDetails v-if="open"
    :selectedEmail="selectedEmail"
    :asigneSelectedContent="asigneSelectedContent"
    :currentSearchTerm="currentSearchTerm"
    :selectedEmailIndex="selectedEmailIndex"
    @close="open=false"
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
