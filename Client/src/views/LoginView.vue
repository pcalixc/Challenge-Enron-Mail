<script setup lang="ts">
import { useRouter } from 'vue-router'
import { ref } from 'vue'

const email = ref('')
const password = ref('')
const errorMsg = ref('')
const router = useRouter()
const errorAlert = ref(false);

async function login() {
  if(password.value.length==0 || email.value.length==0){
    errorMsg.value = 'fill in all fields'
    errorAlert.value = true
    setTimeout(() => {
      errorAlert.value = false
    }, 3000)
    return
  }

  const response = await fetch('http://localhost:8080/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      email: email.value,
      password: password.value
    })
  }
  )
  const data: { message?: string; token?: string } = await response.json()
  if (data.message) {
    errorMsg.value = data.message
    console.log(errorMsg)
    errorAlert.value = true;
    setTimeout(() => {
      errorAlert.value = false;
    }, 3000);
  }
  if (data.token) {
    localStorage.setItem('jwt-token', data.token)
    router.push('/')
  }

}

</script>


<template>

<main
  class="mx-auto flex min-h-screen w-full items-center justify-center bg-gray-900 text-white"
>
  <!-- component -->
  <section class="flex w-[30rem] flex-col space-y-10">
    <div class="text-center text-4xl font-medium">Log In</div>

    <div
      class="w-full transform border-b-2 bg-transparent text-lg duration-300 focus-within:border-indigo-500"
    >
      <input
        type="text"
        v-model="email"
        placeholder="Email"
        class="w-full border-none bg-transparent outline-none placeholder:italic focus:outline-none"
      />
    </div>

    <div
      class="w-full transform border-b-2 bg-transparent text-lg duration-300 focus-within:border-indigo-500"
    >
      <input
        type="password"
        v-model="password"
        placeholder="Password"
        class="w-full border-none bg-transparent outline-none placeholder:italic focus:outline-none"
      />
    </div>

    <div class=" h-8 m-0">
          <div
          v-if="errorAlert"
          class="inline-block align-baseline text-sm text-red-600 ml-2"
        >  {{errorMsg}}
        </div>
        </div>

    <button
    @click="login()"
      class="transform rounded-sm bg-indigo-600 py-2 font-bold duration-300 hover:bg-indigo-400"
    >
      LOG IN
    </button>

    <a
      href="#"
      class="transform text-center font-semibold text-gray-500 duration-300 hover:text-gray-300"
      >FORGOT PASSWORD?</a
    >

    <p class="text-center text-lg">
      No account?
    <router-link to="/register" class="font-medium text-indigo-500 underline-offset-4 hover:underline">Sign up</router-link>

    </p>
  </section>
</main>
</template>
