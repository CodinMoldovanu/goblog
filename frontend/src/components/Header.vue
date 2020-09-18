<template>
<div class="header">
    <div class="header-grid">
    <div class="bio">
    <router-link :to="'/'"><img class="logo" src="@/assets/codinbio.png" alt="Poza cu mine de cand eram mai slab lol"></router-link>
    <h1 class="title">Codin Moldovanu</h1>
    </div>
    <div class="meta">
        <h4>{{ randomQuote }}</h4>
    </div>
    </div>
    <hr>
    <span class="menu">
        <router-link :to="'/'">Home</router-link>
        <a href="/">About</a>
        <a href="/">Contact</a>
    </span>
    <hr>
</div>

</template>
<script>
export default {
    data () {
        return {
            randomQuote: ""
        }
    },
    methods: {
        randomText () {
            var self = this;
            let quotes = ["Make the internet ugly again!", "Served via RPi,Docker,MikroTik!", "This is slow.", "Random text here", "This is a work in progress you know"]
            let todaysPick = Math.floor(Math.random() * quotes.length)
            console.log(quotes[todaysPick])
            self.randomQuote += quotes[todaysPick]
        }
    },
    mounted() {
        this.randomText()
    }
}
</script>
<style scoped>
.header-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  grid-gap: 10px;
  grid-auto-rows: minmax(100px, auto);
}
.bio {
  grid-column: 1 / 3;
  grid-row: 1;
}
.meta {
  grid-column: 4 / 6;
  grid-row: 1;
}

.header {
    /* padding-top: 3em; */
    /* padding-left: 3em;
    padding-right: 3em; */
    padding-left: 1em;
    padding-right: 1em;
    text-align: center;
}
.menu a {
    padding-left: 2em;
    padding-right: 2em;
}
.header .logo {
    width: 10em;
    display: block;
}
hr {
    border: 0;
    height: 1px;
    background: #333;
    background-image: linear-gradient(to right, #333, #ccc, #333);
}
body {
    padding-left: 5em;
    padding-right: 5em;
    padding-top: 1em;
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