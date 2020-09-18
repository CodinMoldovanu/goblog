<template>
<div class="home">
    <link rel="stylesheet" href="https://cdn.quilljs.com/1.3.6/quill.snow.css">
<h1>Add post</h1>


    

<div class="post">
    <div class="grid-post-meta-wrapper">
        <div class="meta">
<h2>Title input here</h2>
        </div>
        <div class="actions">
                <v-btn
            color="blue"
            @click="update"
            >Save</v-btn>
        </div>
    </div>
 <vue-editor  useCustomImageHandler @imageAdded="handleImageAdded" v-model="content" :editorOptions="editorOptions"></vue-editor>


</div>


</div>
</template>

<style scoped>
/* @import "@assets/vue2-editor/vue2-editor.css" */


html {
    background-color: rgb(3, 10, 39);
}
.grid-post-meta-wrapper {
      display: grid;
        grid-template-columns: repeat(4, 1fr);
        grid-gap: 10px;
        grid-auto-rows: minmax(100px, auto);
}
.meta {
    grid-column: 1 / 3;
    grid-row: 1;
    color: white;
}
.actions {
    grid-column: 4 / 5;
    grid-row: 1;
    text-align: right;
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
    color: black;
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
.ck-content {
    color: black !important;
}
.quillWrapper {
    background: white;
}

</style>


<script>
import axios from 'axios';
import Header from '@/components/Header.vue';
import { VueEditor } from "vue2-editor";
import 'vue2-editor/dist/vue2-editor.css';
import '@/assets/quill.snow.css'


export default {
    data() {
        return{
            post: {},
            content: "",
            components: {
                VueEditor
            }

        }
    },
    methods: {
        getPost () {
            var self = this;
            axios.get(this.$api_url + '/post/' + this.$route.params.post_URL).then(function (response) {
                self.post = response.data
                self.content = response.data.Content
            }).catch(function (error) {
                
            })
        },
        update () {
            console.log("HERE?")
            this.post.Content = this.content;
            var self = this;
            axios.post(this.$api_url + '/backend/posts/update', self.post).then(function (response) {
                // console.log(response)
            })
        },
        handleImageAdded (file, Editor, cursorLocation, resetUploader) {
            console.log("GOT HERE");
            var formdata = new FormData();
            formdata.append("image", file);

            axios.post(this.$api_url + '/backend/posts/media', formdata).then(function (response) {
                console.log("GOT INSIDE");
                let url = response.data.url;
                Editor.insertEmbed(cursorLocation, "image", url);
                resetUploader();
            }).catch(err => {
                console.log(err);
            })
        },
    },
    components: {
        Header: Header
    },
    mounted(){
        // this.getPost()
    },
    computed: {
        editorOptions() {
            return {
                theme: 'snow'
            }
        }
    }
}
</script>