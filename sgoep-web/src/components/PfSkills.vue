<template>
  <section>
    <header>
      <span class="material-icons-round">assignment</span>
      <h2>&nbsp;Compétences</h2>
    </header>
    <pfSkillCards :skills="skills" />
  </section>
</template>

<script>
import PfSkillCards from "./PfSkillCards.vue";

export default {
  name: "PfSkills",
  props: {
    skills: {},
  },
  components: {
    PfSkillCards,
  },
};
</script>
