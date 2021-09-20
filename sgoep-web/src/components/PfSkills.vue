<template>
  <section>
    <header>
      <span class="material-icons-round">assignment</span>
      <h2>&nbsp;Compétences</h2>
    </header>
    <pfSkillCard :skills="skills" />
  </section>
</template>

<script>
import PfSkillCard from "./PfSkillCard.vue";

export default {
  name: "PfSkills",
  props: {
    skills: {},
  },
  components: {
    PfSkillCard,
  },
};
</script>
