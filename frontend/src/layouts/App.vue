<template>
 <v-app>
   <!-- <Header /> -->
   <transition name="fade">
     <component :is="layout">
   <router-view></router-view>
     </component>
  </transition>

   <!-- <Login /> -->
   <Footer />
 </v-app>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
import Login from '@/components/Login.vue'
const default_layout = "default";

export default {
  
  computed: {
    layout () {
      return (this.$route.meta.layout || default_layout) + '-layout';
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