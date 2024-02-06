<script setup lang="ts">
import { onBeforeMount, ref } from 'vue'
import DataTable from '@/components/DataTable.vue';
import NavBar from '../components/NavBar.vue'

interface UserResponse {
  ID: string
  email: string
  firstName: string
  lastName: string
}

const userData = ref<UserResponse>({
  ID: '',
  email: '',
  firstName: '',
  lastName: ''
})


async function getUserData() {
  const response = await fetch(`http://${import.meta.env.VITE_API_URL}/users/me`, {
    headers: {
      Authorization: `Bearer ${localStorage.getItem('jwt-token')}`
    }
  })
  const data = await response.json()

  userData.value = data
}

onBeforeMount(() => {
  getUserData()

})
</script>

<template>
  <div class="w-full h-[100vh] bg-slate-800">
    
    <NavBar/>
  
    <DataTable/>
  </div>

</template>
