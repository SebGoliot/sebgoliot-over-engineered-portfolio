<template>
  <PfHello :about="about" />
  <PfSkills :skills="skills" />
  <PfAchievements :achievements="achievements" />
  <PfInterests :interests="interests" />

</template>

<script>
import PfHello from './PfHello';
import PfSkills from './PfSkills';
import PfAchievements from './PfAchievements';
import PfInterests from './PfInterests';

export default {
  name: "Portfolio",
  components: {
    PfHello,
    PfSkills,
    PfAchievements,
    PfInterests,
  },
  props: {
    about: {},
    skills: {},
    achievements: {},
    interests: {},
  },
};

</script>

<style scoped></style>
