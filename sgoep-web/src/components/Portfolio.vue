<template>
  <PfHello :about="about"/>
  <PfSkills :skills="skills"/>
</template>

<script>
import PfHello from './PfHello.vue';
import PfSkills from './PfSkills.vue';

export default {
  name: "Portfolio",
  components: {
    PfHello,
    PfSkills,
  },
  props: {
    about: {},
    skills: {},
  }
};
</script>

<style scoped></style>
