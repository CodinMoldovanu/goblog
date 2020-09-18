<template>
 <v-footer class="pa-3 white--text" color="teal" absolute>
   <div>
     Developed with absolute hate by {{YOUR_NAME}} &copy; {{ new Date().getFullYear() }}
   </div>
 </v-footer>
</template>
<script>
export default {
  data() {
    return {
      YOUR_NAME: 'Codin Moldovanu'
    }
  }
}
</script>