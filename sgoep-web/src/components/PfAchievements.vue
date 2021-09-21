<template>
  <img class="wave-to-blue" src="../assets/svg/wave-to-blue.svg" alt="" />
  <div class="wrapper-blue">
    <section>
      <header>
        <span class="material-icons-round">integration_instructions</span>
        <h2>&nbsp;Réalisations</h2>
      </header>
      <PfAchievementCard :achievements="achievements" />
    </section>
  </div>
  <img src="../assets/svg/wave-to-alpha.svg" alt="" />
</template>

<script>
import PfAchievementCard from "./PfAchievementCard";

export default {
  name: "PfAchievements",
  props: {
    achievements: {},
  },
  components: {
    PfAchievementCard,
  },
};

</script>
