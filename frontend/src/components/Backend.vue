<template>
<div class="home">
    
<h1>Edit Posts</h1>

    <div v-for="post in posts" v-bind:key="post.ID">
    <h3><router-link :to="{ name: 'editPost', params: { 'post_URL': post.URL }}">{{ post.Title }}</router-link></h3>
    </div>

<div class="posts">

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
            posts: [],
        }
    },
    methods: {
        getPosts () {
            var self = this;
            axios.get(this.$api_url).then(function (response) {
                self.posts = response.data
            }).catch(function (error) {
                console.log(error);
            })
        },
    },
    components: {
        Header: Header
    },
    mounted(){
        // this.posts = getPosts().data;
        this.getPosts()
    }
}
</script>