<template>
  <section>
    <header>
      <span class="material-icons-round">assignment</span>
      <h2>&nbsp;Technos qui m'intéressent</h2>
    </header>
    <PfInterestCard :interests="interests" />
  </section>
</template>

<script>
import PfInterestCard from "./PfInterestCard";

export default {
  name: "PfInterests",
  props: {
    interests: {},
  },
  components: {
    PfInterestCard,
  },
};
</script>
