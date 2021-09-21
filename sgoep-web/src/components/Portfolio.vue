<template>
  <PfHello :about="about" />
  <PfSkills :skills="skills" />
  <PfAchievements :achievements="achievements" />
</template>

<script>
import PfHello from './PfHello';
import PfSkills from './PfSkills';
import PfAchievements from './PfAchievements';

export default {
  name: "Portfolio",
  components: {
    PfHello,
    PfSkills,
    PfAchievements,
  },
  props: {
    about: {},
    skills: {},
    achievements: {},
  },
};

</script>

<style scoped></style>
