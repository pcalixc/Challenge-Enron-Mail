<template>
    <div class="splash-screen">
      <div class="splash-content">
        <!-- Contenido de tu pantalla splash -->
        <h1>Mi App</h1>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  export default {
    name: 'SplashScreen',
    mounted() {
      // Simulación de carga o tiempo de espera
      setTimeout(() => {
        // Navegar a la ruta principal o a la siguiente página
        this.$router.push({ name: 'Home' });
      }, 2000); // Cambia este valor al tiempo deseado en milisegundos
    }
  }
  </script>
  
  <style scoped>
  .splash-screen {
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  
  .splash-content {
    text-align: center;
  }
  </style>
  