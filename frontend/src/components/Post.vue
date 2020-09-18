<template>
<div class="home">
    <div class="post">
    <h3 class="post_title">{{ post.Title }}</h3>
    <div class="post_content" v-html="post.Content"></div>
</div>
</div>
</template>

<style scoped>
html {
    background-color: rgb(3, 10, 39);
    color:white;
}
.home {
    padding-left: 5em;
    padding-right: 5em;
    padding-top: 1em;
    margin-bottom: 1em;
    padding-bottom: 5em;
}
.post {
    width: 60vw;
    text-align: justify;
    margin: auto;
}
li {
    padding-top: 1em;
    padding-bottom: 1em;
}
img {
    display: block;
    padding-top: 1em;
    padding-bottom: 1em;
    border-radius: 40px;
    margin: auto;
}
</style>

<script>
import axios from 'axios'
import Header from '@/components/Header.vue'

export default {
    data() {
        return{
            post: {},
        }
    },
    methods: {
        getPost () {
            var self = this;
            axios.get(this.$api_url + '/post/' + this.$route.params.post_URL).then(function (response) {
                self.post = response.data
            }).catch(function (error) {
                
            })
        },
    },
    components: {
        Header: Header
    },
    mounted(){
        this.getPost()
    }
}
</script>