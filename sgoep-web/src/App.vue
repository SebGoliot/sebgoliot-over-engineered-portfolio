<template>
  <Header :author="this.author"/>
  <Portfolio :about="this.author.about" :skills="this.skills"/>
  <Footer/>
</template>

<script>
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';
import Portfolio from './components/Portfolio.vue';

export default {
  name: 'App',
  components: {
    Header,
    Portfolio,
    Footer,
  },
  data() {
    return {
      author: {},
      skills: {}
    };
  },
  created() {
    this.getAuthor();
    this.getSkills();
  },
  methods: {
    async getAuthor() {
      const r = await fetch("http://127.0.0.1:5050/author");
      this.author = await r.json();
      console.log(this.author);
    },
    async getSkills() {
      const r = await fetch("http://127.0.0.1:5050/skills");
      this.skills = await r.json();
      console.log(this.skills);
    },
  },
}
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
