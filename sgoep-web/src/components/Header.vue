<template>
  <div class="main-header wrapper-blue">
    <h1 class="header-name">{{ author.name }}</h1>
    <h3>{{ author.title }}</h3>
    <h4>{{ author.subtitle }}</h4>
    <a :href="author.github" target="_blank" rel="noopener noreferrer">
      <img src="../assets/svg/github-icon.svg" alt="Github" />
    </a>
    <a :href="author.linkedin" target="_blank" rel="noopener noreferrer">
      <img src="../assets/svg/linkedin-icon.svg" alt="Linkedin" />
    </a>
  </div>
  <img src="../assets/svg/wave-to-alpha.svg" alt="" />
</template>

<script>
export default {
  name: "Header",
  props: {
    author: {},
  },
};
</script>

<style scoped>
.main-header {
  padding: 2em 0;
  text-align: center;
}
</style>
