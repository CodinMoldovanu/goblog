<template>
 <v-app>
   <Header />
   <transition name="fade">
   <router-view></router-view>
  </transition>

   <!-- <Login /> -->
   <Footer />
 </v-app>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import Login from '@/components/Login.vue'

export default {
  computed: {
    layout () {
      return (this.$route.meta.layout || layout) + '-layout';
    }
  },
 name: 'App',
 components: { 
   'Footer': Footer, 
   'Login': Login, 
   'Header': Header,
   },
 data() {
   return {}
 }
}
</script>