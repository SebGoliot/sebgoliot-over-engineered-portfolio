<template>
  <div class="vue-wrapper">
    <Header :author="this.author" />
    <Portfolio
      :about="this.author.about"
      :skills="this.skills"
      :achievements="this.achievements"
      :interests="this.interests"
    />
    <Footer />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
import Portfolio from "./components/Portfolio.vue";

export default {
  name: "App",
  components: {
    Header,
    Portfolio,
    Footer,
  },
  data() {
    return {
      author: {},
      skills: {},
      achievements: {},
      interests: {},
    };
  },
  created() {
    this.getAuthor();
    this.getSkills();
    this.getAchievements();
    this.getInterests();
  },
  methods: {
    async getAuthor() {
      const r = await fetch("http://127.0.0.1:5050/author");
      this.author = await r.json();
    },
    async getSkills() {
      const r = await fetch("http://127.0.0.1:5050/skills");
      this.skills = await r.json();
    },
    async getAchievements() {
      const r = await fetch("http://127.0.0.1:5050/achievements");
      this.achievements = await r.json();
    },
    async getInterests() {
      const r = await fetch("http://127.0.0.1:5050/interests");
      this.interests = await r.json();
    },
  },
};
</script>

<style>
#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
