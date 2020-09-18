<template>
 <v-app>
   <transition name="fade">
   <router-view></router-view>
  </transition>

   <!-- <Login /> -->
   <Footer />
 </v-app>
</template>

<script>
import Footer from '@/components/Footer.vue'
import Login from '@/components/Login.vue'

export default {
 name: 'Admin',
 components: { 
   'Footer': Footer, 
   'Login': Login, 
   },
 data() {
   return {}
 }
}
</script>